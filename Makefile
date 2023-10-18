# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)
NOCOLOR_TRIPLE_EQUALS====

# STYLE FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)${1}$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef
define log_action_no_color
$(NOCOLOR_TRIPLE_EQUALS) ${1} $(NOCOLOR_TRIPLE_EQUALS)
endef

gen-wire:
	@echo "$(call log_action_no_color,Generate Wire)"
	wire lib/wire/core/service/auth/wire.go
	wire lib/wire/core/service/employee/wire.go
	wire lib/wire/core/resource/user/wire.go
gen-mock:
	@echo "$(call log_action_no_color,Generate Mock)"
	mockery --all --output=lib/mockery/mocks 

build: gen-wire
	@echo "$(call log_action_no_color,Build Program)"
	go build -o /dist/main cmd/main.go

test-cover: gen-mock
	@echo "$(call log_action,Test Coverage)"
	go test `go list ./... | grep -v mocks` -cover -coverprofile=coverage.out -covermode=count

start-dev: gen-wire
	@echo "$(call log_action,Start Program (Development))"
	go run cmd/main.go -env dev

docker-start:
	@echo "$(call log_action,Start Program (Docker))"
	docker compose --env-file environment/docker.application.env up --build -d
docker-stop:
	@echo "$(call log_action,Stop Program (Docker))"
	docker compose --env-file environment/docker.application.env down
