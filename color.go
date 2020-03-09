//
// log/color.go
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
	"io"
)

const (
	FMT_RESET = 0

	FMT_BOLD       = 1
	FMT_DIM        = 2
	FMT_UNDERLINED = 4
	FMT_BLINK      = 5
	FMT_MINVERTED  = 7
	FMT_HIDDEN     = 8

	FGC_DEFAULT      = 39
	FGC_BLACK        = 30
	FGC_RED          = 31
	FGC_GREEN        = 32
	FGC_YELLOW       = 33
	FGC_BLUE         = 34
	FGC_MAGENTA      = 35
	FGC_CYAN         = 36
	FGC_LIGHTGREY    = 37
	FGC_DARKGREY     = 90
	FGC_LIGHTRED     = 91
	FGC_LIGHTGREEN   = 92
	FGC_LIGHTYELLOW  = 93
	FGC_LIGHTBLUE    = 94
	FGC_LIGHTMAGENTA = 95
	FGC_LIGHTCYAN    = 96
	FGC_LIGHTWHITE   = 97

	BGC_DEFAULT      = 49
	BGC_BLACK        = 40
	BGC_RED          = 41
	BGC_GREEN        = 42
	BGC_YELLOW       = 43
	BGC_BLUE         = 44
	BGC_MAGENTA      = 45
	BGC_CYAN         = 46
	BGC_LIGHTGREY    = 47
	BGC_DARKGREY     = 100
	BGC_LIGHTRED     = 101
	BGC_LIGHTGREEN   = 102
	BGC_LIGHTYELLOW  = 103
	BGC_LIGHTBLUE    = 104
	BGC_LIGHTMAGENTA = 105
	BGC_LIGHTCYAN    = 106
	BGC_LIGHTWHITE   = 107
)

type Color struct {
	attributes []int
	preCrAttrs []int
}

func (this *Color) setAttributes(attrs []int) {
	this.attributes = attrs
}

func (this *Color) setPreCrAttrs(attrs []int) {
	this.preCrAttrs = attrs
}

func (this *Color) empty() bool {
	return len(this.preCrAttrs) <= 0 && len(this.attributes) <= 0
}

func (this *Color) writeH(writer io.Writer, beforeLen int, selfList []int) {
	if writer == nil || len(selfList) <= 0 {
		return
	}

	start := 0

	if beforeLen <= 0 {
		fmt.Fprintf(writer, "%d", selfList[0])
		start = 1
	}

	for ; start < len(selfList); start++ {
		fmt.Fprintf(writer, ";%d", selfList[start])
	}
}

func (this *Color) writeE(writer io.Writer, beforeLen int) {
	if writer == nil || beforeLen <= 0 {
		return
	}

	fmt.Fprintf(writer, "\x1b[%dm", FMT_RESET)
}

func (this *Color) setFormat(writer io.Writer) {
	if writer == nil || this.empty() {
		return
	}

	l := len(this.preCrAttrs)
	this.writeE(writer, l)

	fmt.Fprintf(writer, "\x1b[")
	defer fmt.Fprintf(writer, "m")

	this.writeH(writer, 0, this.preCrAttrs)
	this.writeH(writer, l, this.attributes)
}

func (this *Color) end(writer io.Writer) {
	if writer == nil || this.empty() {
		return
	}

	this.writeE(writer, 1)

	if len(this.preCrAttrs) <= 0 {
		return
	}

	fmt.Fprintf(writer, "\x1b[")
	defer fmt.Fprintf(writer, "m")

	this.writeH(writer, 0, this.preCrAttrs)
}
