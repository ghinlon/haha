# [The GNU Privacy Guard](https://gnupg.org/)

# Links

* [GnuPG - ArchWiki](https://wiki.archlinux.org/index.php/GnuPG)

# Install

```
# Arch
pacman -S gnupg

# Mac
brew install gnupg
```

cfg:

By default, the gnupg directory has its permissions set to 700 and the files it contains have their permissions set to 600. 

CfgFile: `.gnupg/gpg.conf`

```
default-key <secret-key>
default-recipient-self
```

# decrypt

```
$ gpg -o <file> -d foo.gpg

--output file
-o file
      Write output to file.  To write to stdout use - as the filename.

--decrypt
      -d  Decrypt the file given on the command line (or STDIN if no file is specified) and write it to STDOUT (or the file specified with --output). If the
      decrypted file is signed, the signature is also verified. This command differs from the default operation, as it never writes to the filename which is
      included in the file and it rejects files that don't begin with an encrypted message.
```

# encrypt

```
$ gpg -r <user-id> -e <data>
$ gpg -c <data>

--recipient name
      -r Encrypt for user id name. If this option or --hidden-recipient is not specified, GnuPG asks for the user-id unless --default-recipient is given.

--encrypt
      -e  Encrypt  data to one or more public keys. This command may be combined with --sign (to sign and encrypt a message), --symmetric (to encrypt a mes-
      sage that can decrypted using a secret key or a passphrase), or --sign and --symmetric together (for a signed message that can be  decrypted  using  a
      secret key or a passphrase).  --recipient and related options specify which public keys to use for encryption.

--symmetric
      -c  Encrypt with a symmetric cipher using a passphrase. The default symmetric cipher used is AES-128, but may be chosen with the --cipher-algo option.
      This command may be combined with --sign (for a signed and symmetrically encrypted message), --encrypt (for a message that  may  be  decrypted  via  a
      secret  key  or  a  passphrase),  or --sign and --encrypt together (for a signed message that may be decrypted via a secret key or a passphrase).  gpg
      caches the passphrase used for symmetric encryption so that a decrypt operation may not require that the user needs  to  enter  the  passphrase.   The
      option --no-symkey-cache can be used to disable this feature.
```

# Key maintenance

* backup secret key

```
$ gpg --export-secret-keys --armor <user-id> > privkey.asc
```

* import secret key, same as import public key

```
$ gpg --import privkey.asc
$ gpg --import public.key
```

* export public key

```
$ gpg -o public.key -a --export <user-id>

--armor
	     -a Create ASCII armored output.  The default is to create the binary OpenPGP format.
```

# list keys

```
gpg --list-keys
gpg --list-secret-keys
```

# create a key pair

```
$ gpg --full-gen-key

--full-generate-key
	     --full-gen-key Generate a new key pair with dialogs for all options.  This is an extended version of --generate-key.
```

# Change Passphrase of the Private Key

```
gpg --list-keys
gpg --edit-key your-key-ID
	
gpg> passwd
gpg> save
```
