# DNS

# Links

# Flushing your DNS cache in Mac OS X and Linux

[Flushing your DNS cache in Mac OS X and Linux – DreamHost](https://help.dreamhost.com/hc/en-us/articles/214981288-Flushing-your-DNS-cache-in-Mac-OS-X-and-Linux)


```
// OS X 12 (Sierra) and later
sudo killall -HUP mDNSResponder;sudo killall mDNSResponderHelper;sudo dscacheutil -flushcache

// rhel6
sudo service nscd restart 
```
# Clear DNS cache in Firefox

`about:config` -> `network.dnsCacheExpiration` and
`network.dnsCacheExpirationGracePeriod` with the values set as 60.

set both to `0` then `right click` then `reset`.



