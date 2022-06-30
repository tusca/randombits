# Python Bits

## Arguments with argparse

```
from argparse import ArgumentParser


def main(a1, a2, p1, p2, v, i):
    print(f'arg1 is {a1}')
    print(f'arg2 is {a2}')
    print(f'param1 is {p1}')
    print(f'param2 is {p2}')
    print(f'verbosity is {v}')
    print(f'my tailor is {i}')


if __name__ == "__main__":
    p = ArgumentParser(description="I'm a super CLI tool, use me !")
    p.add_argument('arg1', help='Arg1')  # mandatory argument
    p.add_argument('arg2', help='Arg2')  # mandatory argument
    p.add_argument('-p', '--param1', help='Param1') # parameter1 also mandatory, can be used as "-p SOMETHING" or "--param1 SOMETHING"
    p.add_argument('-q', '--param2', required=False, default='Hi', help='Param2') # parameter2 is optional
    p.add_argument('-v', '--verbose', action='store_true', help='print more') # also optional, if not given then it's False, otherwise true
    p.add_argument('-i', type=int, help="give me a number or else") # mandatory and is a number (without type=int it's a string)
    a = p.parse_args()
    main(a.arg1, a.arg2, a.param1, a.param2, a.verbose, a.i)
```
used as
```
$ python argparsing.py hello world --param1 me -q 25 -v -i 36
arg1 is hello
arg2 is world
param1 is me
param2 is 25
verbosity is True
my tailor is 36
```

## Formatting

```
from datetime import datetime

a = 25
b = 'victory'
d = datetime.now()

print(f'{a} {b} {d}')
print('%s %s %s' % (a, b, d))
print('%(x)s %(y)s %(z)s' % {'x': a, 'y': b, 'z': d})
print('{x} {y} {z}'.format(x=a, y=b, z=d))

print(f"now is {d.strftime('%Y-%m-%d %H:%M:%S')}")
e = datetime.strptime('2020-12-12 15:19:17', '%Y-%m-%d %H:%M:%S')
duration = datetime.now() - e
print(f'{e} was {duration} ago, which is approx {duration.days} days ago')

```


## print/get exceptions

```
try:
  ...
except:
  traceback.print_exc() 
  exception = traceback.format_exc()
```

## System and as import
```
from platform   import system as system_name  # Returns the system/OS name
from subprocess import call   as system_call  # Execute a shell command
 
if system_name().lower()=='windows' ...
```

## Json NamedTuple Hook
use obj.field instead of obj['field'] (wink)
```
from json import loads
from collections import namedtuple
 
converter = lambda d: namedtuple('MyClazz', d.keys())(*d.values())
obj = loads('{ "field" : "stuff" }', object_hook=converter)
print(obj.field)
```

also see https://github.com/tusca/python_config_files/blob/main/example_json_simple.py

## Bash complete

```
apt install python3-pip
pip3 completion --bash >> ~/.bashrc
```


## pyenv

```
# Install on linux
curl https://pyenv.run | bash
# Install on Mac
brew install pyenv-virtualenv
# or just brew install pyenv (not sure)

# List available versions
pyenv install --list

# List installed versions
pyenv versions

# List created virtual environents
pyenv virtualenvs

# Install version
pyenv install 3.9.0b5

# Create a virtual environmet named pm39 usign python 3.9.0b5
pyenv virtualenv 3.9.0b5 pm39

# Activate pm39 venv
pyenv activate pm39 

# Deactivate/Leave it
pyenv deactivate
 
# Uninstall version
pyenv uninstall 3.9.0b5

# UPDATE !!!!! to get new versions
pyenv update
```



typically will ask to add to .bashrc
```
export PATH="/home/guy/.pyenv/bin:$PATH"
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"
```

in case of dependency issues, google it and/or see
```
19814  python -m pip install readline
19815  sudo python -m pip install readline
19816  sudo python3 -m pip install readline
19817  pyenv install python3.7.8
19818  pyenv install python-3.7.8
19819  pyenv install 3.7.8
19820  sudo apt-get remove libssl-dev
19821  sudo apt-get update
19822  sudo apt-get install libssl1.0-dev
19823  apt search bz2 python
19824  apt install python3-bz2file python-bz2file
19825  sudo apt install python3-bz2file python-bz2file
19826  apt search readline python
20444  apt search libffi-dev
20445  apt install libffi-dev
20446  sudo apt install libffi-dev
```

## Unbuffered
source: https://stackoverflow.com/questions/107705/disable-output-buffering

```
import sys
class Unbuffered(object):
    def __init__(self, stream):
        self.stream = stream

    def write(self, data):
        self.stream.write(data)
        self.stream.flush()

    def writelines(self, datas):
        self.stream.writelines(datas)
        self.stream.flush()

    def __getattr__(self, attr):
        return getattr(self.stream, attr)
```


## venv

```
sudo apt install -y virtualenv virtualenvwrapper
echo "source /usr/share/virtualenvwrapper/virtualenvwrapper.sh" >> ~/.bashrc
 
 
export WORKON_HOME=~/.virtualenvs
mkdir $WORKON_HOME
echo "export WORKON_HOME=$WORKON_HOME" >> ~/.bashrc
# make sure that if pip creates an extra virtual environment, it is also placed in our WORKON_HOME directory
echo "export PIP_VIRTUALENV_BASE=$WORKON_HOME" >> ~/.bashrc
 
 
#and don't forget
source ~/.bashrc
```
Source: https://askubuntu.com/questions/244641/how-to-set-up-and-use-a-virtual-python-environment-in-ubuntu

### Use

```

# create
mkvirtualenv [-p PYTHON_VERSION] [NAME_OF_ENV]
mkvirtualenv -p $(which python3) NAMEOFENV
 
# check
python -c "import sys; print sys.path"
 
 
# deactivate
deactivate
 
 
# activate
workon NAME_OF_ENV
 
 
# remove
rmvirtualenv NAME_OF_ENV
```


## Unit Tests

```
pip3 install pytest pyfakefs jinja2
python3 -m pytest
python3 -m pytest tests/test_httpd.py -k 'test_noproxy'
# or add details :
python3 -m pytest tests/test_httpd.py -k 'test_noproxy' -vv
```


## MySQL

```
connection = mysql.connector.connect(
    user=arguments.dbUsername,
    password=arguments.dbPassword,
    host=arguments.dbHostname,
    port=int(arguments.dbPort),
    database=arguments.dbSchema)
    

cursor = connection.cursor()
 
updateSql = """
   UPDATE tablename
      SET field1=%(fieldA)s, field2=%(fieldB)s
    WHERE field3 = %(fieldC)s
"""
updateQuery = (updateSql)
cursor.execute(updateSql, {'fieldA': valueA, 'fieldB': valueB, 'fieldC': valueC })
cursor.close()


cursor = connection.cursor()
selectSql = """
   SELECT colA, colB
     FROM tablename
     WHERE field1=%(fieldA)s, field2=%(fieldB)s
"""
 
selectQuery = (selectSql)
cursor.execute(selectSql, {'fieldA': valueA, 'fieldB': valueB, 'fieldC': valueC })
for (colA, colB) in cursor:
    print("(%s, %s)" % (colA, colB)
cursor.close()
```

### show process list using open with

```
#!/usr/bin/env python3
import mysql.connector
from json import loads
from pathlib import Path
from collections import namedtuple
  
 
 
 
 
class connect:
    def __init__(self, user, pwd, host, port, schema):
        self.parameters = (user, pwd, host, port, schema)
 
    def __enter__(self):
        (user, pwd, host, port, schema) = self.parameters
        self.connection = mysql.connector.connect(
            user=user, password=pwd,
            host=host, port=int(port),
            database=schema
        )
        return self.connection
 
    def __exit__(self, type, value, traceback):
        self.connection.close()
 
 
# Create query cursor which is closed after
class execute_query:
    def __init__(self, connection, query, parameters=tuple()):
        self.connection = connection
        self.query = query
        self.parameters = parameters
 
    def __enter__(self):
        self.cursor = self.connection.cursor()
        q = ( self.query )
        self.cursor.execute(q, self.parameters)
        return self.cursor
 
    def __exit__(self, type, value, traceback):
        self.cursor.close()
 
def loadconfig():
    converter = lambda d: namedtuple('Config', d.keys())(*d.values())
    return loads(Path(Path.home(), '.showprocesslist.conf').read_text(encoding='utf8'), object_hook=converter)
 
config = loadconfig()
db = config.db
 
query = 'show processlist;'
with connect(db.user, db.password, db.hostname, db.port, db.schema) as connection:
    with execute_query(connection, query) as cursor:
        for record in cursor:
            (id, user, host, db, command, time, state, info) = record
            if info is not None:
                print('[%s] "%s" on db %s' % (user, info, db))
```
with config
```
{
    "db": {
        "user":"xxxxx",
        "password":"****",
        "hostname":"xxxx",
        "port":3306,
        "schema": "mysql"
    }
}
```
## URLS

### urllib
```
#!/usr/bin/env python3
import urllib
 
 
parameters = {'from': date, 'productId': productId}
 
post_data = urllib.parse.urlencode(parameters)
post_data = post_data.encode('UTF-8')
 
request = urllib.request.Request(url, post_data)
request.add_header("User-Agent", USER_AGENT)
 
response = urllib.request.urlopen(request)
body = response.read().decode('utf8', 'ignore')
```
### urllib2
```
#!/usr/bin/env python2.7
import urllib2
import base64
 
 
auth = base64.encodestring('%s:%s' % (username, password)).replace('\n', '')
request = urllib2.Request(url)
request.add_header("Authorization", "Basic %s" % auth)  
 
body = urllib2.urlopen(request).read()
```
### requests
```
from requests import session
with session() as c:
   response = c.get(url)  
   response.status_code
   response.content
```
### httplib2
```
import httplib2
h = httplib2.Http(".cache")
h.add_credentials('name', 'password')
(resp, content) = h.request("https://example.org/chapter/2",
                            "PUT", body="This is text",
                            headers={'content-type':'text/plain'} )
```
## Parse your json output
```
import json
# deserialize
json.loads(string)
# serialize
json.dumps(object, indent=4)
```

# Flask

## pymysql

```
mysql+pymysql://user:pwd@hostname/schema
```

-  http://flask.pocoo.org/docs/1.0/patterns/sqlalchemy/

# S3

http://s3fs.readthedocs.io/en/latest/api.html 
