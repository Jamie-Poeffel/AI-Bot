package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	"net/http"
)

type prefixWriter struct {
	logger *log.Logger
	prefix string
}

func (p *prefixWriter) Write(b []byte) (int, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s] %s %s", now, p.prefix, b)
	return p.logger.Writer().Write([]byte(msg))
}

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {

	date := time.Now().Format("20060102")
	logFilePath := fmt.Sprintf("logs/%s.log", date)
	os.MkdirAll("logs", os.ModePerm)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Fehler beim Öffnen der Log-Datei: %v", err)
	}

	baseLogger := log.New(file, "", 0) 

	InfoLog = log.New(&prefixWriter{logger: baseLogger, prefix: "[REQUEST]"}, "", 0)
	ErrorLog = log.New(&prefixWriter{logger: baseLogger, prefix: "[ERROR]"}, "", 0)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	InfoLog.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	fmt.Fprintln(w, "Hallo Welt!")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	ErrorLog.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	http.Error(w, "Fehler aufgetreten!", http.StatusInternalServerError)
}
