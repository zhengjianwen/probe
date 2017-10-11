package sql

import "testing"

func TestInitMySQL(t *testing.T) {
	DefaultDatabaseCfg.Host = "192.168.99.100"
	if _, err := InitMySQL(&DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}
}
