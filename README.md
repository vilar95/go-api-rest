# Go API REST - Guia Completo de Desenvolvimento

## 📚 Índice

1. [Introdução](#introdução)
2. [Conceitos e Arquitetura](#conceitos-e-arquitetura)
3. [Pré-requisitos](#pré-requisitos)
4. [Passo a Passo Completo](#passo-a-passo-completo)
5. [Estrutura do Projeto](#estrutura-do-projeto)
6. [Como Usar](#como-usar)
7. [Conceitos Avançados](#conceitos-avançados)

---

## 📖 Introdução

Este é um guia completo para construir uma API REST em Go utilizando as melhores práticas de desenvolvimento, incluindo:

- **Clean Architecture** (Arquitetura Limpa)
- **Separação de responsabilidades em camadas**
- **Princípios SOLID**
- **Injeção de dependências**
- **Validação de dados**
- **Tratamento de erros**
- **Logging estruturado**
- **Middlewares**

---

## 🏗️ Conceitos e Arquitetura

### Clean Architecture (Arquitetura em Camadas)

A aplicação está organizada em camadas com responsabilidades bem definidas:

```
┌─────────────────────────────────────────┐
│         Handler (Presentation)          │  ← Interface HTTP (recebe requests)
├─────────────────────────────────────────┤
│           Service (Business)            │  ← Lógica de negócio
├─────────────────────────────────────────┤
│        Repository (Data Access)         │  ← Acesso aos dados
├─────────────────────────────────────────┤
│              Database                   │  ← Banco de dados
└─────────────────────────────────────────┘
```

**Benefícios:**
- ✅ Código testável e manutenível
- ✅ Baixo acoplamento entre camadas
- ✅ Fácil substituição de componentes
- ✅ Reutilização de código

### Princípios SOLID Aplicados

1. **S** - Single Responsibility: Cada struct tem uma única responsabilidade
2. **O** - Open/Closed: Extensível sem modificar código existente
3. **L** - Liskov Substitution: Interfaces podem ser substituídas
4. **I** - Interface Segregation: Interfaces específicas por necessidade
5. **D** - Dependency Inversion: Dependências invertidas através de interfaces

---

## 🔧 Pré-requisitos

Antes de começar, instale:

- **Go 1.21+** - [Download](https://go.dev/dl/)
- **PostgreSQL 14+** - [Download](https://www.postgresql.org/download/)
- **Docker & Docker Compose** (opcional) - [Download](https://www.docker.com/)
- **Git** - [Download](https://git-scm.com/)
- Um editor de código (**VS Code** recomendado)

---

## 🚀 Passo a Passo Completo

### **FASE 1: Configuração Inicial do Projeto**

#### Passo 1.1: Criar o Diretório do Projeto

```bash
# Crie o diretório do projeto
mkdir go-api-rest
cd go-api-rest
```

#### Passo 1.2: Inicializar o Módulo Go

```bash
# Inicialize o módulo Go
go mod init go-api-rest
```

**O que acontece aqui?**
- Cria o arquivo `go.mod` que gerencia as dependências do projeto
- Define o nome do módulo (usado nos imports)

#### Passo 1.3: Criar a Estrutura de Pastas

```bash
# Crie todas as pastas necessárias
mkdir -p cmd/api
mkdir -p internal/config
mkdir -p internal/dto
mkdir -p internal/handler
mkdir -p internal/middleware
mkdir -p internal/repository
mkdir -p internal/router
mkdir -p internal/service
mkdir -p database
mkdir -p models
mkdir -p pkg/logger
mkdir -p pkg/response
mkdir -p pkg/validator
mkdir -p migration
```

**Explicação da estrutura:**
- `cmd/api/` - Ponto de entrada da aplicação
- `internal/` - Código privado da aplicação (não pode ser importado por outros projetos)
- `models/` - Modelos de dados (structs do banco)
- `pkg/` - Código reutilizável (pode ser importado por outros projetos)
- `database/` - Configuração de conexão com banco
- `migration/` - Scripts SQL de migração

---

### **FASE 2: Configuração do Banco de Dados**

#### Passo 2.1: Instalar Dependências do Banco

```bash
# GORM - ORM para Go
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

**Por que GORM?**
- ORM completo e maduro
- Suporte a migrations automáticas
- Query builder intuitivo
- Relacionamentos facilitados

#### Passo 2.2: Criar Conexão com Banco (`database/db.go`)

```go
package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

// NewDatabase cria uma nova conexão com o banco de dados
func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com o banco: %w", err)
	}

	return &Database{DB: db}, nil
}
```

**Conceitos importantes:**
- `dsn` (Data Source Name): string de conexão com o banco
- `gorm.Config`: configurações do GORM (logs, etc)
- Retornamos um ponteiro para `Database` e um `error` (padrão Go)

#### Passo 2.3: Criar Sistema de Configuração (`internal/config/config.go`)

```go
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Load carrega as configurações das variáveis de ambiente
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 8000),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

// GetDSN retorna a string de conexão do banco de dados
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// Funções auxiliares
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Erro ao converter %s para int, usando valor padrão: %d", key, defaultValue)
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}
```

**Por que configuração separada?**
- ✅ Facilita mudança entre ambientes (dev, prod)
- ✅ Não expõe credenciais no código
- ✅ Usa variáveis de ambiente (12 Factor App)

---

### **FASE 3: Criar o Modelo de Dados**

#### Passo 3.1: Criar Model (`models/personality.go`)

```go
package models

import "time"

// Personality representa uma personalidade histórica
type Personality struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null;size:100"`
	History   string    `json:"history" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName especifica o nome da tabela
func (Personality) TableName() string {
	return "personalities"
}
```

**Tags explicadas:**
- `json:"id"` - Nome do campo no JSON
- `gorm:"primaryKey"` - Chave primária
- `gorm:"unique"` - Valor único no banco
- `gorm:"not null"` - Campo obrigatório
- `gorm:"autoCreateTime"` - Preenche automaticamente na criação

---

### **FASE 4: Camada de Repository (Acesso aos Dados)**

#### Passo 4.1: Criar Interface e Implementação (`internal/repository/personality_repository.go`)

```go
package repository

import (
	"go-api-rest/models"
	"gorm.io/gorm"
)

// PersonalityRepository define o contrato para acesso aos dados
type PersonalityRepository interface {
	Create(personality *models.Personality) error
	GetAll() ([]models.Personality, error)
	GetByID(id uint) (*models.Personality, error)
	Update(personality *models.Personality) error
	Delete(id uint) error
}

// personalityRepository implementa PersonalityRepository
type personalityRepository struct {
	db *gorm.DB
}

// NewPersonalityRepository cria uma nova instância do repository
func NewPersonalityRepository(db *gorm.DB) PersonalityRepository {
	return &personalityRepository{db: db}
}

func (r *personalityRepository) Create(personality *models.Personality) error {
	return r.db.Create(personality).Error
}

func (r *personalityRepository) GetAll() ([]models.Personality, error) {
	var personalities []models.Personality
	err := r.db.Find(&personalities).Error
	return personalities, err
}

func (r *personalityRepository) GetByID(id uint) (*models.Personality, error) {
	var personality models.Personality
	err := r.db.First(&personality, id).Error
	return &personality, err
}

func (r *personalityRepository) Update(personality *models.Personality) error {
	return r.db.Save(personality).Error
}

func (r *personalityRepository) Delete(id uint) error {
	return r.db.Delete(&models.Personality{}, id).Error
}
```

**Por que usar interfaces?**
- ✅ Facilita testes (mock/stub)
- ✅ Desacopla implementação
- ✅ Permite trocar banco de dados facilmente

---

### **FASE 5: Camada de Service (Lógica de Negócio)**

#### Passo 5.1: Criar DTOs (`internal/dto/personality_dto.go`)

```go
package dto

// CreatePersonalityRequest representa os dados para criar uma personalidade
type CreatePersonalityRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=100"`
	History string `json:"history" validate:"required,min=10"`
}

// UpdatePersonalityRequest representa os dados para atualizar uma personalidade
type UpdatePersonalityRequest struct {
	Name    string `json:"name" validate:"omitempty,min=3,max=100"`
	History string `json:"history" validate:"omitempty,min=10"`
}
```

**O que são DTOs?**
- Data Transfer Objects
- Separam dados da API dos modelos do banco
- Permitem validações específicas por endpoint

#### Passo 5.2: Criar Service (`internal/service/personality_service.go`)

```go
package service

import (
	"errors"
	"go-api-rest/internal/dto"
	"go-api-rest/internal/repository"
	"go-api-rest/models"
	"gorm.io/gorm"
)

type PersonalityService interface {
	Create(req *dto.CreatePersonalityRequest) (*models.Personality, error)
	GetAll() ([]models.Personality, error)
	GetByID(id uint) (*models.Personality, error)
	Update(id uint, req *dto.UpdatePersonalityRequest) (*models.Personality, error)
	Delete(id uint) error
}

type personalityService struct {
	repo repository.PersonalityRepository
}

func NewPersonalityService(repo repository.PersonalityRepository) PersonalityService {
	return &personalityService{repo: repo}
}

func (s *personalityService) Create(req *dto.CreatePersonalityRequest) (*models.Personality, error) {
	personality := &models.Personality{
		Name:    req.Name,
		History: req.History,
	}

	if err := s.repo.Create(personality); err != nil {
		return nil, err
	}

	return personality, nil
}

func (s *personalityService) GetAll() ([]models.Personality, error) {
	return s.repo.GetAll()
}

func (s *personalityService) GetByID(id uint) (*models.Personality, error) {
	personality, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("personalidade não encontrada")
		}
		return nil, err
	}
	return personality, nil
}

func (s *personalityService) Update(id uint, req *dto.UpdatePersonalityRequest) (*models.Personality, error) {
	personality, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		personality.Name = req.Name
	}
	if req.History != "" {
		personality.History = req.History
	}

	if err := s.repo.Update(personality); err != nil {
		return nil, err
	}

	return personality, nil
}

func (s *personalityService) Delete(id uint) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
```

**Responsabilidades do Service:**
- ✅ Validações de negócio
- ✅ Transformações de dados (DTO → Model)
- ✅ Orquestração de múltiplos repositories
- ✅ Tratamento de erros específicos

---

### **FASE 6: Pacotes Auxiliares**

#### Passo 6.1: Instalar Dependência de Validação

```bash
go get github.com/go-playground/validator/v10
```

#### Passo 6.2: Criar Validador (`pkg/validator/validator.go`)

```go
package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate valida uma struct com base nas tags
func Validate(data interface{}) error {
	if err := validate.Struct(data); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("erro de validação: %s", validationErrors[0].Field())
		}
		return err
	}
	return nil
}
```

#### Passo 6.3: Criar Response Padronizada (`pkg/response/response.go`)

```go
package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JSON envia uma resposta JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Success envia uma resposta de sucesso
func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	JSON(w, statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error envia uma resposta de erro
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, Response{
		Success: false,
		Error:   message,
	})
}
```

#### Passo 6.4: Criar Logger (`pkg/logger/logger.go`)

```go
package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	infoLogger.Println(message)
}

func Infof(format string, v ...interface{}) {
	infoLogger.Println(fmt.Sprintf(format, v...))
}

func Error(message string) {
	errorLogger.Println(message)
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Println(fmt.Sprintf(format, v...))
}
```

---

### **FASE 7: Camada de Handler (HTTP)**

#### Passo 7.1: Instalar Roteador

```bash
go get -u github.com/gorilla/mux
```

#### Passo 7.2: Criar Handler (`internal/handler/personality_handler.go`)

```go
package handler

import (
	"encoding/json"
	"go-api-rest/internal/dto"
	"go-api-rest/internal/service"
	"go-api-rest/pkg/logger"
	"go-api-rest/pkg/response"
	customValidator "go-api-rest/pkg/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonalityHandler struct {
	service service.PersonalityService
}

func NewPersonalityHandler(service service.PersonalityService) *PersonalityHandler {
	return &PersonalityHandler{service: service}
}

// Home retorna uma mensagem de boas-vindas
func (h *PersonalityHandler) Home(w http.ResponseWriter, r *http.Request) {
	response.Success(w, http.StatusOK, "API de Personalidades - Bem-vindo!", map[string]string{
		"version": "1.0.0",
		"status":  "online",
	})
}

// Create cria uma nova personalidade
func (h *PersonalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePersonalityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := customValidator.Validate(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	personality, err := h.service.Create(&req)
	if err != nil {
		logger.Errorf("Erro ao criar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao criar personalidade")
		return
	}

	response.Success(w, http.StatusCreated, "Personalidade criada com sucesso", personality)
}

// GetAll retorna todas as personalidades
func (h *PersonalityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	personalities, err := h.service.GetAll()
	if err != nil {
		logger.Errorf("Erro ao buscar personalidades: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao buscar personalidades")
		return
	}

	response.Success(w, http.StatusOK, "Personalidades recuperadas com sucesso", personalities)
}

// GetByID retorna uma personalidade por ID
func (h *PersonalityHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	personality, err := h.service.GetByID(uint(id))
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Personalidade encontrada", personality)
}

// Update atualiza uma personalidade
func (h *PersonalityHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var req dto.UpdatePersonalityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := customValidator.Validate(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	personality, err := h.service.Update(uint(id), &req)
	if err != nil {
		logger.Errorf("Erro ao atualizar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Personalidade atualizada com sucesso", personality)
}

// Delete remove uma personalidade
func (h *PersonalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		logger.Errorf("Erro ao deletar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Personalidade deletada com sucesso", nil)
}
```

**Responsabilidades do Handler:**
- ✅ Receber requisições HTTP
- ✅ Validar entrada (JSON, params)
- ✅ Chamar service apropriado
- ✅ Retornar resposta HTTP

---

### **FASE 8: Middlewares**

#### Passo 8.1: Criar Middlewares (`internal/middleware/middleware.go`)

```bash
go get github.com/gorilla/handlers
```

```go
package middleware

import (
	"go-api-rest/pkg/logger"
	"net/http"
	"time"
)

// Logging registra informações sobre cada requisição
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Infof("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logger.Infof("Completado em %v", time.Since(start))
	})
}

// ContentTypeJSON define o Content-Type como JSON
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Recovery recupera de panics
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("Panic recuperado: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"success":false,"error":"Erro interno do servidor"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// CORS adiciona headers CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
```

**O que são Middlewares?**
- Funções que interceptam requisições antes de chegar no handler
- Úteis para: logs, autenticação, CORS, compressão, etc

---

### **FASE 9: Router (Configuração de Rotas)**

#### Passo 9.1: Criar Router (`internal/router/router.go`)

```go
package router

import (
	"go-api-rest/internal/handler"
	"go-api-rest/internal/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(personalityHandler *handler.PersonalityHandler) *mux.Router {
	r := mux.NewRouter()

	// Middlewares globais
	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)
	r.Use(middleware.ContentTypeJSON)

	// Rotas da API
	r.HandleFunc("/", personalityHandler.Home).Methods("GET")
	
	// Rotas de personalidades
	api := r.PathPrefix("/api/personalities").Subrouter()
	api.HandleFunc("", personalityHandler.Create).Methods("POST")
	api.HandleFunc("", personalityHandler.GetAll).Methods("GET")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.GetByID).Methods("GET")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.Update).Methods("PUT")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.Delete).Methods("DELETE")

	return r
}
```

---

### **FASE 10: Main (Ponto de Entrada)**

#### Passo 10.1: Criar Main (`cmd/api/main.go`)

```go
package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/internal/config"
	"go-api-rest/internal/handler"
	"go-api-rest/internal/repository"
	"go-api-rest/internal/router"
	"go-api-rest/internal/service"
	"go-api-rest/models"
	"go-api-rest/pkg/logger"
	"log"
	"net/http"
)

func main() {
	// 1. Carregar configurações
	cfg := config.Load()
	logger.Info("Configurações carregadas")

	// 2. Conectar ao banco de dados
	db, err := database.NewDatabase(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}
	logger.Info("Banco de dados conectado")

	// 3. Auto migrate (criar tabelas)
	if err := db.DB.AutoMigrate(&models.Personality{}); err != nil {
		log.Fatalf("Erro ao executar migrations: %v", err)
	}
	logger.Info("Migrations executadas com sucesso")

	// 4. Inicializar camadas da aplicação (Injeção de Dependência)
	personalityRepo := repository.NewPersonalityRepository(db.DB)
	personalityService := service.NewPersonalityService(personalityRepo)
	personalityHandler := handler.NewPersonalityHandler(personalityService)

	// 5. Configurar rotas
	r := router.SetupRoutes(personalityHandler)

	// 6. Iniciar servidor
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Infof("Servidor iniciado na porta http://localhost%s", addr)
	logger.Infof("Ambiente: %s", cfg.Server.Env)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
```

**Fluxo de Inicialização:**
1. Carregar configurações
2. Conectar ao banco
3. Executar migrations
4. Criar instâncias (dependency injection)
5. Configurar rotas
6. Iniciar servidor HTTP

---

### **FASE 11: Docker e Docker Compose (Opcional mas Recomendado)**

#### Passo 11.1: Criar `docker-compose.yml`

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: go-api-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./migration:/docker-entrypoint-initdb.d
    networks:
      - go-api-network

  app:
    build: .
    container_name: go-api-app
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: postgres
      DB_NAME: postgres
      DB_SSLMODE: disable
      SERVER_PORT: 8000
      ENV: production
    ports:
      - "8000:8000"
    depends_on:
      - postgres
    networks:
      - go-api-network

volumes:
  postgres_data:

networks:
  go-api-network:
    driver: bridge
```

#### Passo 11.2: Criar `Dockerfile`

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
```

#### Passo 11.3: Criar Script SQL Inicial (`migration/docker-database-initial.sql`)

```sql
-- Script executado automaticamente na criação do container

CREATE TABLE IF NOT EXISTS personalities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    history TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Dados de exemplo
INSERT INTO personalities (name, history) VALUES
('Albert Einstein', 'Físico teórico alemão, desenvolveu a teoria da relatividade.'),
('Marie Curie', 'Física e química polonesa, pioneira em radioatividade.')
ON CONFLICT DO NOTHING;
```

---

### **FASE 12: Makefile (Automação)**

#### Passo 12.1: Criar `Makefile`

```makefile
.PHONY: help run build test clean docker-up docker-down migrate

help: ## Mostra este help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Executa a aplicação
	go run cmd/api/main.go

build: ## Compila a aplicação
	go build -o bin/api cmd/api/main.go

test: ## Executa os testes
	go test -v ./...

clean: ## Remove arquivos compilados
	rm -rf bin/

docker-up: ## Inicia containers Docker
	docker-compose up -d

docker-down: ## Para containers Docker
	docker-compose down

docker-logs: ## Mostra logs dos containers
	docker-compose logs -f

install: ## Instala dependências
	go mod download
	go mod tidy
```

---

## 📂 Estrutura Final do Projeto

```
go-api-rest/
├── cmd/
│   └── api/
│       └── main.go                 # Ponto de entrada
├── internal/
│   ├── config/
│   │   └── config.go              # Configurações
│   ├── dto/
│   │   └── personality_dto.go     # DTOs
│   ├── handler/
│   │   └── personality_handler.go # Handlers HTTP
│   ├── middleware/
│   │   └── middleware.go          # Middlewares
│   ├── repository/
│   │   └── personality_repository.go # Acesso a dados
│   ├── router/
│   │   └── router.go              # Configuração de rotas
│   └── service/
│       └── personality_service.go # Lógica de negócio
├── models/
│   └── personality.go             # Modelos do banco
├── pkg/
│   ├── logger/
│   │   └── logger.go              # Sistema de logs
│   ├── response/
│   │   └── response.go            # Respostas padronizadas
│   └── validator/
│       └── validator.go           # Validações
├── database/
│   └── db.go                      # Conexão com banco
├── migration/
│   └── docker-database-initial.sql # Script SQL inicial
├── docker-compose.yml             # Orquestração Docker
├── Dockerfile                     # Build da imagem
├── Makefile                       # Comandos automatizados
├── go.mod                         # Dependências
└── README.md                      # Este arquivo
```

---

## 🚀 Como Usar

### Opção 1: Executar Localmente

```bash
# 1. Instalar dependências
go mod download

# 2. Iniciar PostgreSQL (se não estiver rodando)
# Certifique-se de ter um banco PostgreSQL rodando

# 3. Executar aplicação
go run cmd/api/main.go

# OU usar o Makefile
make run
```

### Opção 2: Executar com Docker

```bash
# 1. Iniciar todos os serviços
docker-compose up -d

# 2. Ver logs
docker-compose logs -f

# 3. Parar serviços
docker-compose down
```

---

## 📡 Testando a API

### 1. Verificar se está online

```bash
curl http://localhost:8000/
```

### 2. Criar uma personalidade

```bash
curl -X POST http://localhost:8000/api/personalities \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Isaac Newton",
    "history": "Físico e matemático inglês, autor dos Principia Mathematica"
  }'
```

### 3. Listar todas

```bash
curl http://localhost:8000/api/personalities
```

### 4. Buscar por ID

```bash
curl http://localhost:8000/api/personalities/1
```

### 5. Atualizar

```bash
curl -X PUT http://localhost:8000/api/personalities/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sir Isaac Newton"
  }'
```

### 6. Deletar

```bash
curl -X DELETE http://localhost:8000/api/personalities/1
```

---

## 🎯 Conceitos Avançados

### 1. **Dependency Injection (Injeção de Dependências)**

```go
// Em vez de criar dependências dentro das structs:
// ❌ Ruim
type Service struct {}
func NewService() *Service {
    repo := NewRepository() // acoplamento forte
    return &Service{repo: repo}
}

// ✅ Bom - recebe dependências
type Service struct {
    repo Repository
}
func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}
```

### 2. **Interface Segregation**

```go
// ✅ Interfaces específicas
type Reader interface {
    Read() error
}

type Writer interface {
    Write() error
}

// Use apenas o que precisa
func ProcessData(r Reader) {
    r.Read()
}
```

### 3. **Error Handling**

```go
// ✅ Sempre retorne e trate erros
func DoSomething() error {
    if err := step1(); err != nil {
        return fmt.Errorf("falha no step1: %w", err)
    }
    return nil
}
```

### 4. **Context para Timeouts**

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := db.WithContext(ctx).Find(&users).Error; err != nil {
    // Handle timeout
}
```

---

## 📚 Próximos Passos

Após dominar esta API básica, você pode adicionar:

1. **Autenticação JWT**
   - Login/Register
   - Middleware de autenticação
   - Refresh tokens

2. **Testes Automatizados**
   - Unit tests
   - Integration tests
   - Mocks

3. **Paginação**
   - Query params (page, limit)
   - Metadados de paginação

4. **Cache**
   - Redis
   - Cache in-memory

5. **Documentação**
   - Swagger/OpenAPI
   - Postman Collection

6. **CI/CD**
   - GitHub Actions
   - Deploy automático

7. **Observabilidade**
   - Prometheus/Grafana
   - Tracing distribuído
   - Health checks

---

## 🔗 Recursos Úteis

- [Go Official Documentation](https://go.dev/doc/)
- [GORM Documentation](https://gorm.io/docs/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Twelve-Factor App](https://12factor.net/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

---
