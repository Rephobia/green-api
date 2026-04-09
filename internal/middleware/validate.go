package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rephobia/green-api-test-task/internal/response"
	"github.com/Rephobia/green-api-test-task/internal/validatorx"
	"github.com/gorilla/schema"
)

func Validate[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req T

		if err := fillRequest(r, &req); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())

			return
		}

		if errs := validatorx.Validate(req); len(errs) > 0 {
			response.Error(w, http.StatusBadRequest, errs)

			return
		}

		next(w, r, req)
	}
}

func fillRequest[T any](r *http.Request, req *T) error {
	if err := schema.NewDecoder().Decode(req, r.URL.Query()); err != nil {
		return fmt.Errorf("invalid query parameters: %w", err)
	}

	if r.Body != nil && r.ContentLength != 0 && r.Method != http.MethodGet {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return fmt.Errorf("invalid json body: %w", err)
		}
	}

	return nil
}
