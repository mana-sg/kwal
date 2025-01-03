package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"

	t "github.com/mana-sg/kv-log-store/types"
)

func EncodeLog(log t.LogEntry) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(log); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DecodeLog(data []byte) (t.LogEntry, error) {
	var log t.LogEntry

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	err := decoder.Decode(&log)
	if err != nil {
		return t.LogEntry{}, fmt.Errorf("error decoding log: %v", err)
	}

	return log, nil
}
