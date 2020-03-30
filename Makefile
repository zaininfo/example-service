.PHONY: start
start:
	@ docker-compose up -d

.PHONY: stop
stop:
	@ docker-compose down -v --rmi local

.PHONY: start-db
start-db:
	@ docker-compose start redis

.PHONY: stop-db
stop-db:
	@ docker-compose stop redis
