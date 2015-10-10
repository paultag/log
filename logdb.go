//

package main

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

//
func Dirpath(when time.Time) string {
	return when.Format("2006/01/02")
}

//
func Rootpath(root string, when time.Time) string {
	return path.Join(root, Dirpath(when))
}

//
func Create(root string, when time.Time) error {
	return os.MkdirAll(Rootpath(root, when), 0700)
}

//
func Logit(root string, when time.Time, log string) error {
	logPath := path.Join(Rootpath(root, when), "log")
	_, err := os.Stat(logPath)
	newFile := false
	if os.IsNotExist(err) {
		newFile = true
		fd, err := os.Create(logPath)
		if err != nil {
			return err
		}
		fd.Close()
	} else if err != nil {
		return err
	}

	fd, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	writer := bufio.NewWriter(fd)
	if !newFile {
		writer.Write([]byte{12, '\n'})
		writer.Flush()
	}
	writer.Write([]byte(log))
	writer.Write([]byte("\n"))
	writer.Flush()
	return nil
}

//
func Readit(root string, when time.Time) []string {
	ret := []string{}

	logPath := path.Join(Rootpath(root, when), "log")
	fd, err := os.Open(logPath)
	if err != nil {
		return []string{}
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString(byte(12))
		ret = append(ret, strings.Trim(line, "\n\r \t\x0C"))
		if err == io.EOF {
			break
		}
	}

	return ret
}
