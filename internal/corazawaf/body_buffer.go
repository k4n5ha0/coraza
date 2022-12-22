// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package corazawaf

import (
	"bytes"
	"io"
	"math"
	"os"

	"github.com/corazawaf/coraza/v3/internal/environment"
	"github.com/corazawaf/coraza/v3/types"
)

// BodyBuffer is used to read RequestBody and ResponseBody objects
// It will handle memory usage for buffering and processing
// It implements io.Copy(bodyBuffer, someReader) by inherit io.Writer
type BodyBuffer struct {
	io.Writer
	options types.BodyBufferOptions
	buffer  *bytes.Buffer
	writer  *os.File
	length  int64
}

const NotOverflow = int64(-1)

// Write appends data to the body buffer by chunks
// You may dump io.Readers using io.Copy(br, reader)
func (br *BodyBuffer) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}

	// If overflow already reached do not write anything.
	if br.length == math.MaxInt64 {
		return 0, nil
	}

	maxWritingDataLenBeforeOverflow := NotOverflow
	var l int64

	// Overflow check
	if br.length > (math.MaxInt64 - int64(len(data))) {
		// Overflow, buffer length will always be at most MaxInt
		l = math.MaxInt64
		maxWritingDataLenBeforeOverflow = math.MaxInt64 - br.length
	} else {
		// No Overflow
		l = br.length + int64(len(data))
	}
	// Check if memory or disk limits are reached
	// Even if Overflow is explicitly checked, MemoryLimit real limits are below maxInt and machine dependenent.
	// bytes.Buffer growth is platform dependent with a growth rate capped at 2x. If the buffer can't grow it will panic with ErrTooLarge.
	// See https://github.com/golang/go/blob/go1.19.4/src/bytes/buffer.go#L117 and https://go-review.googlesource.com/c/go/+/349994
	// Local tests show these buffer limits:
	// 32-bit machine: 2147483647 (2^30, 1GiB)
	// 64-bit machine: 34359738368 (2^35, 32GiB) (Not reached the ErrTooLarge panic, the OS triggered an oom)
	if l > br.options.MemoryLimit {
		if environment.IsTinyGo {
			// TinyGo: Bytes beyond MemoryLimit are not written
			maxWritingDataLen := br.options.MemoryLimit - br.length
			if maxWritingDataLen == 0 {
				return 0, nil
			}
			br.length = br.options.MemoryLimit
			return br.buffer.Write(data[:maxWritingDataLen])
		} else {
			// Default: bytes are buffered to disk
			if br.writer == nil {
				br.writer, err = os.CreateTemp(br.options.TmpPath, "body*")
				if err != nil {
					return 0, err
				}
				// We dump the previous buffer
				if _, err := br.writer.Write(br.buffer.Bytes()); err != nil {
					return 0, err
				}
				br.buffer.Reset()
			}
			// Total (disk) limit is checked
			if l > br.options.Limit {
				// Total limit exceeded, bytes beyond Limit are not buffered
				maxWritingDataLen := br.options.Limit - br.length
				if maxWritingDataLen == 0 {
					return 0, nil
				}
				br.length = br.options.Limit
				return br.writer.Write(data[:maxWritingDataLen])
			} else {
				// Total limit not exceeded, bytes are buffered to disk
				br.length = l
				if maxWritingDataLenBeforeOverflow == NotOverflow {
					return br.writer.Write(data)
				} else {
					return br.writer.Write(data[:maxWritingDataLenBeforeOverflow])
				}
			}
		}
	}

	br.length = l
	if maxWritingDataLenBeforeOverflow == NotOverflow {
		return br.buffer.Write(data)
	} else {
		return br.buffer.Write(data[:maxWritingDataLenBeforeOverflow])
	}
}

// Reader Returns a working reader for the body buffer in memory or file
func (br *BodyBuffer) Reader() (io.Reader, error) {
	if environment.IsTinyGo || br.writer == nil {
		return bytes.NewReader(br.buffer.Bytes()), nil
	}
	if _, err := br.writer.Seek(0, 0); err != nil {
		return nil, err
	}
	return br.writer, nil
}

// Size returns the current size of the body buffer
func (br *BodyBuffer) Size() int64 {
	return br.length
}

// Reset will reset buffers and delete temporary files
func (br *BodyBuffer) Reset() error {
	br.buffer.Reset()
	br.length = 0
	if !environment.IsTinyGo && br.writer != nil {
		w := br.writer
		br.writer = nil
		if err := w.Close(); err != nil {
			return err
		}
		return os.Remove(w.Name())
	}

	return nil
}

// NewBodyBuffer Initializes a body reader
// After writing memLimit bytes to the memory buffer, data will be
// written to a temporary file
// Temporary files will be written to tmpDir
func NewBodyBuffer(options types.BodyBufferOptions) *BodyBuffer {
	return &BodyBuffer{
		options: options,
		buffer:  &bytes.Buffer{},
	}
}
