package xormplus

import (
	"github.com/go-xorm/xorm"
)

type InsertWrapper[T any] struct {
	DB   *xorm.Engine
	isTx bool
}

// NewInsertWrapper 创建构造器
func NewInsertWrapper[T any]() *InsertWrapper[T] {
	w := &InsertWrapper[T]{
		DB: Engine,
	}
	return w
}

// OpenTX 开启事务
func (u *InsertWrapper[T]) OpenTX() *InsertWrapper[T] {
	u.isTx = true
	return u
}

// Insert 添加数据
func (u *InsertWrapper[T]) Insert(t *T) (int64, error) {
	if u.isTx {
		session := u.DB.NewSession()
		session.Begin()
		insert, err := session.Insert(t)
		if err != nil {
			session.Rollback()
			return 0, err
		}
		session.Commit()
		return insert, nil
	}
	insert, err := u.DB.Insert(t)
	if err != nil {
		return 0, err
	}
	return insert, nil
}

// CreateInBatches 批量添加
func (u *InsertWrapper[T]) CreateInBatches(list []*T) (int64, error) {
	if u.isTx {
		session := u.DB.NewSession()
		session.Begin()
		insert, err := session.Insert(list)
		if err != nil {
			session.Rollback()
			return 0, err
		}
		session.Commit()
		return insert, nil
	}
	insert, err := u.DB.Insert(list)
	if err != nil {
		return 0, nil
	}
	return insert, err
}
