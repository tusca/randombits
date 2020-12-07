# Run jenkins in docker

```
mkdir home
export LOCALPORT=8080
docker container run -d -p $LOCALPORT:8080 -v $(pwd)/home:/var/jenkins_home --name myjenkins jenkins/jenkins:lts
docker logs myjenkins
```
