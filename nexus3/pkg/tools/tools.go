package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func JsonMarshalInterfaceToIOReader(data any) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal data: %v", err)
	}

	return bytes.NewReader(b), nil
}

func GetEnv(key string, fallback any) any {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetIntPointer(number int) *int {
	return &number
}

func GetStringPointer(s string) *string {
	return &s
}

func GetBoolPointer(b bool) *bool {
	return &b
}

func GetSeededRandomInteger(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
