package randomstringutils

import (
	"fmt"
	"testing"
)

func TestRandStringBytesMaskImprSrcSB(t *testing.T) {
	fmt.Println(RandStringBytesMaskImprSrcSB(50))
}

func TestRandStringRunes(t *testing.T) {
	fmt.Println(RandStringRunes(50))
}

func TestRandIntRunes(t *testing.T) {
	fmt.Println(RandIntRunes(5))
}

func BenchmarkRandStringBytesMaskImprSrcSB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcSB(50)
	}
}

func BenchmarkRandStringBytesMaskImprSrcUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcUnsafe(50)
	}
}

func BenchmarkRandStringRunes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStringRunes(50)
	}
}

/*
go test -benchmem -run=none gin_graphql/utils/randomStringUtils -bench=.
goos: darwin
goarch: amd64
pkg: gin_graphql/utils/randomStringUtils
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkRandStringBytesMaskImprSrcSB-8          5985272               222.8 ns/op            64 B/op          1 allocs/op
BenchmarkRandStringBytesMaskImprSrcUnsafe-8      9349088               160.9 ns/op            64 B/op          1 allocs/op
BenchmarkRandStringRunes-8                        113139             11196 ns/op             120 B/op          4 allocs/op
PASS
ok      gin_graphql/utils/randomStringUtils     4.558s
*/
