package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func jsonMarshalInterfaceToIOReader(data interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal data: %v", err)
	}

	return bytes.NewReader(b), nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
