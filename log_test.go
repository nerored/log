//
// log_test.go
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

import "testing"

func TestPrintInfo(t *testing.T) {
	InitLog("", true)

	Info("hero world this is %v -----l",
		NewCombo("combo with red", FGC_RED))

	Debu("hello world this is %v",
		NewCombo("combo with blue", FGC_BLUE))

	Trac("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))

	Fata("hello world this is %v and %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD),
		NewCombo("combo with2 3", FMT_UNDERLINED))

	Erro("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))

	Warn("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))
}
