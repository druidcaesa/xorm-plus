package xormplus

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"strings"
)

type QueryWrapper[T any] struct {
	db *xorm.Session
}

func NewQueryWrapper[T any]() *QueryWrapper[T] {
	w := &QueryWrapper[T]{
		db: Engine.NewSession(),
	}
	w.db.Table(new(T))
	return w
}

// Like 模糊查询
func (w *QueryWrapper[T]) Like(column string, attribute interface{}) *QueryWrapper[T] {
	w.db = w.db.Where(fmt.Sprintf(" %s like '%%%s%%' ", column, attribute))
	return w
}

// Eq 等于查询
func (w *QueryWrapper[T]) Eq(column string, attribute interface{}) *QueryWrapper[T] {
	w.db = w.db.Where(fmt.Sprintf("%s = ?", column), attribute)
	return w
}

// Ne 不等于查询
func (w *QueryWrapper[T]) Ne(column string, attribute interface{}) *QueryWrapper[T] {
	w.db = w.db.Where(fmt.Sprintf("%s <> ?", column), attribute)
	return w
}

// OrderBy 正序排序
func (w *QueryWrapper[T]) OrderBy(column string) *QueryWrapper[T] {
	w.db = w.db.OrderBy(fmt.Sprintf("%s", column))
	return w
}

// OrderByDesc 倒序排序
func (w *QueryWrapper[T]) OrderByDesc(column string) *QueryWrapper[T] {
	w.db = w.db.OrderBy(fmt.Sprintf("%s DESC", column))
	return w
}

// SelectList 查询集合
func (w *QueryWrapper[T]) SelectList() *[]T {
	var list []T
	err := w.db.Find(&list)
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &list
}

// OrEq 或查询
func (w *QueryWrapper[T]) OrEq(column string, attribute interface{}) *QueryWrapper[T] {
	w.db = w.db.Or(fmt.Sprintf("%s = ?", column), attribute)
	return w
}

// OrLike 或模糊查询
func (w *QueryWrapper[T]) OrLike(column string, attribute interface{}) *QueryWrapper[T] {
	w.db = w.db.Or(fmt.Sprintf(" %s like '%%%s%%' ", column, attribute))
	return w
}

// OrEntity 使用实体，根据实体内属性字段进行查询
func (w *QueryWrapper[T]) OrEntity(t *T) *QueryWrapper[T] {
	w.db = w.db.Or(t)
	return w
}

// OrMap map参数进行查询
func (w *QueryWrapper[T]) OrMap(m map[string]interface{}) *QueryWrapper[T] {
	w.db = w.db.Or(m)
	return w
}

// NotIn 条件查询
func (w *QueryWrapper[T]) NotIn(query string, args interface{}) *QueryWrapper[T] {
	w.db = w.db.NotIn(query, args)
	return w
}

// SelectOne 单条查询
func (w *QueryWrapper[T]) SelectOne() (*T, error) {
	var model T
	_, err := w.db.Get(&model)

	if err != nil {
		log.Printf("数据查询发生异常%v\n", err.Error())
		return nil, err
	}
	return &model, nil
}

// GetById 根据id查询数据
func (w *QueryWrapper[T]) GetById(id interface{}) *T {
	var model T
	_, err := w.db.ID(id).Get(&model)
	if err != nil {
		log.Printf("数据库查询异常%v\n", err.Error())
		return nil
	}
	return &model
}

// Select 设置查询字段
func (w *QueryWrapper[T]) Select(q ...string) *QueryWrapper[T] {
	w.db = w.db.Cols(strings.Join(q, ","))
	return w
}

// Join 查询
func (w *QueryWrapper[T]) Join(joinOperator string, tableName string, condition string) *QueryWrapper[T] {
	w.db = w.db.Join(joinOperator, tableName, condition)
	return w
}

func (w *QueryWrapper[T]) As(a string) *QueryWrapper[T] {
	w.db = w.db.Alias(a)
	return w
}
