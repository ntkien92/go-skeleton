APP_ENV=$(shell echo $$APP_ENV)

.PHONY: migrate-create
migrate-create:
	@if [ "$(APP_ENV)" = "local" ]; then \
		go run dev/cmd/main.go migrate create; \
	fi