package middleware

import (
	"go-api-rest/pkg/logger"
	"net/http"
	"time"
)

// ContentTypeJSON define o Content-Type como application/json
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// Logging registra informações sobre cada requisição
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log da requisição
		logger.Infof("Iniciando %s %s", r.Method, r.URL.Path)

		// Processar requisição
		next.ServeHTTP(w, r)

		// Log da duração
		duration := time.Since(start)
		logger.Infof("Completado %s %s em %v", r.Method, r.URL.Path, duration)
	})
}

// Recovery recupera de panics e retorna um erro 500
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("Panic recuperado: %v", err)
				http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// CORS configura os headers CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Responder a preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
