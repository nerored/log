//
// log.go
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

func Debu(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_DEBU, PRINT_DEBUG, format, args...)
}

func Trac(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_TRAC, PRINT_UTRACE, format, args...)
}

func Info(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_INFO, PRINT_DEFAULT, format, args...)
}

func Warn(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_WARN, PRINT_DEFINE, format, args...)
}

func Erro(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_ERRO, PRINT_DEFINE, format, args...)
}

func Fata(format string, args ...interface{}) {
	sharedPrinter.print(LOG_LEVEL_FATA, PRINT_DEFINE, format, args...)
}

func Ulog(level LogLv, flags PrintFlag, format string, args ...interface{}) {
	sharedPrinter.print(level, flags, format, args...)
}
