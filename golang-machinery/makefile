up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

worker_run:
	go run ./worker/main.go worker

# user: guest / pass: guest
rabbitui:
	open http://localhost:15672