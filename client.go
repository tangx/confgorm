package confmysql

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	Charset  string
	db       *gorm.DB
}

var lock = sync.Mutex{}

func (m *Mysql) Init() {
	lock.Lock()
	defer lock.Unlock()

	if m.db != nil {
		m.initial()
	}

}
func (m *Mysql) initial() {
	m.SetDefaults()

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	m.db = db
}

func (m *Mysql) SetDefaults() {
	if m.Port == 0 {
		m.Port = 3306
	}
	if m.Charset == "" {
		m.Charset = "utf8mb4"
	}
}
