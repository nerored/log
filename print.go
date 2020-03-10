//
// log/print.go
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
	"runtime"
	"strconv"
	"strings"
	"time"
)

type repeater struct {
	out1, out2 io.Writer
}

func (r *repeater) Write(p []byte) (n int, err error) {
	if r.out1 != nil {
		n, err = r.out1.Write(p)
	}

	if r.out2 != nil {
		n, err = r.out2.Write(p)
	}

	return
}

type printer struct {
	timeFormat string
	writerList []io.Writer
}

var (
	sharedPrinter printer
)

func init() {
	SetTimeFormat("2006/01/02 15:04:05")
}

func SetTimeFormat(format string) {
	sharedPrinter.timeFormat = format
}

func AddWriter(writers ...io.Writer) {
	for _, w := range writers {
		if w == nil {
			continue
		}

		sharedPrinter.writerList = append(sharedPrinter.writerList, w)
	}
}

func (this *printer) each(callback func(io.Writer)) {
	if callback == nil || len(this.writerList) <= 0 {
		return
	}

	for _, writer := range this.writerList {
		callback(writer)
	}
}

func (this *printer) Write(p []byte) (n int, err error) {
	for _, writer := range this.writerList {
		if writer == nil {
			continue
		}

		writer.Write(p)
	}

	return len(p), nil
}

func (this *printer) newRoot(logLevel LogLv) (root Combo) {
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

func (this *printer) makeChan(root *Combo, args []interface{}) {
	if root == nil || len(args) == 0 {
		return
	}

	for _, arg := range args {
		chanObj, ok := arg.(*Combo)

		if !ok || chanObj == nil {
			continue
		}

		chanObj.linkTo(root)
	}
}

func (this *printer) printHeadInfo(writer io.Writer, logLevel LogLv, flags PrintFlag) {
	if writer == nil {
		return
	}

	this.printTimeInfo(writer, flags)

	var prefix, suffix string

	if flags&PRINT_LEVELAB > 0 {
		switch logLevel {
		case LOG_LEVEL_DEBU:
			prefix = "[debu"
			suffix = "] "
		case LOG_LEVEL_INFO:
			prefix = "[info"
			suffix = "] "
		case LOG_LEVEL_TRAC:
			prefix = "[trac"
			suffix = "] "
		case LOG_LEVEL_WARN:
			prefix = "[warn"
			suffix = "] "
		case LOG_LEVEL_ERRO:
			prefix = "[erro"
			suffix = "] "
		case LOG_LEVEL_FATA:
			prefix = "[fata"
			suffix = "] "
		default:
			if flags&PRINT_DEFINE > 0 {
				prefix = "["
				suffix = "] "
			}
		}
	}

	if len(prefix) > 0 {
		fmt.Fprintf(writer, prefix)
	}
	this.printFileInfo(writer, flags)
	this.printFuncName(writer, flags)
	if len(suffix) > 0 {
		fmt.Fprintf(writer, suffix)
	}
}

func (this *printer) printTimeInfo(writer io.Writer, flags PrintFlag) {
	if writer == nil || flags&PRINT_TIMELAB <= 0 {
		return
	}

	fmt.Fprintf(writer, time.Now().Format(this.timeFormat))
}

const LOC_STACK_DEPTH = 4

func (this *printer) printFileInfo(writer io.Writer, flags PrintFlag) {
	if writer == nil || flags&PRINT_FILELOC <= 0 {
		return
	}

	_, fileName, line, ok := runtime.Caller(LOC_STACK_DEPTH)

	if !ok {
		return
	}

	fmt.Fprintf(writer, " ")
	fmt.Fprintf(writer, fileName)
	fmt.Fprintf(writer, ":")
	fmt.Fprintf(writer, strconv.Itoa(line))
}

func (this *printer) printFuncName(writer io.Writer, flags PrintFlag) {
	if writer == nil || flags&PRINT_FILELOC <= 0 {
		return
	}

	pc, _, _, ok := runtime.Caller(LOC_STACK_DEPTH)

	if !ok {
		return
	}

	if funcPC := runtime.FuncForPC(pc); funcPC != nil {
		fmt.Fprintf(writer, " ")
		fmt.Fprintf(writer, funcPC.Name())
	}
}

const STD_STACK_DEPTH = 3

func (this *printer) printStackDep(writer io.Writer, flags PrintFlag) {
	if writer == nil || flags&PRINT_STACKIN <= 0 {
		return
	}

	depLen := int(flags & PRINT_STACKDP)

	if depLen <= 0 {
		return
	}

	fmt.Fprintf(writer, "-------- stack info\n")
	for add := 0; add < depLen; add++ {
		pc, fileName, line, ok := runtime.Caller(add + STD_STACK_DEPTH)

		if !ok {
			continue
		}

		fmt.Fprintf(writer, "\t%d: %s:%d ", add, fileName, line)

		funcPC := runtime.FuncForPC(pc)

		if funcPC == nil {
			fmt.Fprintf(writer, "unknow")
		} else {
			fmt.Fprintf(writer, funcPC.Name())
		}

		fmt.Fprintf(writer, "\n")
	}
}

func (this *printer) print(stdIO io.Writer, logLevel LogLv, flags PrintFlag, format string, args ...interface{}) {
	if len(format) <= 0 && len(args) <= 0 {
		return
	}

	majorW := repeater{

		out1: stdIO,
		out2: this,
	}

	root := this.newRoot(logLevel)

	this.makeChan(&root, args)

	root.begin(majorW.out1)
	defer root.end(majorW.out1)

	this.printHeadInfo(&majorW, logLevel, flags)

	//	fmt.Fprintf(majorW.out2, format, args...)
	fmt.Fprintf(majorW.out1, format, args...)

	if !strings.HasSuffix(format, "\n") {
		fmt.Fprintf(&majorW, "\n")
	}

	this.printStackDep(&majorW, flags)
}
