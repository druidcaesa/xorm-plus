package xormplus

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"strings"
)

type UpdateWrapper[T any] struct {
	DB   *xorm.Session
	isTx bool
}

func NewUpdateWrapper[T any]() *UpdateWrapper[T] {
	w := &UpdateWrapper[T]{
		DB: Engine.NewSession(),
	}
	return w
}

// Set 选择更新字段
func (w *UpdateWrapper[T]) Set(columns ...string) *UpdateWrapper[T] {
	w.DB = w.DB.Cols(strings.Join(columns, ","))
	return w
}

// NotSet 不添加的字段
func (w *UpdateWrapper[T]) NotSet(columns ...string) *UpdateWrapper[T] {
	for _, c := range columns {
		w.DB = w.DB.Omit(c)
	}
	return w
}

// Update 更新数据
func (w *UpdateWrapper[T]) Update(t *T) (int64, error) {
	if w.isTx {
		w.DB.Begin()
		update, err := w.Update(t)
		if err != nil {
			log.Fatalf("更新数据发生异常:%v\n", err)
			w.DB.Rollback()
			return 0, err
		}
		w.DB.Commit()
		return update, nil
	}
	return w.DB.Update(t)
}

// Eq 等于查询
func (w *UpdateWrapper[T]) Eq(column string, attribute interface{}) *UpdateWrapper[T] {
	w.DB = w.DB.Where(fmt.Sprintf("%s = ?", column), attribute)
	return w
}

// Ne 不等于查询
func (w *UpdateWrapper[T]) Ne(column string, attribute interface{}) *UpdateWrapper[T] {
	w.DB = w.DB.Where(fmt.Sprintf("%s <> ?", column), attribute)
	return w
}

// OrEq 或查询
func (w *UpdateWrapper[T]) OrEq(column string, attribute interface{}) *UpdateWrapper[T] {
	w.DB = w.DB.Or(fmt.Sprintf("%s = ?", column), attribute)
	return w
}

// NotIn 条件查询
func (w *UpdateWrapper[T]) NotIn(args ...string) *UpdateWrapper[T] {
	w.DB = w.DB.NotIn(strings.Join(args, ","))
	return w
}

// OpenTX 开启事务
func (w *UpdateWrapper[T]) OpenTX() *UpdateWrapper[T] {
	w.isTx = true
	return w
}

// DeleteById 根据id删除
func (w *UpdateWrapper[T]) DeleteById(id interface{}) (int64, error) {
	t := new(T)
	return w.DB.ID(id).Delete(t)
}

// Delete 删除数据
func (w *UpdateWrapper[T]) Delete() (int64, error) {
	return w.DB.Delete(new(T))
}
