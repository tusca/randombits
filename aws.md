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

## print instances

```
aws ec2 describe-instances --query "Reservations[].Instances[].{id: InstanceId, image:ImageId,type:Tags[?Key=='Type'][]|[0].Value,name:Tags[?Key=='Name'].Value[]|[0], nid:NetworkInterfaces[].PrivateIpAddresses[].[{private:PrivateIpAddress, public:Association.PublicIp}][]}"   --output text | sed '$!N;s/\nNID//;P;D'
```

also use awless tool ?

## Boto 3 pagination

```
import boto3
def aws_fetch(fn, args={}, name=None) -> Iterator:
    while args.get('NextToken', None) != '':
        result = fn(**args)
        args['NextToken'] = result.get('NextToken', '')
        try:
            items = result[name] if name is not None else result
        except KeyError as err:
            logger.error(f'Failed to find value for key {name} where keys are {result.keys()}')
            raise
        for item in items:
            yield item

def flatten_describe_instances(reservations):
    return [instance for reservation in reservations for instance in reservation['Instances']]

reservations = aws_fetch(boto3.client('ec2').describe_instances, name='Reservations')
instances = flatten_describe_instances(reservations)


```

## Boto3 read x.json.tar.gz from s3

```


from io import BytesIO
from json import loads
import boto3
def read_json_tar_gz_from_s3(s3path):
   bucket, path = s3path.split('/', 1)
   with BytesIO() as fh:
      boto3.client('s3').download_fileobj(bucket, path, fh)
      return loads(decompress(fh.getbuffer()))
      
```
