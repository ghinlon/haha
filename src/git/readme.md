# Git

# Links

* [git - the simple guide - no deep shit!](https://rogerdudler.github.io/git-guide/)
* [Connecting to GitHub with SSH - User Documentation](https://help.github.com/articles/connecting-to-github-with-ssh/)
* [Changing a remote's URL - User Documentation](https://help.github.com/articles/changing-a-remote-s-url/#switching-remote-urls-from-https-to-ssh)

# ssh 

**Tip:** with ssh,`http_proxy` in `git config` won't work.may use `proxychains` instead.

```sh
git remote set-url origin git@github.com:USERNAME/REPOSITORY.git
// Verify new remote URL
git remote -v
```

# Delete All Commits History

**Links:** [GitHub - Delete commits history with git commands](https://gist.github.com/heiswayi/350e2afda8cece810c0f6116dadbe651)

```sh
rm -rf .git

git init
git remote add origin git@github.com:USERNAME/REPOSITORY.git
git remote -v

git add --all
git commit -am "Initial commit"
git push -f origin master
```

# Create a new branch with git and manage branches

**Links:** [Create a new branch with git and manage branches · Kunena/Kunena-Forum Wiki · GitHub](https://github.com/Kunena/Kunena-Forum/wiki/Create-a-new-branch-with-git-and-manage-branches)

```sh
// list
git branch

// create at local
git checkout -b [name_of_your_new_branch]

// Change working branch
git checkout [name_of_your_new_branch]

// Push the branch on remote  
git push origin [name_of_your_new_branch]

// create at remote  
git remote add [name_of_your_remote] [name_of_your_new_branch]

// Push changes from your commit into your branch
git push [name_of_your_new_remote] [url]

// Update your branch when the original branch from official repository has been updated 
git fetch [name_of_your_remote]

// Then you need to apply to merge changes, if your branch is derivated from develop you need to do :
git merge [name_of_your_remote]/develop

// delete at local  
git branch -d <branch_name>

// To force the deletion of local branch on your filesystem :
git branch -D [name_of_your_new_branch]

// delete at remote  
git push -d <remote_name> <branch_name>
```

# Delete File From Git Repo

**Links:** [git rm - How can I delete a file from git repo? - Stack Overflow](https://stackoverflow.com/a/2047477)

```sh
git rm file1.txt
git commit -m "remove file1.txt"

// But if you want to remove the file only from the Git repository and not remove it from the filesystem, use:
git rm --cached file1.txt
git commit -m "remove file1.txt"

// And to push changes to remote repo
git push origin branch_name
```

# Merge With Local Modifications

**Links:** [How do I resolve git saying "Commit your changes or stash them before you can merge"? - Stack Overflow](https://stackoverflow.com/a/15745424)

You can't merge with local modifications. Git protects you from losing potentially important changes.
You have three options: 

* Stash

	```sh
	git commit -m "My message"
	// Stashing acts as a stack, where you can push changes, and you pop them in reverse order.
	git stash
	// Do the merge, and then pull the stash:
	git stash pop
	```
* Discard the local changes

	```sh
	git reset --hard
	// or
	git checkout -t -f remote/branch
	```
* Discard local changes for a specific file

	```sh
	git checkout filename
	```
# replace with ssh

`git config --global url."git@git.xxx.com:".insteadOf "https://git.xxx.com/"`
