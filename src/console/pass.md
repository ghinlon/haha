# [pass](https://www.passwordstore.org/)

# Links

* [Pass: The Standard Unix Password Manager](https://www.passwordstore.org/)

`pass` makes managing these individual password files extremely easy. All
passwords live in `~/.password-store`, and pass provides some nice commands for
adding, editing, generating, and retrieving passwords.


# Install

```
brew install pass
```

# First Time

```
pass init <gpg-id>
pass git init
```

# Using the password store

```
pass						// list all
pass Email/zx2c4.com				// show
pass -c Email/zx2c4.com				// copy to clipboard
pass insert Business/account			// insert
pass edit pass-name				// edit passwords
pass generate Email/jasondonenfeld.com 15	// generate password
						// It's possible to generate
// passwords with no symbols using --no-symbols or -n 
pass rm Business/cheese-whiz-factory		// remove

pass find regrex				// find
```

# Data Organization

One approach is to use the multi-line functionality of pass (`--multiline` or
`-m` in `insert`), and store the password itself on the first line of the file,
and the additional information on subsequent lines. For example,
`Amazon/bookreader` might look like this:

```
Yw|ZSNH!}z"6{ym9pI
URL: *.amazon.com/*
Username: AmazonianChicken@example.com
Secret Question 1: What is your childhood best friend's most bizarre superhero fantasy? Oh god, Amazon, it's too awful to say...
Phone Support PIN #: 84719
```

# Import From Keepass

[roddhjav/pass-import: A pass extension for importing data from most of the existing password manager.](https://github.com/roddhjav/pass-import#readme)



