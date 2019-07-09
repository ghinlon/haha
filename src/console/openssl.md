# openssl

# Links

* [Create a Self-Signed TLS Certificate](https://www.linode.com/docs/security/ssl/create-a-self-signed-tls-certificate/)
* [Self Signed Certificate with Custom Root CA · GitHub](https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309)

# Generate Self-Signed Certificate in one command

```
openssl req -x509 -new -newkey rsa:4096 -keyout MyKey.key -sha256 -days 365 -nodes -out MyCertificate.crt 
chmod 400 /root/certs/MyKey.key
```

# Generate an RSA private key and Self-Signed Certificate in two command

```
openssl genrsa -aes256 -out foo.key 4096
openssl req -x509 -new -key foo.key -sha256 -days 1024 -nodes -out foo.crt

	req       PKCS#10 X.509 Certificate Signing Request (CSR) Management.
```

# Generate CA-Signed Certificate

CSR is used to give it to a CA by you, then CA generate Certificate for you.

```
// you create CSR
openssl req -new -sha256 -key mydomain.com.key -subj "/C=US/ST=CA/O=MyOrg, Inc./CN=mydomain.com" -out mydomain.com.csr
// verify the csr's content
openssl req -in mydomain.com.csr -noout -text

// CA do this
openssl x509 -req -in mydomain.com.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out mydomain.com.crt -days 500 -sha256
// verify the certificate's content
openssl x509 -in mydomain.com.crt -noout -text 

	x509      X.509 Certificate Data Management.
```


