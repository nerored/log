//
// log/combo.go
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
	"strings"
)

type Combo struct {
	data interface{}
	Color
}

func NewCombo(data interface{}, attrs ...int) (c *Combo) {
	c = &Combo{
		data: data,
	}
	c.setAttributes(attrs)
	return
}

func (this *Combo) build(preComboList ...*Combo) {
	for _, combo := range preComboList {
		if combo == nil {
			continue
		}

		this.setPreCrAttrs(combo.attributes)
	}
}

func (this *Combo) String() string {
	var buffer strings.Builder

	this.setFormat(&buffer)
	fmt.Fprintf(&buffer, "%v", this.data)
	this.end(&buffer)

	return buffer.String()
}
