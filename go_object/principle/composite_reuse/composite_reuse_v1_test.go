package composite_reuse

import (
	"testing"
	"tortoise/go_object/principle/composite_reuse/bad"
	"tortoise/go_object/principle/composite_reuse/entity"
	"tortoise/go_object/principle/composite_reuse/good"
)

func TestVerifyBad(t *testing.T) {
	fnCallAndLog := func(fn func() (error, int)) {
		e, rows := fn()
		if e != nil {
			t.Errorf("error = %s", e.Error())
		} else {
			t.Logf("affected rows = %v", rows)
		}
	}

	product := entity.NewProduct(1, "手机", 1000)
	bd := bad.NewBadProductDAO("database connection url", "sa", "123")
	fnCallAndLog(func() (error, int) {
		return bd.Insert(product)
	})
	fnCallAndLog(func() (error, int) {
		return bd.Update(product)
	})
	fnCallAndLog(func() (error, int) {
		return bd.Delete(product.ID)
	})
}

func TestVerifyGood(t *testing.T) {
	fnCallAndLog := func(fn func() (error, int)) {
		e, rows := fn()
		if e != nil {
			t.Errorf("error = %s", e.Error())
		} else {
			t.Logf("affected rows = %v", rows)
		}
	}

	con := good.NewGoodMysqlConnection("database connection url", "sa", "123")
	gd := good.NewProductDAO()
	gd.SetDBConnection(con)

	product := entity.NewProduct(1, "手机", 1000)
	fnCallAndLog(func() (error, int) {
		return gd.Insert(product)
	})
	fnCallAndLog(func() (error, int) {
		return gd.Update(product)
	})
	fnCallAndLog(func() (error, int) {
		return gd.Delete(product.ID)
	})
}
