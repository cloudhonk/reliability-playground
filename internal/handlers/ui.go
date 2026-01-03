package handlers

import (
	"html/template"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/cloudhonk/reliability-playground/internal/system"
)

var tpl = template.Must(template.ParseFiles("web/templates/index.html"))

func UI(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"Uptime":     time.Since(system.StartTime).Round(time.Second),
		"Goroutines": runtime.NumGoroutine(),
		"MemoryMB":   system.MemoryMB(),
		"Pod":        os.Getenv("HOSTNAME"),
	}

	tpl.Execute(w, data)
}
