# cut (tr, grep, ...)

## Examples

```
# fields
cut -f 1,5 -d ':'/etc/passwd
cut -f 1-3,7 -d ':'/etc/passwd

# translate char
tr -s ''''
tr -s " "
tr -s "\t" " " | tr -s " "

# limiters
cut -f 1,5 -d ':' --output-delimiter=$'\t' /etc/passwd

# ls group by dates
ls -l | tr -s ' ' | cut -d' ' -f 7,6 | sort | uniq -c

#awk ?
awk -F '===' '{ print $1}'


```

and some log parsing

```
grep "http~ NAME" /var/log/haproxy.log | cut -d' ' -f7 | cut -d':' -f1-2 | sort | uniq -c
grep "http~ NAME" /var/log/haproxy.log | awk '//{print $6}' | cut -d':' -f1 | sort | uniq -c | sort -nr
grep "http~ NAME" /var/log/haproxy.log| grep api|grep -E 'user|login' | cut -d' ' -f7 | cut -d':' -f1-2 | sort | uniq -c
```
