package Jeffail

import (
	"log"
	"testing"
	"time"

	"github.com/Jeffail/tunny"
)


func TestVerify(t *testing.T)  {
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		go pool.Process(i)
	}
	time.Sleep(time.Second * 3)
}
