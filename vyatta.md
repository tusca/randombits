# op : /opt/vyatta/bin/vyatta-op-cmd-wrapper
/opt/vyatta/bin/vyatta-op-cmd-wrapper show vrrp


# cfg : /opt/vyatta/sbin/vyatta-cfg-cmd-wrapper
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper begin
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper delete interfaces openvpn vtun0 disable
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper delete interfaces openvpn vtun1 disable
/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper commit
