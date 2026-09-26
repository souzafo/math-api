package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {
	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedResult int
	}{
		{"Adicao", "/api/sum?term_one=4&term_two=2", http.StatusOK, 6},
		{"Subtracao", "/api/sub?term_one=4&term_two=1", http.StatusOK, 3},
		{"Multiplicacao", "/api/mul?term_one=3&term_two=5", http.StatusOK, 15},
		{"Divisao", "/api/div?term_one=10&term_two=2", http.StatusOK, 5},
		{"Healthcheck", "/healthcheck", http.StatusOK, 0},
		{"Divisao por zero", "/api/div?term_one=10&term_two=0", http.StatusBadRequest, 0},
		{"Parametro ausente", "/api/sum?term_one=10", http.StatusBadRequest, 0},
		{"Parametro invalido", "/api/sum?term_one=abc&term_two=5", http.StatusBadRequest, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tt.path)
			if err != nil {
				t.Fatalf("Erro na requisicao: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Esperava status %d, obteve %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK && tt.path != "/healthcheck" {
				var res map[string]int
				if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
					t.Fatalf("Erro ao decodificar JSON: %v", err)
				}
				if res["result"] != tt.expectedResult {
					t.Errorf("Esperava resultado %d, obteve %d", tt.expectedResult, res["result"])
				}
			}
		})
	}
}