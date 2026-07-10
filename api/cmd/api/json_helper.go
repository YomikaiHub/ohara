package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error APIError `json:"error"`
}

const maxBodyBytes = 1024 * 1024

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.Is(err, io.EOF):
			return errors.New("request body must not be empty")

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("request body contains malformed JSON")

		case errors.As(err, &syntaxError):
			return fmt.Errorf(
				"request body contains badly-formed JSON (at position %d)",
				syntaxError.Offset,
			)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf(
					"request body contains incorrect JSON type for field %q",
					unmarshalTypeError.Field,
				)
			}

			return fmt.Errorf(
				"request body contains incorrect JSON type (at position %d)",
				unmarshalTypeError.Offset,
			)

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("request body contains unknown field %s", field)

		case strings.HasPrefix(err.Error(), "http: request body too large"):
			return fmt.Errorf(
				"request body must not be larger than %d bytes",
				maxBodyBytes,
			)

		case errors.As(err, &invalidUnmarshalError):
			// programmer error, not a client error.
			panic(err)

		default:
			return err
		}
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("request body must only contain a single JSON object")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	if len(headers) > 0 {
		for key, values := range headers[0] {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func writeErrorJSON(
	w http.ResponseWriter,
	status int,
	code string,
	message string,
) error {

	response := ErrorResponse{
		Error: APIError{
			Code:    code,
			Message: message,
		},
	}

	return writeJSON(
		w,
		status,
		response,
		nil,
	)
}
