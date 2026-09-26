package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func parseParams(r *http.Request) (int, int, error) {
	a, err1 := strconv.Atoi(r.URL.Query().Get("term_one"))
	b, err2 := strconv.Atoi(r.URL.Query().Get("term_two"))
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("parâmetros inválidos ou ausentes")
	}
	return a, b, nil
}

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"UP"}`))
	})

	mux.HandleFunc("/api/sum", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := parseParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": a + b})
	})

	mux.HandleFunc("/api/sub", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := parseParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": a - b})
	})

	mux.HandleFunc("/api/mul", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := parseParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": a * b})
	})

	mux.HandleFunc("/api/div", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := parseParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if b == 0 {
			http.Error(w, "divisao por zero nao e permitida", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": a / b})
	})

	return mux
}

func main() {
	mux := setupRouter()
	log.Println("Servidor iniciado na porta :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}