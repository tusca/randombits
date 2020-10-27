# Docker

## bash into running container

```
docker exec -ti ID_OR_NAME /bin/bash
```
(assuming the container does have `/bin/bash`)


## Running image but with shell i/o executing default app

```
docker run --rmi -ti NAME_OR_ID /bin/bash
```
(`--rm` being to delete the container after being created from the image)

## Mysql, also see mysql.md

```
docker run --name mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=dev -v ~/mysql/mydb:/var/lib/mysql -d mysql:5.7
```

this might require some ipv6 sh*t

```
sudo sysctl -w net.ipv6.conf.all.forwarding=1
```



## Delete all

```
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker volume rm $(docker volume ls --filter dangling=true -q)
docker rmi -f $(docker images --filter dangling=true -qa)
docker rmi $(docker images -qa)
```

