package log

import (
	"strings"
	"testing"
)

//
// color_test.go
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

func TestSigleColor(t *testing.T) {
	exs := [][]int{
		{},
		{FMT_BOLD},
		{FGC_BLACK},
		{BGC_BLACK},
		{FMT_BLINK, FGC_BLUE},
		{BGC_CYAN, FGC_BLUE},
	}

	out := []string{
		"",
	}

	for i, ex := range exs {
		if i >= len(out) {
			t.Errorf("[color] no output ex %v", i)
			continue
		}

		var c color
		c.setAttrs(ex...)

		if len(ex) != len(c.attributes) {
			t.Errorf("[color] setAttrs test failed ex %v", 1)
			continue
		}

		if len(ex) > 0 && c.isEmpty() {
			t.Errorf("[color] isEmpty test failed ex %v", i)
			continue
		}

		var builder strings.Builder

		func() {
			c.begin(&builder)
			defer c.end(&builder)
		}()

		if builder.String() != out[i] {
			t.Errorf("[color] buildcolor failed ex %v", i)
			continue
		}
	}
}
