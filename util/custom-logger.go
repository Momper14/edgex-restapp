package util

// This is a copy of negroni's logger which filters healthcheck from logging

import (
	"bytes"
	"strings"

	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/urfave/negroni"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// ALogger implements just enough log.Logger interface to be compatible with other implementations
	negroni.ALogger
	dateFormat string
	template   *template.Template
}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	logger := &Logger{ALogger: log.New(os.Stdout, "[negroni] ", 0), dateFormat: negroni.LoggerDefaultDateFormat}
	logger.SetFormat(negroni.LoggerDefaultFormat)
	return logger
}

func (l *Logger) SetFormat(format string) {
	l.template = template.Must(template.New("negroni_parser").Parse(format))
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	if strings.Split(r.RemoteAddr, ":")[0] == "127.0.0.1" && r.URL.Path == "/api/v1/ping" {
		return
	}

	res := rw.(negroni.ResponseWriter)
	log := negroni.LoggerEntry{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  time.Since(start),
		Hostname:  r.Host,
		Method:    r.Method,
		Path:      r.URL.Path,
		Request:   r,
	}

	buff := &bytes.Buffer{}
	l.template.Execute(buff, log)
	l.Println(buff.String())
}
