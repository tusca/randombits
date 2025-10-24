# Fish


## Activate py env from name of current folder

```shell
#!/bin/bash
FOLDER=$(pwd | rev | cut -d '/' -f 1 | rev)
echo source ~/.pyenv/versions/$FOLDER/bin/activate.fish
```

## Search History

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

eval $cmd

```
