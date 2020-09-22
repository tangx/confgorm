package confgorm

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host     string `env:"host,omitempty"`
	Port     int    `env:"port,omitempty"`
	Username string `env:"username,omitempty"`
	Password string `env:"password,omitempty"`
	Dbname   string `env:"dbname,omitempty"`
	Charset  string `env:"charset,omitempty"`
	*gorm.DB
}

var mysqlLock = sync.Mutex{}

func (m *Mysql) Init() {
	mysqlLock.Lock()
	defer mysqlLock.Unlock()

	if m.DB == nil {
		m.initial()
	}

}

func (m *Mysql) initial() {
	m.SetDefaults()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal(err)
		// return
		panic(err)
	}
	m.DB = db
}

func (m *Mysql) SetDefaults() {
	if m.Port == 0 {
		m.Port = 3306
	}
	if m.Charset == "" {
		m.Charset = "utf8mb4"
	}
}

func (m *Mysql) Ping() {
	var result interface{}
	m.DB.Raw("select 1;").Scan(&result)
}
