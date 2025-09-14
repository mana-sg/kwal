BINARY_NAME = kwal
TARGET = main.go
INSTALL_DIR = /usr/local/bin
CONFIG_DIR = $(HOME)/kwal

build:
	@echo "🔨 Building binary..."
	@bash -c '\
		tput civis; \
		trap "tput cnorm; exit" INT TERM EXIT; \
		( \
			while true; do \
				for c in / - \\ \|; do \
					printf "\r[%s] $$c" "Building"; \
					sleep 0.1; \
				done \
			done \
		) & \
		SPIN_PID=$$!; \
		go build -o $(BINARY_NAME) $(TARGET); \
		kill $$SPIN_PID >/dev/null 2>&1; \
		wait $$SPIN_PID 2>/dev/null || true; \
		tput cnorm; \
		printf "\r[Building] ✓\n"'

install: build
	@echo "🚀 Setting up structure..."
	@bash -c '\
		tput civis; \
		trap "tput cnorm; exit" INT TERM EXIT; \
		( \
			while true; do \
				for c in / - \\ \|; do \
					printf "\r[%s] $$c" "Installing"; \
					sleep 0.1; \
				done \
			done \
		) & \
		SPIN_PID=$$!; \
		mkdir -p $(CONFIG_DIR) && \
		touch $(CONFIG_DIR)/log.bin && \
		chmod +x $(CONFIG_DIR)/log.bin && \
		sudo mv $(BINARY_NAME) $(INSTALL_DIR); \
		kill $$SPIN_PID >/dev/null 2>&1; \
		wait $$SPIN_PID 2>/dev/null || true; \
		tput cnorm; \
		printf "\r[Installing] ✓\n"'

uninstall:
	@echo "🗑️  Deleting records..."
	@sudo rm -rf $(CONFIG_DIR)
	@echo "🧽 Removing $(BINARY_NAME) from $(INSTALL_DIR)..."
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "✅ Uninstallation complete!"

clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "✅ Cleanup complete!"
