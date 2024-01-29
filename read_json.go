package files

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/brass-software/api"
)

func ReadJSON[T any](path string) (*T, error) {
	var v T
	buf := bytes.NewBuffer(nil)
	err := api.GET(buf, path)
	if err != nil {
		if strings.HasPrefix(err.Error(), "404") {
			return nil, nil
		}
		return nil, err
	}
	err = json.NewDecoder(buf).Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
