up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

run_worker:
	go run ./worker/main.go

run_workflow:
	go run ./workflow/main.go

run_worker_delay:
	go run ./worker_delay/main.go

run_workflow_delay:
	go run ./workflow_delay/main.go

run_worker_activity:
	go run ./worker_signup/main.go

run_workflow_activity:
	go run ./workflow_signup/main.go

psql:
	docker exec -it temporal-db psql -U temporal -d temporal

logs:
	docker compose logs temporal -f

temporalui:
	open http://localhost:8080