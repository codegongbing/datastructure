package sqList

import "errors"

type Lister interface {
	Size() int
	IsFull() bool
	Append(data interface{})
	Insert(index int, data interface{}) error
	Delete(index int) error
	Clear()
	Set(index int, data interface{}) error
	Get(index int) (interface{}, error)
}

type SqList struct {
	Data     []interface{}
	Length   int
	Capacity int
}

// TODO 初始化表
func NewList(capacity int) (list *SqList) {
	list = new(SqList)
	list.Data = make([]interface{}, capacity)
	list.Length = 0
	list.Capacity = capacity
	return
}

// TODO 表大小
func (sqList *SqList) Size() int {
	return sqList.Length
}

// TODO 判满
func (sqList *SqList) IsFull() bool {
	if sqList.Length >= sqList.Capacity {
		return true
	}
	return false
}

// TODO 扩容
func (sqList *SqList) expansion() {
	newList := NewList(sqList.Capacity * 2)
	for i := 0; i < sqList.Length; i++ {
		newList.Data[i] = sqList.Data[i]
	}
	sqList.Length += newList.Length
	sqList.Capacity = newList.Capacity
	sqList.Data = newList.Data
}

// TODO 缩容
func (sqList *SqList) shrink() {
	newList := NewList(sqList.Capacity / 2)
	for i := 0; i < sqList.Length; i++ {
		newList.Data[i] = sqList.Data[i]
	}
	sqList.Capacity = newList.Capacity
	sqList.Data = newList.Data
}

// TODO 追加数据
func (sqList *SqList) Append(data interface{}) {
	if sqList.IsFull() {
		sqList.expansion()
	}
	sqList.Data[sqList.Length] = data
	sqList.Length++
}

// TODO 添加数据
func (sqList *SqList) Insert(index int, data interface{}) error {
	if index < 1 || index > sqList.Length {
		return errors.New("索引错误")
	}
	if sqList.IsFull() {
		sqList.expansion()
	}
	for i := sqList.Length; i >= index; i-- {
		sqList.Data[i] = sqList.Data[i-1]
	}
	sqList.Data[index-1] = data
	return nil
}

// TODO 删除数据
func (sqList *SqList) Delete(index int) error {
	if index < 1 || index > sqList.Length {
		return errors.New("索引错误")
	}
	if sqList.Length <= sqList.Capacity {
		sqList.shrink()
	}
	for i := index; i < sqList.Length; i++ {
		sqList.Data[i-1] = sqList.Data[i]
	}
	sqList.Data[sqList.Length-1] = nil
	return nil
}

// TODO 清空表
func (sqList *SqList) Clear() {
	sqList.Data = make([]interface{}, 0, sqList.Capacity)
	sqList.Length = 0
}

// TODO 更改值
func (sqList *SqList) Set(index int, data interface{}) error {
	if index < 1 || index > sqList.Length {
		return errors.New("索引错误")
	}
	sqList.Data[index-1] = data
	return nil
}

// TODO 获取值
func (sqList *SqList) Get(index int) (interface{}, error) {
	if index < 1 || index > sqList.Length {
		return nil, errors.New("索引错误")
	}
	return sqList.Data[index-1], nil
}
