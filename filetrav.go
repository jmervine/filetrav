/*
Package filetrav: Travel around files easily.

This may exist elsewhere, there's probably an easier way to do this. But it was a fun bit of code to write and I find it useful.

Warning: This is not fully tested as of yet.
*/
package filetrav

import (
    "bytes"
    "io/ioutil"
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
    traveler.bottom = (len(traveler.lines) - 1)
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

    traveler.position = n
    traveler.current = traveler.lines[n]
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

func (traveler *FileTraveler) HasNext() (hasNext bool) {
    return traveler.position+1 < len(traveler.lines)
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
