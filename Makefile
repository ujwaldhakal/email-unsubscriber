.PHONY: *

parse-email:
	docker-compose run unsubscriber go run main.go parse-email
