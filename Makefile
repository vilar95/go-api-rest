# Makefile para Go API REST

.PHONY: help run build test clean docker-up docker-down migrate install

# Variáveis
APP_NAME=go-api-rest
BUILD_DIR=./bin
MAIN_FILE=cmd/api/main.go

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Instala todas as dependências
	@echo "Instalando dependências..."
	go mod download
	go mod tidy

run: ## Executa a aplicação
	@echo "Executando aplicação..."
	go run $(MAIN_FILE)

build: ## Compila a aplicação
	@echo "Compilando aplicação..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

test: ## Executa os testes
	@echo "Executando testes..."
	go test -v ./...

test-coverage: ## Executa testes com cobertura
	@echo "Executando testes com cobertura..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Remove arquivos de build
	@echo "Limpando arquivos de build..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

docker-up: ## Inicia os containers Docker
	@echo "Iniciando containers..."
	docker-compose up -d

docker-down: ## Para os containers Docker
	@echo "Parando containers..."
	docker-compose down

docker-logs: ## Mostra logs dos containers
	docker-compose logs -f

fmt: ## Formata o código
	@echo "Formatando código..."
	go fmt ./...

lint: ## Executa linter
	@echo "Executando linter..."
	golangci-lint run

tidy: ## Limpa e organiza dependências
	@echo "Organizando dependências..."
	go mod tidy

dev: docker-up run ## Inicia banco e aplicação

all: clean install build ## Limpa, instala dependências e compila
