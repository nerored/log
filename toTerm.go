//
// toTerm.go
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
package log

import (
	"bytes"
	"io"
	"os"
)

type termWriter struct {
	infoBuffer *bytes.Buffer
	erroBuffer *bytes.Buffer
}

func newTermWriter() Writer {
	return &termWriter{
		infoBuffer: new(bytes.Buffer),
		erroBuffer: new(bytes.Buffer),
	}
}

func (t *termWriter) init() {

}

func (t *termWriter) exit() {

}

func (t *termWriter) needColor() bool {
	return true
}

func (t *termWriter) info() io.Writer {
	return t.infoBuffer
}

func (t *termWriter) erro() io.Writer {
	return t.erroBuffer
}

func (t *termWriter) reflush() {
	if t.infoBuffer != nil && t.infoBuffer.Len() > 0 {
		io.Copy(os.Stdout, t.infoBuffer)
		t.infoBuffer.Reset()
	}

	if t.erroBuffer != nil && t.erroBuffer.Len() > 0 {
		io.Copy(os.Stderr, t.erroBuffer)
		t.erroBuffer.Reset()
	}
}
