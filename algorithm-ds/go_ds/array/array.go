package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []interface{}
	length uint
}

func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]interface{}, capacity, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index uint) (interface{}, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of range index")
	}
	return this.data[index], nil
}

/**
1: index 插入的位置
2：v插入的值
*/
func (this *Array) Insert(index uint, v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("out of range index")
	}

	if index != this.Len() && this.isIndexOutOfRange(index) {
		return errors.New("out of range index")
	}

	for l := this.Len(); l > index; l-- {
		this.data[l] = this.data[l-1]
	}
	this.data[index] = v
	this.length++

	return nil
}

func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

func (this *Array) Delete(index uint) (interface{}, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of range index")
	}

	v := this.data[index]
	for l := index; l < this.Len()-1; l++ {
		this.data[l] = this.data[l+1]
	}
	this.length--
	return v, nil
}

func (p *Array) Print() {
	var format string
	for i := uint(0); i < p.Len(); i++ {
		format += fmt.Sprintf("|%+v", p.data[i])
	}
	fmt.Println(format)
}
