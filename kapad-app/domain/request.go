package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

func DecodeBody[T any](body io.Reader) (*T, error) {
	var data T
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields() // optional: catch unknown fields
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode body: %w", err)
	}
	return &data, nil
}

func ToJSON[T any](data T) ([]byte, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	return jsonBytes, nil
}

func ToJSONString[T any](data T) (string, error) {
	jsonBytes, err := ToJSON(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
