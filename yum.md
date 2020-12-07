# Cheatsheet

https://access.redhat.com/sites/default/files/attachments/rh_yum_cheatsheet_1214_jcs_print-1.pdf

# list

```
yum list available
yum list installed
```

# epel

## awsl2

```
cd /tmp
wget -O epel.rpm –nv \
https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum install -y ./epel.rpm
```

link https://www.cyberciti.biz/faq/installing-rhel-epel-repo-on-centos-redhat-7-x/

# Dev tools

```
yum groupinstall 'Development Tools'
```
