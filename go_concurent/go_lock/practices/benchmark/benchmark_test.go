package benchmark

import (
	"sync"
	"testing"
)

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup

		for r := 0; r < read*100; r++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}

		for w := 0; w < write*100; w++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B) {
	benchmark(b, &Lock{}, 9, 1)
}

func BenchmarkReadMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 9, 1)
}

func BenchmarkWriteMore(b *testing.B) {
	benchmark(b, &Lock{}, 1, 9)
}

func BenchmarkWriteMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 1, 9)
}

func BenchmarkEqual(b *testing.B) {
	benchmark(b, &Lock{}, 5, 5)
}

func BenchmarkEqualRW(b *testing.B) {
	benchmark(b, &RWLock{}, 5, 5)
}
