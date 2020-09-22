package confgorm

import (
	"path/filepath"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	DbFile string
	*gorm.DB
}

var sqliteLock = sync.Mutex{}

func (s *Sqlite) Init() {
	sqliteLock.Lock()
	defer sqliteLock.Unlock()

	if s.DB == nil {
		s.initial()
	}
}

func (s *Sqlite) initial() {
	s.SetDefaults()

	dir := filepath.Dir(s.DbFile)
	if !DirExists(dir) {
		MustMkdir(dir)
	}

	db, err := gorm.Open(sqlite.Open(s.DbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	s.DB = db
}

func (s *Sqlite) SetDefaults() {
	if s.DbFile == "" {
		s.DbFile = "sqlite.db"
	}
}
