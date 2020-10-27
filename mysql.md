# MySQL related bits

## Start server in docker

with latest mariadb (see versions other than `latest` on https://hub.docker.com/_/mariadb

```
docker run --name mariadb -p3306:3306 -e MYSQL_ROOT_PASSWORD=ROOT_PWD -v ~/mariadb_data:/var/lib/mysql -d mariadb:latest

#then 

docker start mariadb
docker stop mariadb
docker exec -ti mariadb /bin/bash
```

with mysql, replace `mariadb` and `mariadb:latest` with `mysql` and `mysql:latest` or `mysql:5.7` or similar (see https://hub.docker.com/_/mysql).

## Backup and restore

usage
```
mysqldump -u [user] -p -h [host] [database] > db_backup.dump
mysql -p -u[user] -h [host] [database] < db_backup.dump
```

## See connections

```
select * from information_schema.processlist
```

## IP Issues with docker

```
sudo sysctl -w net.ipv6.conf.all.forwarding=1
```

