package handlerutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func JsonRequestBody(r *http.Request, dest any) error {
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return errors.New("content type not json")
	}

	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json body: %w", err)
	}

	return nil
}
