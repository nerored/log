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
	color
	data interface{}
}

func NewCombo(data interface{}, attrs ...int) (c *Combo) {
	c = new(Combo)
	c.data = data
	c.color.setAttrs(attrs)
	return
}

func (c *Combo) linkTo(comboChan *Combo) {
	if comboChan == nil {
		return
	}

	c.color.linkTo(&comboChan.color)
}

func (c *Combo) String() string {
	var builder strings.Builder

	c.start(&builder)
	defer c.end(&builder)

	fmt.Fprintf(&builder, "%v", c.data)

	return builder.String()
}
