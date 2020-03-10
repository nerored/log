//
// writer.go
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
	"io"
	"sync"
)

type Writer interface {
	init()
	exit()
	needColor() bool
	info() io.Writer
	erro() io.Writer
}

type writerChan struct {
	writers []io.Writer
}

func (w *writerChan) isEmpty() bool {
	return len(w.writers) == 0
}

func (w *writerChan) free() {
	w.writers = w.writers[:0]
	writerChanPool.Put(w)
}

func (w *writerChan) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		if writer == nil {
			continue
		}

		writer.Write(p)
	}

	return
}

var writerChanPool = sync.Pool{
	New: func() interface{} {
		return new(writerChan)
	},
}

func buildIOChanByWriters(level LogLv, writers []Writer) (newChan *writerChan) {
	newChan = writerChanPool.Get().(*writerChan)

	if cap(newChan.writers) < len(writers) {
		newChan.writers = make([]io.Writer, 0, len(writers))
	}

	switch level {
	case LOG_LEVEL_ERRO, LOG_LEVEL_FATA:
		for _, writer := range writers {
			if writer == nil {
				continue
			}
			newChan.writers = append(newChan.writers, writer.erro())
		}
	default:
		for _, writer := range writers {
			if writer == nil {
				continue
			}
			newChan.writers = append(newChan.writers, writer.info())
		}
	}

	return
}

func getMajorChan(level LogLv, colored, nocolor []Writer) (major repeater) {
	return repeater{
		colored: buildIOChanByWriters(level, colored),
		nocolor: buildIOChanByWriters(level, nocolor),
	}
}

func freeMajorChan(major *repeater) {
	if major == nil {
		return
	}

	if major.colored != nil {
		major.colored.free()
	}

	if major.nocolor != nil {
		major.nocolor.free()
	}
}
