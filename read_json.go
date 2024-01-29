package files

import (
	"bytes"
	"encoding/json"

	"github.com/brass-software/api"
)

func ReadJSON[T any](path string) (*T, error) {
	buf := bytes.NewBuffer(nil)
	err := api.GET(buf, path)
	if err != nil {
		return nil, err
	}
	var v T
	err = json.NewDecoder(buf).Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
