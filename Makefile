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

.PHONY: start-server-packet-capture
start-server-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d server timeout 180  tcpdump -i any -w "/captured/server/${DATE}.pcap"
