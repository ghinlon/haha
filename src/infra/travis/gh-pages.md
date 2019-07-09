# GitHub Pages Deployment

# Links

* [GitHub Pages Deployment - Travis CI](https://docs.travis-ci.com/user/deployment/pages/)
* [Personal Access Tokens](https://github.com/settings/tokens) 
* [Environment Variables - Travis CI](https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings)

For a minimal configuration, add the following to your `.travis.yml`:

```yaml
deploy:
  provider: pages
  skip-cleanup: true
  github-token: $GITHUB_TOKEN  # Set in the settings page of your repository, as a secure variable
  keep-history: true
  on:
    branch: master
```

# Setting the GitHub token

You’ll need to generate a personal access token with the public_repo or repo scope (repo is required for private repositories). Since the token should be private, you’ll want to pass it to Travis securely in your repository settings or via encrypted variables in .travis.yml.

[Personal Access Tokens](https://github.com/settings/tokens) 

[Environment Variables - Travis CI](https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings)


# 报错

因为这个报错，我整理这个资料

`invalid option "--github-token="`

