# Python Bits


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

## Bash complete

```
apt install python3-pip
pip3 completion --bash >> ~/.bashrc
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
