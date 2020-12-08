# MongoDB

## Dump

```
mongodump -h HOST -d DB -o FOLDER
```

it'll probably save to `FOLDER/DB/*`


## Get DB Names

```
mongo admin --eval "var databaseNames = db.getMongo().getDBNames(); for (var i in databaseNames) { printjson(db.getSiblingDB(databaseNames[i]).getCollectionNames()) };"
```
