# Documentação da Arquitetura

## 📐 Estrutura do Projeto

A aplicação foi completamente feita seguindo os princípios de **Clean Architecture**, **SOLID** e **Clean Code**.

## 🏛️ Camadas da Arquitetura

### 1. Handler (Camada de Apresentação)
**Localização:** `internal/handler/`

**Responsabilidade:** Receber requisições HTTP, validar entrada e retornar respostas.

- Não contém lógica de negócio
- Converte dados HTTP para DTOs
- Chama a camada de serviço
- Retorna respostas padronizadas

```go
// Exemplo de handler
func (h *PersonalityHandler) Create(w http.ResponseWriter, r *http.Request) {
    // 1. Parse do request
    // 2. Validação
    // 3. Chama o service
    // 4. Retorna resposta
}
```

### 2. Service (Camada de Negócio)
**Localização:** `internal/service/`

**Responsabilidade:** Implementar regras de negócio e orquestrar operações.

- Validações de negócio (ex: nome único)
- Orquestração de múltiplas operações
- Transformação de dados entre DTOs e Models
- Tratamento de erros específicos

```go
// Exemplo de service
func (s *personalityService) Create(req *dto.CreatePersonalityRequest) (*dto.PersonalityResponse, error) {
    // 1. Validações de negócio
    // 2. Conversão para modelo
    // 3. Chama repository
    // 4. Retorna DTO
}
```

### 3. Repository (Camada de Dados)
**Localização:** `internal/repository/`

**Responsabilidade:** Acesso ao banco de dados e persistência.

- Operações CRUD
- Queries específicas
- Abstração do banco de dados (fácil trocar de DB)
- Trabalha com models do domínio

```go
// Interface do repository
type PersonalityRepository interface {
    Create(personality *models.Personality) error
    FindAll() ([]models.Personality, error)
    FindByID(id uint) (*models.Personality, error)
    Update(personality *models.Personality) error
    Delete(id uint) error
}
```

## 📊 Fluxo de Dados

```
Request → Middleware → Handler → Service → Repository → Database
                          ↓         ↓          ↓
                        DTO     Business    Model
                                 Logic
```

## 🔧 Componentes Principais

### DTOs (Data Transfer Objects)
**Localização:** `internal/dto/`

Separa a representação externa (API) da interna (domínio).

**Benefícios:**
- Controla o que é exposto na API
- Permite diferentes formatos de request/response
- Facilita versionamento da API
- Validações específicas por operação

### Models
**Localização:** `models/`

Representa o domínio da aplicação.

```go
type Personality struct {
    ID        uint
    Name      string
    History   string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Middleware
**Localização:** `internal/middleware/`

Funções que interceptam requisições antes dos handlers.

**Implementados:**
- `Recovery`: Captura panics e retorna erro 500
- `Logging`: Log de todas as requisições
- `CORS`: Configuração de CORS
- `ContentTypeJSON`: Define Content-Type JSON

### Configuration
**Localização:** `internal/config/`

Gerenciamento centralizado de configurações via variáveis de ambiente.

**Benefícios:**
- Configuração por ambiente (dev, staging, prod)
- Valores padrão seguros
- Fácil deploy em diferentes ambientes

## 🎯 Princípios Aplicados

### 1. Single Responsibility Principle (SRP)
Cada camada/struct tem uma única responsabilidade:
- Handler: HTTP
- Service: Lógica de negócio
- Repository: Persistência

### 2. Dependency Inversion Principle (DIP)
Dependência de abstrações (interfaces), não de implementações:

```go
type PersonalityService interface {
    Create(req *dto.CreatePersonalityRequest) (*dto.PersonalityResponse, error)
    // ...
}
```

### 3. Interface Segregation Principle (ISP)
Interfaces pequenas e específicas.

### 4. Open/Closed Principle (OCP)
Código aberto para extensão, fechado para modificação.

### 5. Dependency Injection
Todas as dependências são injetadas no `main.go`:

```go
personalityRepo := repository.NewPersonalityRepository(db.DB)
personalityService := service.NewPersonalityService(personalityRepo)
personalityHandler := handler.NewPersonalityHandler(personalityService)
```

## 🧪 Testabilidade

A arquitetura facilita testes em múltiplos níveis:

### Unit Tests
- Testar services com repository mockado
- Testar handlers com service mockado

### Integration Tests
- Testar com banco de dados real ou em memória

### Exemplo de Mock
```go
type mockRepository struct{}

func (m *mockRepository) Create(p *models.Personality) error {
    return nil
}
```

## 📦 Organização de Pacotes

### `internal/`
Código privado da aplicação (não pode ser importado por outros projetos).

### `pkg/`
Código reutilizável que pode ser importado.

### `models/`
Modelos de domínio (entities).

## 🔄 Vantagens da Arquitetura

1. **Manutenibilidade**: Fácil localizar e modificar código
2. **Testabilidade**: Cada camada pode ser testada independentemente
3. **Escalabilidade**: Fácil adicionar novas features
4. **Flexibilidade**: Trocar implementações (ex: mudar de banco)
5. **Reutilização**: Código bem organizado e modular
6. **Documentação**: Estrutura autoexplicativa

## 🚀 Como Adicionar Novas Features

### 1. Adicionar nova entidade
```
1. Criar model em models/
2. Criar DTOs em internal/dto/
3. Criar repository interface e implementação
4. Criar service interface e implementação
5. Criar handler
6. Registrar rotas no router
```

### 2. Adicionar novo endpoint
```
1. Criar método no handler
2. Adicionar rota no router
3. Implementar método no service (se necessário)
4. Implementar query no repository (se necessário)
```

## 📚 Referências

- [Clean Architecture - Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go)
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID)
