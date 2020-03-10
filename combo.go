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
	"sync"
)

type Combo struct {
	color
	coloredCount int
	data         interface{}
}

var comboPool = sync.Pool{
	New: func() interface{} {
		return new(Combo)
	},
}

func NewCombo(data interface{}, attrs ...int) (c *Combo) {
	c = comboPool.Get().(*Combo)

	c.data = data
	c.color.setAttrs(attrs...)
	return
}

func (c *Combo) linkTo(comboChan *Combo) {
	if comboChan == nil {
		return
	}

	c.color.linkTo(&comboChan.color)
	c.addColoredChange(comboChan.coloredCount)
}

func (c *Combo) free() {
	c.data = nil
	c.coloredCount = 0
	c.color.attributes = nil
	c.color.colorChan = nil

	comboPool.Put(c)
}

func (c *Combo) addColoredChange(delta int) {
	c.coloredCount += delta
}

func (c *Combo) String() string {
	writing := func() (buffer *strings.Builder) {
		buffer = new(strings.Builder)

		c.begin(buffer)
		defer c.end(buffer)

		fmt.Fprintf(buffer, "%v", c.data)
		return
	}

	return writing().String()
}

func freeCombos(args []interface{}) {
	for _, arg := range args {
		combo, ok := arg.(*Combo)

		if !ok || combo == nil {
			continue
		}

		combo.free()
	}
}

//------ combo root maker

func newRoot(logLevel LogLv) (root Combo) {
	switch logLevel {
	case LOG_LEVEL_DEBU:
		root.setAttrs(FGC_LIGHTCYAN)
	case LOG_LEVEL_INFO:
		root.setAttrs(FGC_DEFAULT)
	case LOG_LEVEL_TRAC:
		root.setAttrs(FGC_LIGHTYELLOW, FMT_UNDERLINED)
	case LOG_LEVEL_WARN:
		root.setAttrs(FGC_YELLOW)
	case LOG_LEVEL_ERRO:
		root.setAttrs(FGC_RED)
	case LOG_LEVEL_FATA:
		root.setAttrs(FGC_LIGHTWHITE, BGC_RED)
	}

	return
}
