package storage

import (
	"fmt"
	"strings"

	"github.com/mana-sg/kv-log-store/pkg/wal"
)

type KVStore struct {
	Store map[string]string
}

var Kv_store = KVStore{
	Store: make(map[string]string),
}

func (kv *KVStore) Set(key, value string) error {
	if strings.Compare(key, "") == 0 {
		return fmt.Errorf("cannot have empty key")
	}

	if strings.Compare(value, "") == 0 {
		return fmt.Errorf("cannot have empty value")
	}
	err := wal.WriteLog("SET", key, value)
	if err != nil {
		return fmt.Errorf("error writing to log file: %v", err)
	}

	kv.Store[key] = value
	return nil
}

func (kv *KVStore) Get(key string) (string, error) {
	if strings.Compare(key, "") == 0 {
		return "", fmt.Errorf("cannot have empty key")
	}
	value, ok := kv.Store[key]

	if !ok {
		return "", fmt.Errorf("key does not exist")
	}

	return value, nil
}

func (kv *KVStore) Remove(key string) error {
	_, ok := kv.Store[key]
	if !ok {
		return fmt.Errorf("key does not exist")
	}

	err := wal.WriteLog("DELETE", key, "")
	if err != nil {
		return fmt.Errorf("error writing to log file: %v", err)
	}

	delete(kv.Store, key)
	return nil
}

func (kv *KVStore) BuildStore() error {
	logs, err := wal.GetLogs()
	if err != nil {
		return fmt.Errorf("error getting logs: %v", err)
	}

	for _, log := range logs {
		if log.Operation == "SET" {
			err := kv.Set(log.Key, log.Value)
			if err != nil {
				return fmt.Errorf("error setting value in BuildStore: %v", err)
			}
		} else if log.Operation == "DELETE" {
			err := kv.Remove(log.Key)
			if err != nil {
				return fmt.Errorf("del log created for key does not exist")
			}
		}
	}

	return nil
}
