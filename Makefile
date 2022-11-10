OUTPUT_DIR = out

build:
	mkdir -p "$(OUTPUT_DIR)"
	go build -o "$(OUTPUT_DIR)/ocm-toolbox" "cmd/main.go"
