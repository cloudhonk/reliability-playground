package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Fast(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("fast response\n"))
}

func Slow(w http.ResponseWriter, r *http.Request) {
	msStr := r.URL.Query().Get("ms")
	ms, _ := strconv.Atoi(msStr)
	if ms <= 0 {
		ms = 500
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("slow response\n"))
}

func Error(w http.ResponseWriter, r *http.Request) {
	rateStr := r.URL.Query().Get("rate")
	rate, _ := strconv.ParseFloat(rateStr, 64)
	if rate <= 0 {
		rate = 0.3
	}

	if rand.Float64() < rate {
		http.Error(w, "simulated error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success\n"))
}

func Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok\n"))
}
