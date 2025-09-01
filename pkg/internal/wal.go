package internal

import (
	"bufio"
	"fmt"
	"os"

	t "github.com/mana-sg/kv-log-store/types"
	"github.com/mana-sg/kv-log-store/utils"
)

var LOGFILE string = "log.bin"

func WriteLog(op, key, value string) error {
	log := CreateLog(op, key, value)
	encodedLog, err := utils.EncodeLog(log)
	if err != nil {
		return fmt.Errorf("error encoding log in WriteLog: %v", err)
	}

	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error reading log file in WriteLog: %v", err)
	}
	defer file.Close()

	_, err = file.Write(encodedLog)
	if err != nil {
		return fmt.Errorf("error writing to log file in WriteLog: %v", err)
	}

	_, err = file.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("error writing newline to log file in WriteLog: %v", err)
	}

	return nil
}

func GetLogs() ([]t.LogEntry, error) {
	var logs []t.LogEntry

	if _, err := os.Stat(LOGFILE); os.IsNotExist(err) {
		f, err := os.Create(LOGFILE)
		defer f.Close()
		if err != nil {
			return nil, fmt.Errorf("error creating log file in GetLogs: %v", err)
		}
		return logs, nil
	}

	file, err := os.Open(LOGFILE)
	if err != nil {
		return nil, fmt.Errorf("error reading log file in GetLogs: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		log, err := utils.DecodeLog([]byte(line))
		if err != nil {
			return nil, fmt.Errorf("error decoding log in GetLogs: %v", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func CreateLog(op, key, value string) t.LogEntry {
	return t.LogEntry{
		Operation: op,
		Key:       key,
		Value:     value,
	}
}
