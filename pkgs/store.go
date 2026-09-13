package pkgs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var AppStateFile string

func init() {
	var err error
	AppStateFile, err = GetAppStateFile()

	if err != nil {
		log.Panic(err)
	}
}

func WriteStateFile(todos any) error {
	file, err := os.Create(AppStateFile)
	if err != nil {
		return fmt.Errorf("Error opening state file %w\n", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(todos); err != nil {
		return fmt.Errorf("Error writing state file %w\n", err)
	}

	return nil
}
