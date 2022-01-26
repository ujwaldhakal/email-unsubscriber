.PHONY: *

dockerup:
	docker-compose up

publish:
	docker-compose run unsubscriber go run main.go sync-inbox

consume:
	docker-compose run unsubscriber go run main.go parse-email

serve-frontend:
	docker-compose run -p 1323:1323 unsubscriber go run main.go serve-frontend
