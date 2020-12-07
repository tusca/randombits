
packets are stored in /var/cache/apt/archives.

- purge old packets

```
sudo apt autoclean
```

- fully clean cache

```
sudo apt clean
```

- redownload cached packets

```
dpkg -l | grep "^ii" | awk ' {print $2} ' | xargs sudo apt-get -y --force-yes install --reinstall --download-only
```


- remove unneeded packages

```
sudo apt autoremove --purge
```
