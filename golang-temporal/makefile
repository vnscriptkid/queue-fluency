up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

### Case 1
run_worker:
	go run ./worker/main.go

run_workflow:
	go run ./workflow/main.go

### Case 2
run_worker_delay:
	go run ./worker_delay/main.go

run_workflow_delay:
	go run ./workflow_delay/main.go

### Case 3
run_worker_activity:
	go run ./worker_signup/main.go

run_workflow_activity:
	go run ./workflow_signup/main.go

### Case 4
run_worker_error:
	go run ./worker_error/main.go

run_workflow_error:
	go run ./workflow_error/main.go

###
psql:
	docker exec -it temporal-db psql -U temporal -d temporal

logs:
	docker compose logs temporal -f

ui:
	open http://localhost:8080