//
// log/color_test.go
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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColorSet(t *testing.T) {
	Convey("cli fg color test", t, func() {
		for _, color := range []int8{
			FGC_DEFAULT,
			FGC_BLACK,
			FGC_RED,
			FGC_GREEN,
			FGC_YELLOW,
			FGC_BLUE,
			FGC_MAGENTA,
			FGC_CYAN,
			FGC_LIGHTGREY,
			FGC_DARKGREY,
			FGC_LIGHTRED,
			FGC_LIGHTGREEN,
			FGC_LIGHTYELLOW,
			FGC_LIGHTBLUE,
			FGC_LIGHTMAGENTA,
			FGC_LIGHTCYAN,
			FGC_LIGHTWHITE,
		} {
			buffer := new(strings.Builder)
			So(buffer, ShouldNotBeNil)
			obj := NewColorWriter(color)
			So(obj, ShouldNotBeNil)
			obj.write(buffer, "吾有一言。曰「問天地好在」。")
			fmt.Println(buffer.String())
		}
	})
}

func TestFormatSet(t *testing.T) {
	Convey("cli format output test", t, func() {
		for _, format := range []int8{
			FMT_BOLD,
			FMT_DIM,
			FMT_UNDERLINED,
			FMT_BLINK,
			FMT_MINVERTED,
			FMT_HIDDEN,
		} {
			buffer := new(strings.Builder)
			So(buffer, ShouldNotBeNil)
			obj := NewColorWriter(format)
			So(obj, ShouldNotBeNil)
			obj.write(buffer, "落霞与孤鹜齐飞，秋水共长天一色")
			fmt.Println(buffer.String())
		}
	})
}

func TestBGColorSet(t *testing.T) {
	Convey("cli bg color test", t, func() {
		for _, color := range []int8{
			BGC_DEFAULT,
			BGC_BLACK,
			BGC_RED,
			BGC_GREEN,
			BGC_YELLOW,
			BGC_BLUE,
			BGC_MAGENTA,
			BGC_CYAN,
			BGC_LIGHTGREY,
			BGC_DARKGREY,
			BGC_LIGHTRED,
			BGC_LIGHTGREEN,
			BGC_LIGHTYELLOW,
			BGC_LIGHTBLUE,
			BGC_LIGHTMAGENTA,
			BGC_LIGHTCYAN,
			BGC_LIGHTWHITE,
		} {
			buffer := new(strings.Builder)
			So(buffer, ShouldNotBeNil)
			obj := NewColorWriter(color)
			So(obj, ShouldNotBeNil)
			obj.write(buffer, "仿佛兮若轻云之蔽月,飘飘兮若流风之回雪")
			fmt.Println(buffer.String())
		}
	})
}
