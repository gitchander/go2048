package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

var ErrFileLoggerClosed = errors.New("FileLogger is closed")

type FileLogger struct {
	file   *os.File
	bw     *bufio.Writer
	opened bool
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{
		file:   file,
		bw:     bufio.NewWriter(file),
		opened: true}, nil
}

func (fl *FileLogger) Close() error {
	if !fl.opened {
		return ErrFileLoggerClosed
	}
	fl.bw.Flush()
	err := fl.file.Close()
	fl.opened = false
	return err
}

var _ io.Writer = &FileLogger{}

func (fl *FileLogger) Write(data []byte) (n int, err error) {
	if !fl.opened {
		return 0, ErrFileLoggerClosed
	}
	return fl.bw.Write(data)
}

func usageFileLogger() {
	logWriter, err := NewFileLogger("out.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logWriter.Close()
	log.SetOutput(logWriter)

	// do something
}
