package client

import (
	"testing"
	"tortoise/go_learning/ch6/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
}
