package io

import (
	"io"
	"sync/atomic"
)

type SpeedWriter struct {
	W int64
	F io.Writer
}

func NewSpeedWriter(f io.Writer) *SpeedWriter {
	return &SpeedWriter{F: f}
}

type SpeedReader struct {
	R int64
	F io.Reader
}

func NewSpeedReader(f io.Reader) *SpeedReader {
	return &SpeedReader{F: f}
}

func (s *SpeedWriter) Write(p []byte) (n int, err error) {
	write, err := s.F.Write(p)
	atomic.AddInt64(&s.W, int64(write))
	return write, err
}

func (s *SpeedReader) Read(p []byte) (n int, err error) {
	read, err := s.F.Read(p)
	atomic.AddInt64(&s.R, int64(read))
	return read, err
}
