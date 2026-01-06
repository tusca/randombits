# Create custom CA for private use

```
#!/usr/bin/env bash
set -euo pipefail
CA_KEY="ca.key"
CA_CERT="ca.crt"
umask 077

if [[ -f "$CA_KEY" || -f "$CA_CERT" ]]; then
    echo "Error: CA files ($CA_KEY or $CA_CERT) already exist in current directory." >&2
    exit 1
fi

openssl ecparam -genkey -name prime256v1 -out "$CA_KEY"

CA_CONFIG=$(mktemp)
cat >"$CA_CONFIG" <<EOF
[ req ]
distinguished_name = dn
x509_extensions = v3_ca
prompt = no
[ dn ]
CN = OpenVPN-CA
[ v3_ca ]
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical, CA:true
keyUsage = critical, digitalSignature, keyCertSign, cRLSign
EOF

openssl req -x509 -new -key "$CA_KEY" -days 36525 -config "$CA_CONFIG" -extensions v3_ca -out "$CA_CERT"
rm -f "$CA_CONFIG"

echo "CA initialized successfully."
echo "  Key:  $CA_KEY"
echo "  Cert: $CA_CERT"
```

# Create cert signed by CA with CommonName CN parameter

Usage: `SCRIPT <CN>`

```
#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "Usage: $0 <CN>" >&2
    exit 1
fi

CN="$1"
KEY_FILE="${CN}.key"
CRT_FILE="${CN}.crt"
CA_KEY="ca.key"
CA_CERT="ca.crt"

umask 077

if [[ ! -f "$CA_KEY" || ! -f "$CA_CERT" ]]; then
    echo "Error: CA files ($CA_KEY, $CA_CERT) not found in current directory." >&2
    echo "Run ca_init.sh first." >&2
    exit 1
fi

if [[ -f "$KEY_FILE" || -f "$CRT_FILE" ]]; then
    echo "Error: Certificate files for '$CN' already exist." >&2
    exit 1
fi

openssl ecparam -genkey -name prime256v1 -out "$KEY_FILE"

CSR_FILE=$(mktemp)
openssl req -new -key "$KEY_FILE" -out "$CSR_FILE" -subj "/CN=${CN}"

EXT_FILE=$(mktemp)
printf "[v3_req]\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth,clientAuth" > "$EXT_FILE"

openssl x509 -req -in "$CSR_FILE" -CA "$CA_CERT" -CAkey "$CA_KEY" -CAcreateserial \
    -out "$CRT_FILE" -days 3650 -sha256 \
    -extensions v3_req -extfile "$EXT_FILE"

rm -f "$CSR_FILE" "$EXT_FILE"

echo "Certificate created successfully."
echo "  Key:  $KEY_FILE"
echo "  Cert: $CRT_FILE"
```

