//
// benchmark_test.go
// Copyright (c) 2020 nerored <nero_stellar@icloud.com>
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

func init() {
	InitLog("", true)
}

func BenchmarkLogNoComboInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("benchmark info")
	}
}

func BenchmarkLogNoComboWarn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Warn("benchmark info")
	}
}

func BenchmarkLogNoComboTrac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trac("benchmark info")
	}
}

func BenchmarkLog1ComboInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("benchmark info with %v", NewCombo("combo", FGC_LIGHTBLUE))
	}
}

func BenchmarkLog1ComboWarn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Warn("benchmark info with %v", NewCombo("combo", FGC_LIGHTBLUE))
	}
}

func BenchmarkLog1ComboTrac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trac("benchmark info with %v", NewCombo("combo", FGC_LIGHTBLUE))
	}
}

func BenchmarkFMTNoCombo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("benchmark info")
	}
}

func BenchmarkFMT1Combo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("benchmark info with %v\n", "combo")
	}
}
