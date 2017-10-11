package sql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DefaultDatabaseCfg DatabaseConfig = DatabaseConfig{
	Port:     3306,
	DB:       "probe",
	User:     "root",
	Password: "123456",
	ConnIdle: 20,
	ConnMax:  200,
	ShowSQL:  true,
}

type DatabaseConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	ConnIdle int
	ConnMax  int
	ShowSQL  bool
}

func (c DatabaseConfig) Addr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai",
		c.User, c.Password, c.Host, c.Port, c.DB)
}

func InitMySQL(cfg *DatabaseConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", cfg.Addr())
	if err != nil {
		return nil, err
	}
	if err = engine.Ping(); err != nil {
		return nil, err
	}
	engine.SetMaxIdleConns(cfg.ConnIdle)
	engine.SetMaxOpenConns(cfg.ConnMax)
	engine.ShowSQL(cfg.ShowSQL)

	return engine, nil
}
