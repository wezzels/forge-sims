# FORGE-Sims Makefile
# Build system for simulator binaries

# Binary directory
BUILD_DIR := build/linux-x86
BIN_DIR := binaries/linux-x86

# Pre-built simulators (from binaries/)
SIMULATORS := $(notdir $(wildcard $(BIN_DIR)/*))

.PHONY: all build clean verify list help

all: build

# Build: copy pre-built binaries to build/
build:
	@echo "Setting up FORGE-Sims..."
	@mkdir -p $(BUILD_DIR)
	@for sim in $(SIMULATORS); do \
		cp $(BIN_DIR)/$$sim $(BUILD_DIR)/$$sim 2>/dev/null || true; \
	done
	@echo "Built binaries in $(BUILD_DIR)/"

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Run verification suite
verify:
	cd verification && go run verify_suite.go

# List available simulators
list:
	@echo "Available simulators:"
	@for sim in $(SIMULATORS); do \
		size=$$(du -h "$(BIN_DIR)/$$sim" 2>/dev/null | cut -f1 || echo "?"); \
		echo "  - $$sim ($$size)"; \
	done

# Help
help:
	@echo "FORGE-Sims Build System"
	@echo ""
	@echo "Targets:"
	@echo "  all          Setup build directory (default)"
	@echo "  build        Copy binaries to build/"
	@echo "  clean        Remove build artifacts"
	@echo "  verify       Run verification suite"
	@echo "  list         List available simulators"
	@echo "  help         Show this help"
	@echo ""
	@echo "Note: Source code for launch-veh-sim is at:"
	@echo "      https://idm.wezzel.com/crab-meat-repos/launch-veh-sim"
