APP_ENV=$(shell echo $$APP_ENV)

.PHONY: migrate-create
migrate-create:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go migrate create; \
	fi

.PHONY: migrate-up
migrate-up:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go migrate up; \
	fi

.PHONY: migrate-down
migrate-down:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go migrate down; \
	fi

.PHONY: migrate-status
migrate-status:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go migrate status; \
	fi

.PHONY: seed
seed:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go seed; \
	fi

.PHONY: test-coverage
test-coverage:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go test $$(go list ./... | grep -v /cmd) -coverprofile=coverage.out && \
		go tool cover -func=coverage.out && \
		go tool cover -html=coverage.out -o coverage.html; \
	fi