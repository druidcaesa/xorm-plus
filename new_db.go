package xormplus

import "github.com/go-xorm/xorm"

var Engine *xorm.Engine

func NewEngine(driverName string, dataSourceName string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	Engine = engine
	return Engine, nil
}
