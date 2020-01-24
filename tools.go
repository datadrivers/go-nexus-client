package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func jsonMarshalInterfaceToIOReader(data interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal data: %v", err)
	}

	return bytes.NewReader(b), nil
}
