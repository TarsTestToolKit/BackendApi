package mysql

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"

	"github.com/TarsTestToolKit/BackendApi/config"
)

type dbManager struct {
	dbs map[string]*xorm.Engine
	mux *sync.RWMutex
}

var dbm *dbManager

func init() {
	if dbm == nil {
		dbm = new(dbManager)
		dbm.dbs = make(map[string]*xorm.Engine)
		dbm.mux = new(sync.RWMutex)
	}
}

func newSession(cfgName string) (*xorm.Session, error) {
	dbm.mux.RLock()
	if engine, ok := dbm.dbs[cfgName]; ok {
		dbm.mux.RUnlock()
		return engine.NewSession(), nil
	}
	dbm.mux.RUnlock()
	connStr := config.GetDBCfg(cfgName)
	engine, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		return nil, err
	}
	dbm.mux.Lock()
	dbm.dbs[cfgName] = engine
	dbm.mux.Unlock()

	return engine.NewSession(), nil
}

func newTestDBSess() (*xorm.Session, error) {
	return newSession(config.MySQLTestDB)
}

func newTarsDBSess() (*xorm.Session, error) {
	return newSession(config.MySQLTarsDB)
}
