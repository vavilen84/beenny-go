watch:
	air -c .air.toml

migrate-test-db:
	go run main.go migrate-test-db-up

make-migration:
	go run main.go migration-create $(name)

run-all-tests:
	go test ./... -v -count=1 -failfast -cover -timeout 60m