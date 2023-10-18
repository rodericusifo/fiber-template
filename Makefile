# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)

# FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)$(1)$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef
define check_env
	@if [ -z "$(ENV)" ]; then \
		echo "ERROR: ENV is required. Usage: make $(1) ENV=dev"; \
		exit 1; \
	fi
endef
define set_env
	@export ENV=$(ENV)
endef

# ENV VARIABLE
ENV ?=

gen-wire:
	@echo -e "$(call log_action,Generate Wire)"
	wire lib/wire/core/service/auth/wire.go
	wire lib/wire/core/service/employee/wire.go
	wire lib/wire/core/resource/user/wire.go
	wire lib/wire/core/resource/role/wire.go
	wire lib/wire/core/resource/permission/wire.go
	wire lib/wire/core/resource/role_permission/wire.go
gen-mock:
	@echo -e "$(call log_action,Generate Mock)"
	mockery

build: gen-wire
	@echo -e "$(call log_action,Build Program)"
	go build -o /dist/main cmd/main.go

test-cover: gen-mock
	@echo -e "$(call log_action,Test Coverage)"
	go test `go list ./... | grep -v mocks` -cover -coverprofile=coverage.out -covermode=count

start:
	$(call check_env,start)
	$(call set_env)
	@echo -e "$(call log_action,Start Program ($(ENV)))"
	docker volume ls | grep mysql_fiber_data_$(ENV) || docker volume create --name mysql_fiber_data_$(ENV)
	# docker volume ls | grep postgres_fiber_data_$(ENV) || docker volume create --name postgres_fiber_data_$(ENV)
	docker volume ls | grep redis_fiber_data_$(ENV) || docker volume create --name redis_fiber_data_$(ENV)
	docker network ls | grep fiber_backend_$(ENV) || docker network create fiber_backend_$(ENV)
	docker compose --env-file env/$(ENV).application.env up --build -d

stop:
	$(call check_env,stop)
	$(call set_env)
	@echo -e "$(call log_action,Stop Program ($(ENV)))"
	docker compose --env-file env/$(ENV).application.env down
