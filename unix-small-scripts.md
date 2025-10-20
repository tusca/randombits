# h - history

```
#!/usr/bin/env fish

set cmd "history"
for arg in $argv
    if test "$arg" = "|"
        break
    end
    set cmd "$cmd | grep \"$arg\""
end
set cmd "$cmd | tac -r"

eval $cmd⏎
```
