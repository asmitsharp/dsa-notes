package test

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitching(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

// ------ Results ------ //
// > GO111MODULE=off go test -run '^$' -bench=. -cpu=1 contex_switch_test.go
// goos: darwin
// goarch: arm64
// cpu: Apple M1
// BenchmarkContextSwitching        9627457               125.5 ns/op
// PASS
// ok      command-line-arguments  1.714s
