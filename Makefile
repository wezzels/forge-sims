# FORGE-Sims Makefile
# Build system for all simulator binaries

GO := go
GOFLAGS := -ldflags="-s -w"  # Strip debug info for smaller binaries
BUILD_DIR := build/linux-x86
CMD_DIR := cmd

# All simulators (in cmd/ directory)
SIMULATORS := $(shell find $(CMD_DIR) -maxdepth 1 -type d -not -name 'cmd' | sed 's|cmd/||')

.PHONY: all build clean test verify

all: build

# Build all simulators
build:
	@echo "Building FORGE-Sims..."
	@mkdir -p $(BUILD_DIR)
	@for sim in $(SIMULATORS); do \
		echo "  [BUILD] $$sim"; \
		$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$$sim ./$(CMD_DIR)/$$sim; \
	done
	@echo "Done. Binaries in $(BUILD_DIR)/"

# Build specific simulator
$(BUILD_DIR)/%: $(CMD_DIR)/%/main.go
	@mkdir -p $(BUILD_DIR)
	@echo "  [BUILD] $*"
	$(GO) build $(GOFLAGS) -o $@ ./$(CMD_DIR)/$*

# Launch vehicle simulator (standalone target)
launch-veh-sim:
	@mkdir -p $(BUILD_DIR)
	@echo "  [BUILD] launch-veh-sim"
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/launch-veh-sim ./$(CMD_DIR)/launch-veh-sim

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f *.test

# Run verification suite
verify:
	cd verification && $(GO) run verify_suite.go

# Run specific simulator
run-%:
	@if [ -f $(BUILD_DIR)/$* ]; then \
		$(BUILD_DIR)/$*; \
	else \
		echo "Binary $* not found. Run 'make build' first."; \
		exit 1; \
	fi

# List available simulators
list:
	@echo "Available simulators:"
	@for sim in $(SIMULATORS); do \
		echo "  - $$sim"; \
	done
	@echo ""
	@echo "Built binaries:"
	@for bin in $(BUILD_DIR)/*; do \
		if [ -f "$$bin" ]; then \
			size=$$(du -h "$$bin" | cut -f1); \
			echo "  - $$(basename $$bin) ($$size)"; \
		fi; \
	done

# Cross-compile for multiple platforms
cross:
	@mkdir -p build/linux-x86 build/linux-arm64 build/darwin-amd64 build/darwin-arm64
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o build/linux-x86/ ./...
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) -o build/linux-arm64/ ./...
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -o build/darwin-amd64/ ./...
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -o build/darwin-arm64/ ./...

# Format code
fmt:
	$(GO) fmt ./...

# Run tests
test:
	$(GO) test ./...

# Update dependencies
tidy:
	$(GO) mod tidy

# Help
help:
	@echo "FORGE-Sims Build System"
	@echo ""
	@echo "Targets:"
	@echo "  all          Build all simulators (default)"
	@echo "  build        Build all simulators"
	@echo "  clean        Remove build artifacts"
	@echo "  verify       Run verification suite"
	@echo "  test         Run tests"
	@echo "  fmt          Format code"
	@echo "  tidy         Update go.mod"
	@echo "  cross        Cross-compile for all platforms"
	@echo "  list         List available simulators"
	@echo "  help         Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build               # Build all"
	@echo "  make launch-veh-sim      # Build specific simulator"
	@echo "  make run-launch-veh-sim  # Build and run"
	@echo "  make verify              # Run verification suite"