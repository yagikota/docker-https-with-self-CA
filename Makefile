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
