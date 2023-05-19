
# add self root certificate to Client
.PHONY: add-self-root-certificate
add-self-root-certificate:
	cat client/crt/rootCA.pem >> /etc/ssl/certs/ca-certificates.crt

.PHONY: up-CA
up-CA:
	docker compose up myca -d --build

.PHONY: up-client
up-cleint:
	docker compose up client -d --build

.PHONY: up-server
up-server:
	docker compose up server -d --build

.PHONY: up-client-server
up-client-server: up-cleint up-server
