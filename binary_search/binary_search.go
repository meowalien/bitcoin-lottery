package binary_search

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// BinarySearchOnDisk performs a binary search on a lexicographically sorted file
// without loading the entire file into memory. The 'target' is the line we want to find.
// Returns true if the line is found, otherwise false.
func BinarySearchOnDisk(file *os.File, target string) (bool, error) {
	stat, err := file.Stat()
	if err != nil {
		return false, err
	}
	fileSize := stat.Size()

	low, high := int64(0), fileSize-1
	for low <= high {
		mid := (low + high) / 2

		line, lineStartOffset, nextLineStartOffset, err := readLineFromOffset(file, mid)
		if err != nil && err != io.EOF {
			return false, err
		}

		cmp := strings.Compare(line, target)
		if cmp == 0 {
			// Found the target line
			return true, nil
		} else if cmp < 0 {
			// line < target, search after this line
			low = nextLineStartOffset
		} else {
			// line > target, search before this line
			high = lineStartOffset - 1
		}
	}

	return false, nil
}

// readLineFromOffset attempts to read a full line starting at or after the given offset.
// It works by seeking to 'pos', then moving backwards until it finds the start of a line,
// and then reading that line. Returns the line, the start offset of that line, and the
// start offset of the next line.
func readLineFromOffset(file *os.File, pos int64) (string, int64, int64, error) {
	if pos < 0 {
		pos = 0
	}
	_, err := file.Seek(pos, io.SeekStart)
	if err != nil {
		return "", 0, 0, err
	}

	// Move backwards until start of line (or beginning of file)
	buf := make([]byte, 1)
	for {
		if pos == 0 {
			// At start of file, so this is the start of a line
			break
		}
		pos--
		_, err := file.Seek(pos, io.SeekStart)
		if err != nil {
			return "", 0, 0, err
		}
		n, err := file.Read(buf)
		if err != nil {
			return "", 0, 0, err
		}
		if n == 1 && buf[0] == '\n' {
			// The line starts after this newline
			pos++
			break
		}
	}

	lineStart := pos
	_, err = file.Seek(pos, io.SeekStart)
	if err != nil {
		return "", 0, 0, err
	}

	reader := bufio.NewReader(file)
	lineBytes, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", 0, 0, err
	}

	// Trim newline character if present
	lineStr := strings.TrimRight(lineBytes, "\n")

	// nextLineStart is current lineStart + line length + newline
	// If EOF is encountered without a trailing newline, then next line start is lineStart + len(line)
	nextLineStart := lineStart + int64(len(lineStr)) + 1
	if err == io.EOF {
		nextLineStart = lineStart + int64(len(lineStr))
	}

	return lineStr, lineStart, nextLineStart, nil
}
