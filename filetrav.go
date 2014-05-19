/*
Package filetrav: Travel around files easily.

This may exist elsewhere, there's probably an easier way to do this. But it was a fun bit of code to write and I find it useful.

Warning: This is not fully tested as of yet.
*/
package filetrav

import (
    "bytes"
    "io/ioutil"
    "regexp"
)

// FileTraveler is a wrapper for the file being worked with.
type FileTraveler struct {
    lines    [][]byte
    current  []byte
    position int
    bottom   int
}

// NewFileTraveler creates a new FileTraveler from a []byte, like the
// one that would be read in from ioutil.ReadFile.
func NewFileTraveler(data []byte) (traveler FileTraveler, err error) {
    traveler.lines = bytes.Split(data, []byte("\n"))
    traveler.current = traveler.lines[0]
    traveler.bottom = len(traveler.lines) - 1
    return
}

// ReadFileTraveler creates a new FileTraveler from a path to a file on
// disk.
func ReadFileTraveler(file string) (traveler FileTraveler, err error) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }

    return NewFileTraveler(data)
}

// Position within the current file.
func (traveler *FileTraveler) Position() (position int) {
    return traveler.position
}

// GoTo 'n' line in file.
func (traveler *FileTraveler) GoTo(n int) (moved bool) {
    if n == traveler.position || n < 0 || n > traveler.bottom {
        return false
    }

    traveler.current = traveler.lines[n]
    traveler.position = n
    return true
}

// Move +/- 'n' lines in file.
func (traveler *FileTraveler) Move(n int) (moved bool) {
    return traveler.GoTo(traveler.position + n)
}

// Current returns the current line
func (traveler *FileTraveler) Current() (current []byte) {
    return traveler.current
}

// IsTop checks to see if this current line is the first line.
func (traveler *FileTraveler) IsTop() (isTop bool) {
    return traveler.position == 0
}

// Top makes the top line in the file the current line.
func (traveler *FileTraveler) Top() (moved bool) {
    return traveler.GoTo(0)
}

// GetTop makes the top line in the file the current line
// and then returns it.
func (traveler *FileTraveler) GetTop() (top []byte) {
    traveler.Top()
    return traveler.Current()
}

// IsBottom checks to see if this current line is the last line.
func (traveler *FileTraveler) IsBottom() (isBottom bool) {
    return traveler.position == traveler.bottom
}

// Bottom makes the bottom line in the file the current line.
func (traveler *FileTraveler) Bottom() (moved bool) {
    return traveler.GoTo(traveler.bottom)
}

// GetBottom makes the bottom line in the file the current line
// and then returns it.
func (traveler *FileTraveler) GetBottom() (bottom []byte) {
    traveler.Bottom()
    return traveler.Current()
}

// HasNext checks to see if there is a line after the current line.
func (traveler *FileTraveler) HasNext() (hasNext bool) {
    return traveler.position < traveler.bottom
}

// Next makes the next line in file the current line.
func (traveler *FileTraveler) Next() (next bool) {
    return traveler.Move(1)
}

// GetNext makes the next line in the file the current line
// and then returns it.
func (traveler *FileTraveler) GetNext() (next []byte) {
    if traveler.Next() {
        return traveler.Current()
    }
    return nil
}

// HasPrev checks to see if there is a line before the current line.
func (traveler *FileTraveler) HasPrev() (hasPrev bool) {
    return traveler.position-1 > 0
}

// Prev makes the previous line in file the current line.
func (traveler *FileTraveler) Prev() (prev bool) {
    return traveler.Move(-1)
}

// GetPrev makes the previous line in the file the current line
// and then returns it.
func (traveler *FileTraveler) GetPrev() (prev []byte) {
    if traveler.Prev() {
        return traveler.Current()
    }
    return nil
}

// Length returns the number of lines in the file.
func (traveler *FileTraveler) Length() (length int) {
    return len(traveler.lines)
}

// CurrentLength returns the number of bytes in the current line.
func (traveler *FileTraveler) CurrentLength() (length int) {
    return len(traveler.Current())
}

// Find takes a compiled regexp and starting at the top of the file
// returns the line number of all lines that match.
func (traveler *FileTraveler) Find(rx *regexp.Regexp) (matches []int) {
    traveler.ForEach(func(pos int, line []byte) {
        if rx.Match(line) {
            matches = append(matches, pos)
        }
    })
    return
}

// FindString takes a regexp as a string and starting at the top of the file
// returns the line number of all lines that match.
func (traveler *FileTraveler) FindString(pattern string) (matches []int) {
    rx := regexp.MustCompile(pattern)
    return traveler.Find(rx)
}

// ForEach iterates over the file from top to bottom, calling the passed method
// on each line. It leaves the current line unchanged when complete.
func (traveler *FileTraveler) ForEach(act func(position int, line []byte)) {
    defer func(t *FileTraveler, c int) {
        t.GoTo(c)
    }(traveler, traveler.Position())

    traveler.Top()
    act(traveler.Position(), traveler.Current())

    for traveler.Next() {
        act(traveler.Position(), traveler.Current())
    }

    return
}

// ForRange iterates over the file from 'start' to 'end', calling the passed method
// on each line. It leaves the current line unchanged when complete.
func (traveler *FileTraveler) ForRange(start int, end int, act func(position int, line []byte)) {
    defer func(t *FileTraveler, c int) {
        t.GoTo(c)
    }(traveler, traveler.Position())

    traveler.GoTo(start)
    act(traveler.Position(), traveler.Current())

    for traveler.Next() {
        act(traveler.Position(), traveler.Current())
        if traveler.Position() == end {
            return
        }
    }

    return
}
