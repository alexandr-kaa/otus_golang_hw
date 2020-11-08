package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrConvertFailed         = errors.New("type conversion error")
	ErrInvalidSize           = errors.New("invalid size param")
)

type copyFileContext interface {
	copy() error
	closeBar()
}

func createCopyFile(offset int64, len int64, fromSrc string, toDst string) (copyFileContext, error) {
	bar := make(chan int64)
	if offset < 0 || len < 0 {
		return nil, ErrInvalidSize
	}
	stat, err := os.Stat(fromSrc)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if stat.Size() == 0 {
		return nil, fmt.Errorf("%w", ErrUnsupportedFile)
	}
	if stat.Size() < offset {
		return nil, fmt.Errorf("%w", ErrOffsetExceedsFileSize)
	}
	if err = checkEmptyName(toDst); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if len > 0 {
		if offset+len > stat.Size() {
			CreateBar(stat.Size()-offset, bar)
		} else {
			CreateBar(len, bar)
		}
	} else {
		CreateBar(stat.Size(), bar)
	}
	file := copyFile{offset: offset, length: len, from: fromSrc, to: toDst, bar: bar}
	return file, nil
}

type copyFile struct {
	offset, length int64
	from           string
	to             string
	bar            chan int64
}

func (f copyFile) closeBar() {
	close(f.bar)
}

func checkEmptyName(name string) error {
	if name == "" {
		return &os.PathError{Err: syscall.ENOENT, Op: "Empty path From", Path: ""}
	}
	return nil
}

func readChan(out chan []byte, errCh chan error, reader io.Reader, len int64) {
	var buffSize int64 = 1024
	if len < buffSize && len > 0 {
		buffSize = len
	}
	if len == 0 {
		len = 1 << 62
	}
	limReader, ok := io.LimitReader(reader, buffSize).(*io.LimitedReader)
	if !ok {
		errCh <- ErrConvertFailed
	}
	defer close(out)
	var total int64 = 0
	for total < len {
		buffer := make([]byte, buffSize)
		length, err := limReader.Read(buffer)
		if !errors.Is(err, io.EOF) && err != nil {
			errCh <- err
			return
		}
		out <- buffer[:length]
		if errors.Is(err, io.EOF) {
			out <- nil
			return
		}
		total += int64(length)
		if total+buffSize > len && len > 0 {
			limReader.N = len - total
		} else {
			limReader.N = buffSize
		}
	}
}

func writeChan(out chan []byte, errCh chan error, writer io.Writer, bar chan int64) {
	defer close(errCh)
	total := 0
	for s := range out {
		buffer := s
		if s == nil {
			errCh <- io.EOF
			return
		}
		total += len(s)
		bar <- int64(total)
		_, err := writer.Write(buffer)
		if err != nil {
			errCh <- err
			return
		}
	}
	errCh <- nil
}

func (f copyFile) copy() error {
	s := make(chan []byte)
	errChan := make(chan error)
	defer close(f.bar)
	defer func() {
		f.bar <- f.length
	}()
	reader, err := os.Open(f.from)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err = reader.Seek(f.offset, io.SeekStart); err != nil {
		return fmt.Errorf("%v", err)
	}
	fileTo, err := os.Create(f.to)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	go readChan(s, errChan, reader, f.length)
	go writeChan(s, errChan, fileTo, f.bar)
	errCh := <-errChan
	if errors.Is(errCh, io.EOF) {
		return nil
	}
	if errCh != nil {
		return fmt.Errorf("%w", errCh)
	}
	return nil
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	// Place your code here
	file, err := createCopyFile(offset, limit, fromPath, toPath)
	if err != nil {
		return err
	}
	err = file.copy()
	return fmt.Errorf("%w", err)
}

type progressBar struct {
	limit int64
}

func (p progressBar) write(pos int64) {
	len := 25.
	c := len / float64(p.limit) * float64(pos)

	if _, err := io.WriteString(os.Stdout, "\r"); err != nil {
		log.Fatal(err)
	}

	str := fmt.Sprintf("[%s", strings.Repeat("*", int(c)))
	if _, err := io.WriteString(os.Stdout, str); err != nil {
		log.Fatal(err)
	}
	str = strings.Repeat(" ", int(len-c))
	if _, err := io.WriteString(os.Stdout, str); err != nil {
		log.Fatal(err)
	}
	if _, err := io.WriteString(os.Stdout, "]"); err != nil {
		log.Fatal(err)
	}

	proc := strconv.FormatFloat(c/len*100, 'f', 1, 64)
	if _, err := io.WriteString(os.Stdout, proc); err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(os.Stdout, "%"); err != nil {
		log.Fatal(err)
	}
}

func CreateBar(limit int64, c <-chan int64) {
	go func() {
		bar := progressBar{limit: limit}
		for s := range c {
			if s > bar.limit {
				bar.write(bar.limit)
				continue
			}
			bar.write(s)
		}
	}()
}
