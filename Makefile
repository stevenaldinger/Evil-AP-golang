# ==================== [START] Global Variable Declaration =================== #
SHELL := /bin/bash
BASE_DIR := $(shell pwd)
UNAME_S := $(shell uname -s)
APP_NAME := evil-twin

export
# ===================== [END] Global Variable Declaration ==================== #

# =========================== [START] Build Scripts ========================== #
build_evil_twin:
	@cd $(BASE_DIR)/cmd/$(APP_NAME) && \
		echo "Building" $(APP_NAME) && \
		CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o $(BASE_DIR)/$(APP_NAME)

build_all: build_evil_twin
# ============================ [END] Build Scripts =========================== #

# ============================ [START] Run Scripts =========================== #
run:
	@$(BASE_DIR)/$(APP_NAME)
# ============================= [END] Run Scripts ============================ #

# ========================= [START] Formatting Script ======================== #
gofmt:
	@go fmt github.com/stevenaldinger/$(APP_NAME)/...

golint:
	@golint github.com/stevenaldinger/$(APP_NAME)/cmd/...
	@golint github.com/stevenaldinger/$(APP_NAME)/internal/...

govet:
	@go vet github.com/stevenaldinger/$(APP_NAME)/cmd/...
	@go vet github.com/stevenaldinger/$(APP_NAME)/internal/...

lint: gofmt golint govet
# ========================== [END] Formatting Script ========================= #

gotest:
	@cd $(BASE_DIR)/cmd/$(APP_NAME) && \
		go test -v -cover
# ======================= [START] Documentation Scripts ====================== #
godoc:
	@godoc -http=":6060"
# ==============-========= [END] Documentation Scripts =========-============= #
