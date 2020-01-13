package initialiser

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// for current configuration
const cdmgrFolder = ".cdmgr"
const db = "cdmgr.db"

// To set default bucket in main.go
const TaskBucket = "cdmgr"
const TagMasterData = "tagMasterData"

func ReturnDbPath() string {
	home, _ := homedir.Dir()
	home = filepath.Join(home, cdmgrFolder)
	dbPath := filepath.Join(home, db)
	return dbPath
}

func InitiiseRootFolder() {
	home, _ := homedir.Dir()
	home = filepath.Join(home, cdmgrFolder)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		os.Mkdir(home, os.ModeDir)
	}
}
