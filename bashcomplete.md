# Bash Autocomplete

## Interdependent parameters

```
SCRIPT=SCRIPT.py
 
 
 
 
_suggest_PARAMs()
{
    local cur opts   # define some variables
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"  # get the current value on which you are auto-completing
    case $COMP_CWORD in
       1) # parameter in 1st position
        opts=`$SCRIPT list PARAM1s`
        ;;
       2) # parameter in 2st position
        PARAM1="${COMP_WORDS[COMP_CWORD-1]}"
        opts=`$SCRIPT list PARAM2s --PARAM1 $PARAM1`
        ;;
       3) # parameter in 3st position
        PARAM1="${COMP_WORDS[COMP_CWORD-2]}"
        PARAM2="${COMP_WORDS[COMP_CWORD-1]}"
        opts=`$SCRIPT list PARAM3s --PARAM1 $PARAM1 --PARAM2 $PARAM2`
        ;;
       4) # parameter in 4st position
        PARAM1="${COMP_WORDS[COMP_CWORD-3]}"
        PARAM2="${COMP_WORDS[COMP_CWORD-2]}"
        PARAM3="${COMP_WORDS[COMP_CWORD-1]}"
        opts=`$SCRIPT list PARAM4s --PARAM1 $PARAM1 --PARAM2 $PARAM2 --PARAM3 $PARAM3`
        ;;
       *) # parameter in other positions
        opts=""
        ;;
    esac
    #echo $opts
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )   # uses the compgen command to generate all possible autocompletes based on a list of possibilities (opts) and the part already written on the command line (cur)
    return 0
}
 
complete -F _suggest_PARAMs mycmd

```
command to use autocomplete on is "mycmd" that you have to implement
SCRIPT.py does generate the possible parameters as follows:
- "SCRIPT.py list PARAM1s" returns a space separated list of parameters to be used as 1st parameter
- "SCRIPT.py list PARAM2s --PARAM1 XXX" returns a space separated list of parameters to be used as 2nd parameter taking in consideration that the 1st parameter "XXX "

