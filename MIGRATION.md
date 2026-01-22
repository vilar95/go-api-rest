# Guia de Migração da Aplicação

## 📝 Resumo das Mudanças

A aplicação foi completamente refatorada de uma estrutura simples para uma arquitetura em camadas seguindo Clean Architecture.

## 🔄 O Que Mudou

### Estrutura Antiga
```
controllers/controllers.go  → Tudo misturado
database/db.go              → Conexão global
models/personality.go       → Model simples
routes/routes.go            → Rotas simples
```

### Estrutura Nova
```
internal/
  ├── config/          → Configurações centralizadas
  ├── dto/             → Request/Response objects
  ├── handler/         → HTTP handlers (ex-controllers)
  ├── middleware/      → Middlewares melhorados
  ├── repository/      → Acesso a dados
  ├── router/          → Configuração de rotas
  └── service/         → Lógica de negócio
pkg/
  ├── logger/          → Logger customizado
  ├── response/        → Helpers de resposta
  └── validator/       → Validação de dados
models/                → Models com timestamps
database/              → Database refatorado
```

## ⚠️ Breaking Changes

### 1. Models
**Antes:**
```go
type Personality struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    History string `json:"history"`
}
```

**Depois:**
```go
type Personality struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    History   string    `json:"history"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**Impacto:** Respostas da API agora incluem `created_at` e `updated_at`.

### 2. Respostas de Erro
**Antes:**
```
Personalidade não encontrada
```

**Depois:**
```json
{
  "error": "Not Found",
  "message": "personalidade não encontrada"
}
```

**Impacto:** Clientes devem lidar com o novo formato JSON de erro.

### 3. Validações
Agora há validações mais rigorosas:
- `name`: 3-100 caracteres
- `history`: 10-5000 caracteres

**Resposta de validação:**
```json
{
  "error": "Erro de validação",
  "message": "Os dados fornecidos são inválidos",
  "details": {
    "name": "O campo name deve ter no mínimo 3 caracteres"
  }
}
```

## 🚀 Como Migrar

### Passo 1: Backup
```bash
# Fazer backup do banco de dados
docker exec -t postgres pg_dump -U vilar postgres > backup.sql
```

### Passo 2: Parar aplicação antiga
Pare qualquer instância em execução na porta 8000.

### Passo 3: Instalar dependências
```bash
go mod download
go mod tidy
```

### Passo 4: Configurar variáveis de ambiente (opcional)
```bash
cp .env.example .env
# Editar .env conforme necessário
```

### Passo 5: Executar migrations
A aplicação executa auto-migration ao iniciar, mas você pode verificar:
```bash
# As migrations são automáticas
go run cmd/api/main.go
```

### Passo 6: Testar
```bash
# Executar testes
go test ./...

# Ou usar o Makefile
make test
```

## 🧹 Arquivos que Podem Ser Removidos

Após confirmar que tudo está funcionando:

```bash
# Arquivos antigos (backup antes de remover!)
controllers/controllers.go  # Substituído por internal/handler/
routes/routes.go           # Substituído por internal/router/
middleware/middleware.go   # Substituído por internal/middleware/
```

**⚠️ NÃO REMOVA até ter certeza que está tudo funcionando!**

## 📊 Comparação de Performance

A nova estrutura não deve ter impacto significativo na performance, mas oferece:

✅ Melhor organização do código
✅ Mais fácil de testar
✅ Mais fácil de manter
✅ Validações mais robustas
✅ Melhor tratamento de erros
✅ Logs estruturados

## 🐛 Solução de Problemas

### Erro: "Personality not found"
**Causa:** IDs agora são `uint` em vez de `int`
**Solução:** Verificar se está enviando IDs válidos (> 0)

### Erro: "Erro de validação"
**Causa:** Dados não atendem aos critérios mínimos
**Solução:** Verificar os detalhes do erro no campo `details`

### Erro: "Já existe uma personalidade com esse nome"
**Causa:** Validação de unicidade foi implementada
**Solução:** Usar um nome diferente ou atualizar a existente

### Erro de Compilação
**Causa:** Falta de dependências
**Solução:** 
```bash
go mod download
go mod tidy
```

## 🔧 Configuração

### Variáveis de Ambiente Disponíveis
```env
SERVER_PORT=8000           # Porta do servidor
ENV=development            # Ambiente (development, production)
DB_HOST=localhost          # Host do banco
DB_PORT=5432              # Porta do banco
DB_USER=vilar             # Usuário do banco
DB_PASSWORD=vilar123      # Senha do banco
DB_NAME=postgres          # Nome do banco
DB_SSLMODE=disable        # SSL mode
```

## 📈 Próximos Passos

### Melhorias Futuras Sugeridas
1. Adicionar autenticação JWT
2. Implementar paginação
3. Adicionar cache (Redis)
4. Implementar rate limiting
5. Adicionar métricas (Prometheus)
6. Adicionar documentação OpenAPI/Swagger
7. Implementar CI/CD
8. Adicionar testes de integração
9. Adicionar health check endpoint
10. Implementar graceful shutdown

## 📚 Documentação Adicional

- [README.md](README.md) - Documentação geral
- [ARCHITECTURE.md](ARCHITECTURE.md) - Detalhes da arquitetura
- [Makefile](Makefile) - Comandos úteis

## 💬 Suporte

Se encontrar problemas:
1. Verificar os logs da aplicação
2. Verificar se o banco de dados está acessível
3. Verificar se todas as dependências foram instaladas
4. Consultar a documentação da arquitetura
