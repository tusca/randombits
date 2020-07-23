# Ubuntu Bits

## How Tos

Change default screensaver folder

  - dconf-editor (alt+F2) -> org / gnome / gnome-screenshot / auto-save-directory

https://discourse.ubuntu.com/t/how-to-set-default-screenshots-folder-in-ubuntu-18-04/6578

## IDEA Icon

~/.local/share/applications/jetbrains-idea.desktop
```
[Desktop Entry]
Version=1.0
Type=Application
Name=IntelliJ IDEA Ultimate Edition
Icon=/opt/idea-X/bin/idea.svg
Exec="/opt/idea-X/bin/idea.sh" %f
Comment=Capable and Ergonomic IDE for JVM
Categories=Development;IDE;
Terminal=false
StartupWMClass=jetbrains-idea
```

## Nemo

Source: https://itsfoss.com/install-nemo-file-manager-ubuntu/

```
sudo apt install nemo
xdg-mime default nemo.desktop inode/directory application/x-gnome-saved-search
sudo apt-get install dconf-tools
gsettings set org.gnome.desktop.background show-desktop-icons false
xdg-open $HOME
```

## APT

```
sudo apt autoclean
sudo apt clean
sudo apt autoremove --purge
```

## Window Control

Requires `sudo apt-get install wmctrl`

then

bringToFront bash script is

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

then associate with keyboard shortcut
