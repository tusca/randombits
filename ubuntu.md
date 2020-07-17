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
