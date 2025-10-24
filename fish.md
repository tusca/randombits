# Fish


## Activate py env from name of current folder

```shell
#!/bin/bash
FOLDER=$(pwd | rev | cut -d '/' -f 1 | rev)
echo source ~/.pyenv/versions/$FOLDER/bin/activate.fish
```
