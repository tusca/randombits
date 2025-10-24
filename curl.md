# CURL

## Host Header

```
curl --verbose --header 'Host: www.example.com' 'http://10.1.1.36:8000/the_url_to_test' [--insecure]
```

## Post with json

```
#!/bin/bash

export FILENAME=$1
export URL=$2
export TOKEN=$3

if [ -z "$TOKEN" ]; then
  echo "Usage: $0 <filename> <url> <token>"
  exit 1
fi

echo "Posting $FILENAME to $URL with token $TOKEN"
echo curl -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" --data @$FILENAME $URL
curl -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" --data @$FILENAME $URL



```
