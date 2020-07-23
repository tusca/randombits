# Find

## Find and Delete

```
find /PATH -name FILENAME -exec rm -f {} \;
```

## Delete older files

```
# 3 days
find /PATH  -mtime +3 -exec rm -rf {} \;
# *.tmp files of more than 40 min old and only 1 level deep
find /PATH -name "*.tmp" -maxdepth 1 -mmin 40 -exec rm -rf {} \;
```

## Grep

```
find . -name FILENAME -exec grep PATTERN {} \; -print
```
