# nftables

# Links

* [Performing Network Address Translation (NAT) - nftables wiki](https://wiki.nftables.org/wiki-nftables/index.php/Performing_Network_Address_Translation_(NAT))
* [nftables - ArchWiki](https://wiki.archlinux.org/index.php/nftables)
* [Simple rule management - nftables wiki](https://wiki.nftables.org/wiki-nftables/index.php/Simple_rule_management)
* [Differences between iptables and nftables explained - Linux Audit](https://linux-audit.com/differences-between-iptables-and-nftables-explained/)
* [Quick reference-nftables in 10 minutes - nftables wiki](https://wiki.nftables.org/wiki-nftables/index.php/Quick_reference-nftables_in_10_minutes)


# Install

```
// Ubuntu 18.04
apt remove ufw firewalld

apt install nftables
```

# Enough for a Router


```
nft add table nat
nft add chain nat prerouting { type nat hook prerouting priority 0 \;  }
nft add chain nat postrouting { type nat hook postrouting priority 100 \;  }

nft add rule nat postrouting oif ens4 masquerade


// nft list ruleset

table ip nat {
        chain prerouting {
                type nat hook prerouting priority 0; policy accept;
        }

        chain postrouting {
                type nat hook postrouting priority 100; policy accept;
                oif "enp0s8" masquerade
        }
}
```

# TCP Redirecter

* [iptables - nftables dnat forwarding doesn't work properly - Server Fault](https://serverfault.com/questions/895611/nftables-dnat-forwarding-doesnt-work-properly/899152#899152)

1. Type nat hook prerouting priority 0; <-Should be **-100** (minus hundred) according to `NF_IP_PRI_NAT_DST` netfilter constant
1. Use `meta nftrace set 1` in prerouting and `nft monitor` for debug you packet flow
1. Simplest ruleset to isolate problem and make example


```
table ip nat {
        chain prerouting {
                type nat hook prerouting priority -100; policy accept;
                ip daddr { 0.0.0.0/8, 10.0.0.0/8, 127.0.0.0/8, 169.254.0.0/16, 172.16.0.0/12, 192.168.0.0/16, 224.0.0.0/4, 240.0.0.0/4, 172.16.39.0/24} return
                ip daddr <remote_server> return
                # By using redirect, packets will be forwarded to local machine. 
                # Is a special case of DNAT where the destination is the current machine. 
                # tcp sport {32768-61000} redirect to 12345     # /proc/sys/net/ipv4/ip_local_port_range
                ip protocol tcp redirect to 12345  


        }

        chain input {
                type nat hook input priority 0; policy accept;
        }


        # this is for local packages generated from "lo"
        chain output {
                type nat hook output priority 0; policy accept;
                ip daddr { 0.0.0.0/8, 10.0.0.0/8, 127.0.0.0/8, 169.254.0.0/16, 172.16.0.0/12, 192.168.0.0/16, 224.0.0.0/4, 240.0.0.0/4, 172.16.39.0/24} return 
                ip daddr <remote_server> return 
                ip protocol tcp redirect to 12345  
        }

        chain postrouting {
                type nat hook postrouting priority 0; policy accept;
        }

}
```

# DNAT

This redirects the incoming traffic for TCP ports 80 and 443 to 192.168.1.120. 

```
nft add table nat
nft add chain nat prerouting { type nat hook prerouting priority -100 \; }
nft add chain nat postrouting { type nat hook postrouting priority 100 \; }

nft add rule nat prerouting iif eth0 tcp dport { 80, 443 } dnat 192.168.1.120

nft add rule ip nat prerouting iif eth0 tcp dport 12345 counter dnat to 1.2.3.4:8388
nft add rule nat postrouting ip daddr 1.2.3.4 tcp dport 8388  oif wg0 masquerade
```

[Linux 2.4 NAT HOWTO: Destination NAT Onto the Same Network](https://www.netfilter.org/documentation/HOWTO/NAT-HOWTO-10.html)

**the Destination server may have no route to the true client, so we still do
a SNAT or masquerade.**

# Redirect

By using redirect, packets will be forwarded to local machine. Is a special
case of DNAT where the destination is the current machine. 

```
nft add rule nat prerouting redirect
nft add rule nat prerouting tcp dport 22 redirect to 2222
```

# SNAT

This matches for all traffic from the 192.168.1.0/24 network to the interface
eth0. The IPv4 address 1.2.3.4 is used as source for the packets that match
this rule. 

```
nft add table nat
nft add chain nat prerouting { type nat hook prerouting priority 0 \; }
nft add chain nat postrouting { type nat hook postrouting priority 100 \; }

nft add rule nat postrouting ip saddr 192.168.1.0/24 oif eth0 snat 1.2.3.4

nft add rule ip nat postrouting daddr 10.0.7.2 tcp dport 8388 counter snat to 10.0.7.1

```

# Masquerading

Masquerade is a special case of SNAT, where the source address is automagically
set to the address of the output interface.

```
nft add table nat
nft add chain nat prerouting { type nat hook prerouting priority 0 \; }
nft add chain nat postrouting { type nat hook postrouting priority 100 \; }

nft add rule nat postrouting oif ens4 masquerade
```

# Basic

```
nft list tables
nft list table nat [-a]
nft list ruleset [-a]

// Tables
% nft list tables [<family>]
% nft list table [<family>] <name> [-n] [-a]
% nft (add | delete | flush) table [<family>] <name>

// Chains
% nft (add | create) chain [<family>] <table> <name> [ { type <type> hook <hook> [device <device>] priority <priority> \; [policy <policy> \;] } ]
% nft (delete | list | flush) chain [<family>] <table> <name>
% nft rename chain [<family>] <table> <name> <newname>

// Rules
% nft add rule [<family>] <table> <chain> <matches> <statements>
% nft insert rule [<family>] <table> <chain> [position <position>] <matches> <statements>
% nft replace rule [<family>] <table> <chain> [handle <handle>] <matches> <statements>
% nft delete rule [<family>] <table> <chain> [handle <handle>]
```

# interactive mode

```
nft -i
```

# Save and Reload Configuration

```
// save
nft list ruleset > /etc/nftables.conf
// load
nft -f filename
```

# iptables-restore-translate

[Moving from iptables to nftables - nftables wiki](https://wiki.nftables.org/wiki-nftables/index.php/Moving_from_iptables_to_nftables)

```
apt install iptables-nftables-compat

iptables-save > save.txt
iptables-restore-translate -f save.txt
// or
iptables-restore-translate -f save.txt > ruleset.nft
nft -f ruleset.nft
```

# Removing rules

```
nft list table nat -a
       -a, --handle
              Show rule handles in output.

// nft delete rule <family> <table_name> <chain_name> handle <number>
nft delete rule ip nat prerouting handle 7
```

# Incompatibilities

You cannot use iptables and nft to perform NAT at the same time. So make sure
that the iptable_nat module is unloaded: 

```
rmmod iptable_nat
modprobe nft_nat
```

# Netfilter hooks

* [Netfilter hooks - nftables wiki](https://wiki.nftables.org/wiki-nftables/index.php/Netfilter_hooks)



