up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

run_worker:
	go run ./worker/main.go

psql:
	docker exec -it temporal-db psql -U temporal -d temporal

logs:
	docker compose logs temporal -f

temporal_ui:
	open http://localhost:8088