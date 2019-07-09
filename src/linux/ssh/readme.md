# ssh

# show ssh key length

```
ssh-keygen -lf .ssh/id_rsa
```

# change private key passphrase

```
ssh-keygen -p -f ~/.ssh/id_dsa

	-p      Requests changing the passphrase of a private key file instead of creating a new private key. 
```

# copy the public key to a ssh server

```
cat ~/.ssh/id_rsa.pub | ssh <user>@<hostname> 'umask 0077; mkdir -p .ssh; cat >> .ssh/authorized_keys && echo "Key copied"'
```

# Host config

cfg:

`.ssh/config`:

```sh
Host host1
User    user1
HostName host1.com

Host host2
User    user2
HostName host2.com
Port    22322
```

Now can easy `ssh`:

```sh
ssh host1
```


