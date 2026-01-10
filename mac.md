# Fix Keyboard Shortcuts

Download https://gist.githubusercontent.com/trusktr/1e5e516df4e8032cbc3d/raw/ab7c868a8354219782b37971f984564c66a53d78/DefaultKeyBinding.dict to

```
~/Library/KeyBindings/DefaultKeyBinding.dict
```

# Sudo can use touch id
see In macOS Sonoma, Touch ID for sudo can survive updates 

https://sixcolors.com/post/2023/08/in-macos-sonoma-touch-id-for-sudo-can-survive-updates/

```
cd /etc/pam.d
sudo cp sudo_local.template sudo_local
```

then edit sudo_local to remove the comment and have


```
# sudo_local: local config file which survives system update and is included for sudo
# uncomment following line to enable Touch ID for sudo
auth       sufficient     pam_tid.so
```
