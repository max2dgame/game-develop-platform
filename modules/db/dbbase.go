package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/max2dgame/game-develop-platform/modules/config"
	"sync"
	"xorm.io/xorm"
)

type DBBase struct {
	DBList  map[string]*sql.DB
	Once    sync.Once
	Configs config.DatabaseList
}

func (db *DBBase) GetConfig(name string) config.DatabaseConfig {
	v, has := db.Configs[name]
	if !has {
		panic("Database config not found with key: " + name)
	}
	return v
}

func (db *DBBase) Create(name string, params ...interface{}) error {
	cfg := db.GetConfig(name)
	engine, err := xorm.NewEngine("mysql", cfg.GetDSN())
	if err != nil {
		return err
	}
	defer func() {
		_ = engine.Close()
	}()
	err = engine.Sync(params...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBBase) Close() []error {
	errs := make([]error, 0)
	for _, d := range db.DBList {
		errs = append(errs, d.Close())
	}
	return errs
}
