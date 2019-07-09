# ssh-agent

# Links

* [Generating a new SSH key and adding it to the ssh-agent - User Documentation](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/)
* [SSH 免除重複輸入金鑰密碼教學：SSH Agent 與 Forwarding - G. T. Wang](https://blog.gtwang.org/linux/using-ssh-agent-forwarding-to-avoid-being-asked-passphrase/)
* [linux - Can I automatically add a new host to known_hosts? - Server Fault](https://serverfault.com/questions/132970/can-i-automatically-add-a-new-host-to-known-hosts/316100#316100)
* [Using ssh-agent with ssh](http://mah.everybody.org/docs/ssh)

# Overview

之前想要免密码登陆其它ssh-server, 统都要先把public key拷贝到server里，现在我已经烦这种方法罢。我感觉麻烦。

所以研究一下ssh-agent. 

# HowTo

```sh
    SSH_ENV="$HOME/.ssh/environment"

    function start_agent {
         echo "Initialising new SSH agent..."
         /usr/bin/ssh-agent | sed 's/^echo/#echo/' > "${SSH_ENV}"
         echo succeeded
         chmod 600 "${SSH_ENV}"
         . "${SSH_ENV}" > /dev/null
         /usr/bin/ssh-add;
    }

    # Source SSH settings, if applicable

    if [ -f "${SSH_ENV}" ]; then
         . "${SSH_ENV}" > /dev/null
         #ps ${SSH_AGENT_PID} doesn't work under cywgin
         ps -ef | grep ${SSH_AGENT_PID} | grep ssh-agent$ > /dev/null || {
             start_agent;
         }
    else
         start_agent;
    fi
```

* Run ssh-agent

```sh
eval $(ssh-agent)
```

* Generate a new key

```bash
# ssh-keygen [-q] [-b bits] [-t dsa | ecdsa | ed25519 | rsa] [-N new_passphrase] [-C comment] [-f output_keyfile]
ssh-keygen -t rsa -b 4096 -C "your_email@example.com" -N *** -f .ssh/id20181223_rsa
```

* Adding your SSH key to the ssh-agent

```bash
ssh-add .ssh/id20181223_rsa
ssh-add -l
```
* cfg

`~/.ssh/config`:

```
Host example.com
  ForwardAgent yes
```

or 

`/etc/ssh/ssh_config`:

```
Host *
 ForwardAgent yes
```

or without config, just run `ssh` with `-A` option.

* run

```sh
ssh <host>
# or
ssh -A <host>
```

# Help

```
$ ssh-add --help
ssh-add: illegal option -- -
usage: ssh-add [options] [file ...]
Options:
  -l          List fingerprints of all identities.
  -E hash     Specify hash algorithm used for fingerprints.
  -L          List public key parameters of all identities.
  -k          Load only keys and not certificates.
  -c          Require confirmation to sign using identities
  -m minleft  Maxsign is only changed if less than minleft are left (for XMSS)
  -M maxsign  Maximum number of signatures allowed (for XMSS)
  -t life     Set lifetime (in seconds) when adding identities.
  -d          Delete identity.
  -D          Delete all identities.
  -x          Lock agent.
  -X          Unlock agent.
  -s pkcs11   Add keys from PKCS#11 provider.
  -e pkcs11   Remove keys provided by PKCS#11 provider.
  -q          Be quiet after a successful operation.
```

# Conclusion

研究么半天，终于发觉这个不是我想要的东西。这个实现的是服务器之间免密码，而你本地还是需要把公钥copy到所有服务器上面去的。

# Troubleshooting

其实困扰我的问题是这个：

```
The authenticity of host '[xxxxxxx]:12345 ([123.123.123.123]:12345)' can't be established.
RSA key fingerprint is SHA256:832423werwqer9dsfioWERdafj@#$@%KJSERWe2ewrsd
Are you sure you want to continue connecting (yes/no)? 
```

这种情况就不方便放在脚本里处理罢，不推荐是解决办法是：`ssh -o StrictHostKeyChecking=no username@hostname.com`，这个有安全问题。

正确的方法是：

```sh
#    -R hostname
#                 Removes all keys belonging to hostname from a known_hosts file.  This option is useful to delete hashed hosts (see the -H option above).
ssh-keygen -R [hostname]
ssh-keygen -R [ip_address]
# same as above two
ssh-keygen -R [hostname],[ip_address]

#    -H      Hash all hostnames and addresses in the output.  Hashed names may be used normally by ssh(1) and sshd(8), but they do not reveal identifying informa-
#                 tion should the file's contents be disclosed.

#    -p port
#            Connect to port on the remote host.
ssh-keyscan -H [ip_address] >> ~/.ssh/known_hosts
ssh-keyscan -H [hostname] >> ~/.ssh/known_hosts
# same as above two
ssh-keyscan -H [hostname],[ip_address] >> ~/.ssh/known_hosts
```
