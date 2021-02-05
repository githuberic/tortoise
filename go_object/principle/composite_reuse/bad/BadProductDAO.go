package bad

import (
	"tortoise/go_object/principle/composite_reuse/entity"
)

type BadProductDAO struct {
	BadDBConnection
}

func NewBadProductDAO(url string, uid string, pwd string) *BadProductDAO {
	return &BadProductDAO{*NewBadDBConnection(url, uid, pwd)}
}

func (p *BadProductDAO) Insert(it *entity.Product) (error, int) {
	return p.Execute("insert into product(id,name,price) values(?, ?, ?)", it.ID, it.Name, it.Price)
}

func (p *BadProductDAO) Update(it *entity.Product) (error, int) {
	return p.Execute("update product set name=? price=? where id=?", it.Name, it.Price, it.ID)
}

func (p *BadProductDAO) Delete(id int) (error, int) {
	return p.Execute("delete from product where id=?", id)
}