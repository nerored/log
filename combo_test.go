//
// log/combo_test.go
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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTermLog(t *testing.T) {
	Convey("cli term log test", t, func() {
		Info("对酒当歌，人生几何")
		Warn("譬如朝露，去日苦多")
		Debu("慨当以慷，忧思难忘")
		Erro("何以%v？唯%v杜康", NewCombo("解忧", FGC_BLUE), NewCombo("有", BGC_YELLOW, FGC_LIGHTMAGENTA))
		Trac("%v子衿，悠悠我心", NewCombo("青青", FMT_BLINK))
		Fata("但为君故，沉吟至今")
		Info("呦呦%v之苹", NewCombo("鹿鸣，食野", FGC_MAGENTA, FMT_UNDERLINED))
		Warn("我有嘉宾，鼓瑟吹笙")
		Debu("明明如月，何时可掇")
		Erro("忧从中来，不可断绝")
	})
}
