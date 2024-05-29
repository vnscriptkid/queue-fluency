up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

run_worker:
	go run ./worker/main.go

run_workflow:
	go run ./workflow/main.go

psql:
	docker exec -it temporal-db psql -U temporal -d temporal

logs:
	docker compose logs temporal -f

temporalui:
	open http://localhost:8080