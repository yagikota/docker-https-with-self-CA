# docker-https-with-self-CA

## How to

1. create secret key and CSR at Server

``` shell
cd server/cert
openssl req -nodes -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=JP/ST=Osaka/CN=server"
```


- create self CA

``` shell
cd ..
make up-CA
docker compose exec myca /bin/bash
# folowing: in container(myca)
root@myca:/# cd ~
root@myca:~# pwd
/root
root@myca:~# mkcert -install
Created a new local CA 💥
The local CA is now installed in the system trust store! ⚡️
```

- upload CSR to CA

``` shell
cp server/cert/server-req.pem mkcert/
```

- create a new certificate from CSR valid for Server

``` shell
root@myca:~# mkcert -csr server-req.pem

Created a new certificate valid for the following names 📜
 - "server"

The certificate is at "./server.pem" ✅

It will expire on 16 August 2025 🗓


openssl x509 -in server.pem -text -noout
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            c5:0a:56:83:d2:6a:3c:64:f9:44:fe:b3:ab:ce:e5:5d
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: O = mkcert development CA, OU = root@myca, CN = mkcert root@myca
        Validity
            Not Before: May 16 16:31:34 2023 GMT
            Not After : Aug 16 16:31:34 2025 GMT
        Subject: C = JP, ST = Osaka, L = Paris, CN = server
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (4096 bit)
                Modulus:
                    00:b7:d1:e2:2e:8e:b2:ee:a1:6c:36:09:75:21:50:
                    e4:ef:72:e4:c2:cb:51:5d:27:9b:a5:db:08:3b:a3:
                    4e:a7:76:22:00:a3:95:d4:2b:2d:6d:2f:8e:df:39:
                    11:df:79:94:d7:eb:43:99:dc:b6:d4:bd:af:27:4e:
                    65:68:00:e1:9b:39:fc:f8:5a:e1:61:ab:7a:49:51:
                    fe:1e:29:a4:a3:6b:a5:e9:1e:f8:b6:08:ec:f7:3b:
                    b0:28:31:12:8e:1d:df:c4:fa:bb:da:c4:86:18:a7:
                    f5:32:14:8a:7e:d1:4a:2d:ab:f3:39:60:a9:bd:32:
                    b1:0a:c8:74:74:80:a7:d4:5d:b9:0e:9a:4c:45:4e:
                    65:0b:95:47:a6:ae:c4:a5:8b:35:49:3d:f4:38:61:
                    52:3e:8a:24:15:62:73:eb:48:48:89:fa:46:c2:4d:
                    47:20:bb:91:0c:f9:d6:2a:90:6b:19:75:22:23:01:
                    ea:dd:fd:ef:a6:6e:25:81:af:79:37:fc:f3:4d:14:
                    6c:ce:2c:29:a8:96:0a:9d:91:9d:93:40:3d:79:61:
                    5d:a4:c0:9a:bc:c1:e8:a3:14:74:24:ca:17:48:7d:
                    ad:fb:43:57:cd:3d:64:e4:4f:e2:f8:fe:2b:fa:3b:
                    d5:d6:2e:4c:27:f5:00:cf:ee:e4:ea:31:8f:c0:e4:
                    db:7b:01:4f:58:89:4c:32:fb:4e:28:dd:f1:cc:5f:
                    3a:a6:fb:1b:31:11:18:2b:e8:e0:ba:47:5b:0d:c4:
                    ac:24:3e:7f:cb:44:ed:ce:8f:0a:7c:ee:82:f1:22:
                    d3:bf:12:c2:00:01:a3:05:5a:7c:d9:14:98:24:bb:
                    1e:cf:3e:05:61:73:07:bb:87:b2:cb:b3:22:3a:48:
                    5a:e3:5a:67:f4:49:b3:44:ce:33:a9:f3:6e:59:7a:
                    43:23:b9:b5:5a:23:74:6d:36:b3:13:57:5b:37:ac:
                    24:ef:5b:13:a5:c7:16:d7:0b:9b:5d:f1:0f:83:ca:
                    d9:40:91:52:1a:7f:93:92:ac:a6:19:b6:d9:ae:bb:
                    47:a4:9a:e9:56:f6:5c:e0:28:17:92:01:0e:ab:ee:
                    66:81:58:3b:03:f2:d1:27:39:45:a3:7d:31:f1:6b:
                    28:3c:a4:a8:b1:7c:a3:24:5c:83:17:88:00:1e:f6:
                    77:23:f2:98:64:9d:a3:3b:3a:b3:99:28:06:3b:0b:
                    d8:9a:6b:02:b8:f1:24:f9:e5:55:24:d1:af:8f:d9:
                    de:06:7a:37:ce:44:e0:d3:ae:01:1b:e6:af:26:58:
                    2d:b5:fe:9a:28:5d:2e:d8:74:41:50:ee:56:ac:ca:
                    43:2a:70:e7:5c:d4:b1:fc:a1:cd:b3:6e:bc:62:bf:
                    2c:6d:45
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication
            X509v3 Authority Key Identifier:
                A9:EE:82:7B:46:AA:B5:B7:75:2D:08:46:40:3E:62:BA:EA:FE:CB:70
            X509v3 Subject Alternative Name:
                DNS:server
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        67:1c:d0:fa:cd:e1:33:f6:2e:3a:56:84:76:2f:61:62:0f:83:
        db:bf:9c:a4:28:c7:7d:52:4d:c6:17:82:72:d5:e9:1c:75:61:
        ce:9a:0b:72:69:e7:61:3a:ff:8d:d7:08:2a:41:51:7a:44:b0:
        28:0e:c9:4b:26:be:6d:f5:27:94:09:36:7c:dd:44:c3:8d:a1:
        a5:fc:c4:78:48:4a:47:43:87:d0:d9:43:b0:52:3e:e7:1e:b1:
        17:e5:0c:f3:82:38:48:6d:6d:fe:85:15:57:3b:40:dd:40:35:
        e1:4c:3e:9e:7f:ae:3b:15:fd:f8:c5:b4:de:e9:89:a1:7b:42:
        51:cd:f9:03:bb:9f:cf:e3:cc:65:cb:8b:ac:19:f1:63:83:cc:
        8b:21:75:69:60:ac:0c:bd:ba:56:db:29:53:0c:ed:65:b3:1b:
        e6:76:88:a1:eb:88:f4:17:a4:c3:34:56:a0:ee:ea:26:8b:e7:
        f5:1b:9c:bc:37:3f:a8:a4:4c:37:3a:dd:88:3a:f4:8f:b3:3f:
        a5:5d:61:a1:9c:ef:ea:55:3d:f5:a5:e4:23:10:0b:68:c8:a1:
        a8:a5:8f:d1:4a:24:bc:39:d9:8b:49:ae:6b:73:dd:59:ef:2b:
        85:d8:f8:5c:31:b2:b7:8e:e9:76:e4:71:ad:72:9d:14:93:15:
        60:d5:30:9b:c3:3e:c4:7c:33:29:04:5d:ad:71:a9:86:d1:04:
        b1:cd:88:2f:4e:8e:7a:be:fa:68:f3:e2:95:86:6b:92:9f:df:
        1b:ed:b9:7e:d2:26:81:f7:c4:af:5b:ce:15:4f:39:68:7b:cc:
        ff:5e:2a:35:52:44:87:62:e1:38:7c:a2:8d:99:ef:74:90:b3:
        3a:2f:2a:c0:d3:6e:58:97:3a:63:94:76:0a:c2:88:32:eb:96:
        09:7f:c6:c7:57:c4:05:4e:54:36:a9:f2:04:ec:f2:00:51:2a:
        27:fe:e1:10:45:25:ff:eb:dc:8a:a2:25:db:cf:2b:dc:73:41:
        34:86:8c:ec:f4:f3
```

``` shell
cp mkcert/server.pem server/cert
```

- add rootCA certificate to Client

``` shell
# docker compose exec client sh -c "cp client/cert/rootCA.pem >> /usr/local/share/ca-certificates/rootCA.pem"
docker compose exec client sh -c "cp client/cert/rootCA.pem >> /etc/ssl/certs/rootCA.pem"

```
