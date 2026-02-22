up:
	docker-compose up --build -d

down:
	docker-compose down

logs:
	docker-compose logs -f

test:
	cd backend && go test ./... -v
	cd matching-service && cargo test
