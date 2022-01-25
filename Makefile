.PHONY: *

dockerup:
	docker-compose up

publish:
	docker-compose run unsubscriber go run main.go sync-inbox

consume:
	docker-compose run unsubscriber go run main.go parse-email

asyncConsume:
	make -j consume
