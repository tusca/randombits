# fish history script

Intent: easy search history with grep without typing grep with multiple keys (reversed as used on mac)

Second intent: support multiple histories represented by commands history, history2, history3, ...

Script name: h

Usage:
```
h MATCH1 MATCH2 ...
```

Usage alternate history:
```
h -2 MATCH1 MATCH2 ...
```

script
```
#!/usr/bin/env fish

set history_cmd "history"
set args $argv

if test (count $args) -gt 0
    set -l matches (string match -r "^-([0-9]+)\$" -- $args[1])
    if test (count $matches) -eq 2
        set history_cmd "history$matches[2]"
        set args $args[2..-1]
    end
end

set cmd "$history_cmd"
for arg in $args
    if test "$arg" = "|"
        break
    end
    set cmd "$cmd | grep \"$arg\""
end
set cmd "$cmd | tac -r"

eval $cmd
```
