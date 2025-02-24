package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of http requests.",
	},
)

func main() {

	prometheus.MustRegister(httpRequestsTotal)

	httpMux := http.NewServeMux()

	httpMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	httpMux.HandleFunc("/", handleHomePage)
	httpMux.HandleFunc("/api/pay", handlePayment)
	httpMux.Handle("/metrics", promhttp.Handler())

	log.Printf("Payment-microservice listening on 8080")
	http.ListenAndServe(":8080", httpMux) // Run the service on port 8080

}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "templates/index.html")
}

func handlePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
		return
	}

	httpRequestsTotal.Inc()

	type PaymentRequest struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	var req PaymentRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input"})
		return
	}

	// Simulate payment processing logic
	response := map[string]interface{}{
		"status":   "success",
		"message":  "Payment processed successfully!",
		"amount":   req.Amount,
		"currency": req.Currency,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
