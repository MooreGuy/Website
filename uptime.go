package main

import (
	"fmt"
	"net/http"
	"time"
)

type UptimeHandler struct {
	Started time.Time
	format  string
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uptime := fmt.Sprintf("Uptime: %s", time.Since(h.Started))
	fmt.Fprintf(w, uptime)
}
