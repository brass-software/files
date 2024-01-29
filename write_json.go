package files

import (
	"bytes"
	"encoding/json"

	"github.com/brass-software/api"
)

func WriteJSON(path string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return api.PUT(path, bytes.NewReader(b))
}
