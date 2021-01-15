# Handy aliases

```
alias readcrt="openssl x509 -inform PEM -text -in"
alias readcsr="openssl req -text -noout -verify -in"
alias readpk="openssl rsa -check -in"
```

used as
```
readcrt file.crt.pem
```

# End Date ... just need the end date

```
openssl x509 -enddate -noout -in client.pem
```

# Check crt from website

```
echo | openssl s_client -showcerts -servername FQN -connect FQN_OR_IP:443 2>/dev/null | openssl x509 -inform pem -noout -text
echo | openssl s_client -showcerts -servername FQN -connect FQN_OR_IP:443 2>/dev/null | openssl x509 -inform pem -noout -enddate
```

where:
- FQN_OR_IP : can be the fully qualified name (eg. x.domain.com or x.x.x.x)
- FQN : fully qualified name
- 443 : if needed change port

# More OpenSSL Checks

https://langui.sh/2009/03/14/checking-a-remote-certificate-chain-with-openssl/

# Generate using available CA

```
# KEY
openssl genrsa -out NAME.key -aes256 -passout pass:PASSWORD 4096
# REQUEST
openssl req -new -key NAME.key -out NAME.csr -subj "/C=?/ST=?/L=?/O=?/OU=?/CN=NAME/emailAddress=name@domain.com" -passin pass:PASSWORD
# SIGN and GEN
openssl x509 -req -in NAME.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out deviceX.crt -days 3650 -sha256
```

- replacing `NAME` and `/C=?/ST=?/L=?/O=?/OU=?/CN=NAME/emailAddress=name@domain.com`
- replace 3650 if needed, this is 10 years validity

# Remove password from private key

```
openssl rsa -in source.key -out destination.key -passin pass:ZE_PASSWORD
```

# Print System certs

```
awk -v cmd='openssl x509 -noout -subject' '
    /BEGIN/{close(cmd)};{print | cmd}' < /etc/ssl/certs/ca-certificates.crt
```

# Convert from DER to PEM

```
# key
function keyder2pem()  {
   openssl rsa -inform der -outform pem -in $1 -out $1.pem
}
 
 
# cert
function crtder2pem() {
   openssl x509 -inform der -outform pem -in $1 -out $1.pem;
}
```

# Keytores / JKS

```
keytool -list -keystore cacerts
```

and import

```
keytool -import -alias NAME -file NAME.cer -keystore /srv/java/jre/lib/security/cacerts -storepass changeit  -noprompt
```

more at https://www.sslshopper.com/article-most-common-java-keytool-keystore-commands.html

