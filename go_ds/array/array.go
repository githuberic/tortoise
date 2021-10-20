package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (p *Array) Len() uint {
	return p.length
}

//判断索引是否越界
func (p *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(p.data)) {
		return true
	}
	return false
}

func (p *Array) Find(index uint) (int, error) {
	if p.isIndexOutOfRange(index) {
		return 0, errors.New("out of range index")
	}
	return p.data[index], nil
}

func (p *Array) Insert(index uint, v int) error {
	if p.Len() == uint(cap(p.data)) {
		return errors.New("out of range index")
	}

	if index != p.Len() && p.isIndexOutOfRange(index) {
		return errors.New("out of range index")
	}

	for l := p.Len(); l > index; l-- {
		p.data[l] = p.data[l-1]
	}
	p.data[index] = v
	p.length++

	return nil
}

func (p *Array) InsertToTail(v int) error {
	return p.Insert(p.Len(), v)
}

func (p *Array) Delete(index uint) (int, error) {
	if p.isIndexOutOfRange(index) {
		return 0, errors.New("out of range index")
	}

	v := p.data[index]
	for l := index; l < p.Len()-1; l++ {
		p.data[l] = p.data[l+1]
	}
	p.length--
	return v, nil
}

func (p *Array) Print() {
	var format string
	for i := uint(0); i < p.Len(); i++ {
		format += fmt.Sprintf("|%+v", p.data[i])
	}
	fmt.Println(format)
}
