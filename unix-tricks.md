# use PID file to kill processes and delete files

```
FOLDER=/path/to/somewhere
# kill all
for f in $(find $FOLDER -name RUNNING_PID -exec echo {} \;); do kill $(cat $f); done
# delete files
find $FOLDER -name RUNNING_PID -exec rm -f {} \;
```

set correct folder ;-)

