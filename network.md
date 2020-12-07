# Static routes

https://my.esecuredata.com/index.php?/knowledgebase/article/2/add-a-static-route-on-centos

# iptables save 

centos7

```

sudo yum -y install iptables-services
sudo systemctl enable iptables 
sudo service iptables save
```

and debian
```
sudo apt-get install iptables-persistent
sudo netfilter-persistent save
```

# masquerade

one possible way
```
sudo iptables -t nat -A POSTROUTING -s x.x.x.x/x -d x.x.x.x/x -j MASQUERADE
```

# ipsec network routing

```
ip xfrm policy

ip xfrm state -s
```
