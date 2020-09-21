package confmysql

import "testing"

var (
	m = Mysql{
		Username: "root",
		Password: "Mysql12345",
		Host:     "127.0.0.1",
	}
)

func TestPing(t *testing.T) {

	m.Init()
	m.Ping()

}
