//
// log_test.go
// Copyright (c) 2019 nerored <nero_stellar@icloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
package log

import (
	"fmt"
	"testing"
)

func BenchmarkLog1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ulog(LOG_LEVEL_USER, PRINT_DEFAULT, "hello %v", "world")
	}
}

func BenchmarkLog2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Erro("hello world")
	}
}

func BenchmarkLog3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fata("hello world %v %v", NewCombo("lalala", FGC_YELLOW), NewCombo("ohohoh", FMT_BLINK))
	}
}

func BenchmarkLog4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ulog(LOG_LEVEL_USER, PRINT_STACKIN|5, "asdsad")
	}
}

func BenchmarkLog5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello %v", "world")
	}
}

func BenchmarkLog6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("hello %v", "world")
	}
}

func BenchmarkLog7(b *testing.B) {
	type a struct {
		a int
		b int
		c int
	}

	for i := 0; i < b.N; i++ {
		fmt.Printf("hello %+v", a{
			a: 100,
			b: 2200,
			c: 1231,
		})
	}
}
