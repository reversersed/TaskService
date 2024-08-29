general:
	@echo === Available makefile commands:
	@echo run - run the service
	@echo install (alias - i) - install project dependencies
	@echo gen - generate documentation and mock files
	@echo upgrade - upgrade dependencies
	@echo clean - clean mod files
	@echo start - start docker compose with rebuild
	@echo up - start docker compose without rebuild
	@echo stop - stop docker container

run: gen start
	
i: install

install:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(MAKE) clean

gen:
	@swag init --parseDependency -d ./internal/endpoint -g ../app/app.go -o ./docs

upgrade: clean i
	@go get -u ./... && go mod tidy

clean:
	@go mod tidy

start:
	@docker compose --env-file ./config/.env up --build --timestamps --wait --wait-timeout 1800 --remove-orphans -d

stop:
	@docker compose --env-file ./config/.env stop

up:
	@docker compose --env-file ./config/.env up --timestamps --wait --wait-timeout 1800 --remove-orphans -d