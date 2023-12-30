openssl req -x509 -out localhost.crt -keyout localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=192.168.0.200\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:192.168.0.200\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")