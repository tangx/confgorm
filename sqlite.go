package confgorm

import (
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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
