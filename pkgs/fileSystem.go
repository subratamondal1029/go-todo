package pkgs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const appDir string = "todo"

func GetAppStateDir() (string, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "", fmt.Errorf("User config directory not found")
	}

	joinedPath := filepath.Join(configDir, appDir)

	pathStat, err := os.Stat(joinedPath)

	if err == nil {
		if pathStat.IsDir() {
			return joinedPath, nil
		} else {
			return "", fmt.Errorf("App state directory is not a directory")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(joinedPath, 0755)
		if err != nil {
			return "", fmt.Errorf("Error occurred while creating app state directory")
		}
		return joinedPath, nil
	} else {
		return "", fmt.Errorf("Error occurred while getting app state directory")
	}
}

func GetAppStateFile() (string, error) {
	appDir, err := GetAppStateDir()
	if err != nil {
		return "", err
	}

	stateFile := filepath.Join(appDir, "todos.json")

	fileStat, err := os.Stat(stateFile)

	if err == nil {
		if fileStat.IsDir() {
			return "", fmt.Errorf("App state file is not a file")
		}

		return stateFile, nil
	} else if errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(stateFile, []byte("[]"), 0644)

		if err != nil {
			return "", fmt.Errorf("Error occurred while creating app state file")
		}

		return stateFile, nil
	} else {
		return "", fmt.Errorf("Error occurred while getting app state file")
	}
}
