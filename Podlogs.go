package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client represents a WebSocket client
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// PodBuffer contains the buffer and the lock for a specific pod
type PodBuffer struct {
	buffer [][]byte
	mu     sync.Mutex
}

// Hub maintains the set of active clients and broadcasts messages to them
type Hub struct {
	clients          map[*Client]bool
	register         chan *Client
	unregister       chan *Client
	broadcast        chan []byte
	podSubscribers   map[string]map[*Client]bool
	podBuffers       map[string]*PodBuffer
	stopReadChannels map[string]chan bool
}

var hub = Hub{
	clients:          make(map[*Client]bool),
	register:         make(chan *Client),
	unregister:       make(chan *Client),
	broadcast:        make(chan []byte),
	podSubscribers:   make(map[string]map[*Client]bool),
	podBuffers:       make(map[string]*PodBuffer),
	stopReadChannels: make(map[string]chan bool),
}

// JSON message structs
type ClientRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type ServerResponse struct {
	EventType string `json:"eventtype"`
	Value     string `json:"value"`
}

// Function to read and broadcast logs
func getAndBroadcastLogs(clientset *kubernetes.Clientset, podName string) {
	podLogOpts := v1.PodLogOptions{
		Follow: true,
	}
	req := clientset.CoreV1().Pods("default").GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		hub.broadcast <- []byte(fmt.Sprintf("Error in opening stream for pod %s: %v", podName, err))
		return
	}
	defer podLogs.Close()

	stopChan := hub.stopReadChannels[podName]
	scanner := bufio.NewScanner(podLogs)
	for scanner.Scan() {
		select {
		case <-stopChan:
			return
		default:
			line := scanner.Bytes()
			message := ServerResponse{
				EventType: "podlog",
				Value:     fmt.Sprintf("%s: %s", podName, line),
			}
			msgJSON, _ := json.Marshal(message)
			hub.podBuffers[podName].mu.Lock()
			hub.podBuffers[podName].buffer = append(hub.podBuffers[podName].buffer, msgJSON)
			hub.podBuffers[podName].mu.Unlock()
			hub.broadcast <- msgJSON
		}
	}
	if err := scanner.Err(); err != nil {
		hub.broadcast <- []byte(fmt.Sprintf("Error in scanner for pod %s: %v", podName, err))
	}
}

// WebSocket handler
func serveWs(clientset *kubernetes.Clientset, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	hub.register <- client

	go client.writePump()
	client.readPump(clientset)
}

// Client read pump to handle incoming messages
func (c *Client) readPump(clientset *kubernetes.Clientset) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		var request ClientRequest
		if err := json.Unmarshal(message, &request); err != nil {
			continue
		}
		if request.Action == "subscribe" && request.Key == "pod" {
			hub.mu.Lock()
			podName := request.Value
			if _, ok := hub.podSubscribers[podName]; !ok {
				hub.podSubscribers[podName] = make(map[*Client]bool)
				hub.podBuffers[podName] = &PodBuffer{
					buffer: make([][]byte, 0),
					mu:     sync.Mutex{},
				}
				hub.stopReadChannels[podName] = make(chan bool)
				go getAndBroadcastLogs(clientset, podName)
			}
			hub.podSubscribers[podName][c] = true
			hub.mu.Unlock()

			// Send buffered messages to new client
			hub.podBuffers[podName].mu.Lock()
			for _, msg := range hub.podBuffers[podName].buffer {
				c.send <- msg
			}
			hub.podBuffers[podName].mu.Unlock()
		}
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.mu.Lock()
				for podName, subs := range h.podSubscribers {
					if _, exists := subs[client]; exists {
						delete(subs, client)
						if len(subs) == 0 {
							close(h.stopReadChannels[podName])
							delete(h.podSubscribers, podName)
							delete(h.podBuffers, podName)
							delete(h.stopReadChannels, podName)
						}
					}
				}
				h.mu.Unlock()
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	go hub.run()

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		serveWs(clientset, c)
	})

	r.Run(":8080")
}
