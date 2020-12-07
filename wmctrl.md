# bringToFront

Using the title of a window of an app either launch or bring to front this app

```
sudo apt-get install wmctrl
```

## script

```
#!/bin/bash
# $1 : command to launch
# $2 : part of the window title (unique enough)
if [ `wmctrl -l | grep -c "$2"` != 0 ] 
then
    wmctrl -a "$2"
else
    $1 &
fi
```

## usage

create a shortcut (eg. gnome keyboard F9)

- example command : `bringToFront tilix "Tilix"
