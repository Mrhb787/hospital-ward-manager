package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

type healthRequest struct {
}

type healthResponse struct {
	Status string `json:"status"`
}

type httpResponse interface {
	error() error
}

type data interface {
	data() interface{}
}

func (r healthResponse) error() error {
	return nil
}

func (r healthResponse) data() interface{} {
	return r.Status
}

func GenericSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GenericErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func DecodeHealthRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return healthRequest{}, nil
}

func EncodeHealthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if resp, ok := response.(httpResponse); ok && resp.error() != nil {
		resp.error()
	}

	w.Header().Set("Content-Type", "application/health+json; charset=utf-8")
	if response.(data).data() != "pass" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"status": response.(data).data(),
	})
}
