func getAndBroadcastLogs(clientset *kubernetes.Clientset, podName string) {
	// Channels to signal the completion of existing logs and the buffering of new logs
	existingLogsDone := make(chan bool)
	newLogsBuffer := make(chan []byte, 1000)
	stopChan := hub.stopReadChannels[podName]

	// Function to handle existing logs
	go func() {
		podLogOpts := v1.PodLogOptions{
			Follow: false,
		}
		req := clientset.CoreV1().Pods("default").GetLogs(podName, &podLogOpts)
		podLogs, err := req.Stream(context.Background())
		if err != nil {
			errorMessage := ServerResponse{
				EventType: "error",
				Value:     fmt.Sprintf("Error in opening stream for pod %s: %v", podName, err),
			}
			msgJSON, _ := json.Marshal(errorMessage)
			hub.broadcast <- msgJSON
			existingLogsDone <- true
			return
		}
		defer podLogs.Close()

		scanner := bufio.NewScanner(podLogs)
		for scanner.Scan() {
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
		if err := scanner.Err(); err != nil {
			errorMessage := ServerResponse{
				EventType: "error",
				Value:     fmt.Sprintf("Error in scanner for pod %s: %v", podName, err),
			}
			msgJSON, _ := json.Marshal(errorMessage)
			hub.broadcast <- msgJSON
		}
		existingLogsDone <- true
	}()

	// Function to handle new logs
	go func() {
		podLogOpts := v1.PodLogOptions{
			Follow: true,
		}
		req := clientset.CoreV1().Pods("default").GetLogs(podName, &podLogOpts)
		podLogs, err := req.Stream(context.Background())
		if err != nil {
			errorMessage := ServerResponse{
				EventType: "error",
				Value:     fmt.Sprintf("Error in opening stream for pod %s: %v", podName, err),
			}
			msgJSON, _ := json.Marshal(errorMessage)
			hub.broadcast <- msgJSON
			return
		}
		defer podLogs.Close()

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
				newLogsBuffer <- msgJSON
			}
		}
		if err := scanner.Err(); err != nil {
			errorMessage := ServerResponse{
				EventType: "error",
				Value:     fmt.Sprintf("Error in scanner for pod %s: %v", podName, err),
			}
			msgJSON, _ := json.Marshal(errorMessage)
			hub.broadcast <- msgJSON
		}
	}()

	// Wait for the existing logs to be processed before adding new logs to the buffer
	go func() {
		<-existingLogsDone
		for msg := range newLogsBuffer {
			hub.podBuffers[podName].mu.Lock()
			hub.podBuffers[podName].buffer = append(hub.podBuffers[podName].buffer, msg)
			hub.podBuffers[podName].mu.Unlock()
			hub.broadcast <- msg
		}
	}()
}
