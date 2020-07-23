# AWS

## Ssh into instanceid

Prerequisites: configured `aws` command, installed `ssh` as well as pem file(s) in `$PATH_TO_PEMS`
(replace PrivateIpAddress with PublicIpAddress if needed)

```
function sshi() {
   INSTANCE=$1
   PATH_TO_PEMS=/home/USER/.ssh/
   if [[ "$INSTANCE" == "" ]]; then
      echo "usage: sshi INSTANCEID"
   else
      echo "ssh to $INSTANCE"
      CMD=$(aws ec2 describe-instances --instance-id $INSTANCE | jq -r '.Reservations[].Instances[] | "ssh -i $PATH_TO_PEMS\(.KeyName).pem ec2-user@\(.PrivateIpAddress)"')
      echo $CMD
      $CMD
   fi
 
}
```


