# YUM

* Links

* [8.4.6. Creating a Yum Repository - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/deployment_guide/sec-yum_repository)

# Creating a Yum Repository

```
yum install createrepo
createrepo --database /mnt/local_repo
```

# Working with Package Groups


```
yum groups summary
yum group info glob_expression…

yum group install "group name"
yum group install groupid
// same as group install
yum install @group
// Replace group with the groupid or quoted group name. 
yum install @^group

yum group remove group_name
yum group remove groupid
yum remove @group
yum remove @^group
```
