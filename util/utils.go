package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	WorkspacePath string
	ConfigPath    string
)

func init() {
	_, currentPath, _, _ := runtime.Caller(0)

	WorkspacePath = filepath.Join(filepath.Dir(currentPath), "../..")
	ConfigPath = filepath.Join(WorkspacePath, "config")
}

func NormalizeFileName(fileName string) (*string, error) {
	unixTime := strconv.Itoa(int(time.Now().Unix()))

	fileName = strings.ToLower(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "_")

	fileName, err := transformChain(fileName)

	if err != nil {
		errMsg := fmt.Sprintf("Fail to transform file name - %s", err.Error())

		return nil, errors.New(errMsg)
	}

	formattedFileName := fmt.Sprintf("%s-%s", unixTime, fileName)

	return &formattedFileName, nil
}

func transformChain(str string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	transformResult, _, err := transform.String(t, str)

	return transformResult, err
}

func GetEndTimeDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func CountCsvLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	var lastChar rune = '\n'
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		if c > 0 {
			lastChar = rune(buf[c-1])
		}

		switch {
		case err == io.EOF:
			if lastChar != '\n' {
				count++
			}
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
