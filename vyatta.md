# op : /opt/vyatta/bin/vyatta-op-cmd-wrapper
```
/opt/vyatta/bin/vyatta-op-cmd-wrapper show vrrp
```

# cfg : /opt/vyatta/sbin/vyatta-cfg-cmd-wrapper
```
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper begin
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper delete interfaces openvpn vtun0 disable
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper delete interfaces openvpn vtun1 disable
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper commit
```
# VRRP - updown tunnels

```
#!/bin/vbash
WR=/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper
LOG=/var/log/vrrp_vtun.log
 
/opt/vyatta/bin/vyatta-op-cmd-wrapper show vrrp>$LOG
 
if /opt/vyatta/bin/vyatta-op-cmd-wrapper show vrrp | grep master; then
 
   echo "$(date) start">>$LOG
   $WR begin
   $WR delete interfaces openvpn vtun0 disable 2>&1 >>$LOG
   $WR delete interfaces openvpn vtun1 disable 2>&1 >>$LOG
   $WR commit
   echo "$(date) started">>$LOG
 
else
 
   echo "$(date) start">>$LOG
   $WR begin
   $WR set interfaces openvpn vtun0 disable 2>&1 >>$LOG
   $WR set interfaces openvpn vtun1 disable 2>&1 >>$LOG
   $WR commit
   echo "$(date) started">>$LOG
 
fi
```
