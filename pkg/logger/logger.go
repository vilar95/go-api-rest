package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info registra uma mensagem informativa
func Info(message string) {
	InfoLogger.Println(message)
}

// Error registra uma mensagem de erro
func Error(message string) {
	ErrorLogger.Println(message)
}

// Infof registra uma mensagem informativa formatada
func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

// Errorf registra uma mensagem de erro formatada
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}
