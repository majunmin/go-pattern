// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-24 20:55
package singleton

import (
	"testing"

	"github.com/majunmin/go-pattern/creator/singleton/lock"
)

//➜  singleton git:(main) ✗ go test -bench=. -run=none
//goos: darwin
//goarch: amd64
//pkg: github.com/majunmin/go-patttern/creator/singleton
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkLazy-12        666121430                1.774 ns/op
//BenchmarkLock-12        518335168                2.295 ns/op
//BenchmarkLock2-12       585673497                2.032 ns/op
//PASS
//ok      github.com/majunmin/go-patttern/creator/singleton       4.337s

func BenchmarkLazy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstance()
	}
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lock.GetInstance()
	}
}

func BenchmarkLock2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lock.GetInstance2()
	}
}
