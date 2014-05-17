# filetrav

[![GoDoc](https://godoc.org/github.com/jmervine/filetrav?status.png)](https://godoc.org/github.com/jmervine/filetrav) [![Build Status](https://travis-ci.org/jmervine/filetrav.svg?branch=master)](https://travis-ci.org/jmervine/filetrav)

## [Documentation](https://godoc.org/github.com/jmervine/filetrav)

```go
import "github.com/jmervine/filetrav"
```

Package filetrav: Travel around files easily.

This may exist elsewhere, there's probably an easier way to do this. But it was
a fun bit of code to write and I find it useful.


### Types

#### FileTraveler

```go
type FileTraveler struct {
    // contains filtered or unexported fields
}
```




### Functions
#### NewFileTraveler

```go
func NewFileTraveler(data []byte) (traveler FileTraveler, err error)
```
NewFileTraveler creates a new FileTraveler from a []byte, like the one that
would be read in from ioutil.ReadFile.


#### ReadFileTraveler

```go
func ReadFileTraveler(file string) (traveler FileTraveler, err error)
```
ReadFileTraveler creates a new FileTraveler from a path to a file on disk.


#### Bottom

```go
func (traveler *FileTraveler) Bottom() (moved bool)
```
Bottom makes the bottom line in the file the current line.



#### Current

```go
func (traveler *FileTraveler) Current() (current []byte)
```
Current returns the current line



#### GetBottom

```go
func (traveler *FileTraveler) GetBottom() (bottom []byte)
```
GetBottom makes the bottom line in the file the current line and then returns
it.



#### GetNext

```go
func (traveler *FileTraveler) GetNext() (next []byte)
```
GetNext makes the next line in the file the current line and then returns it.



#### GetPrev

```go
func (traveler *FileTraveler) GetPrev() (prev []byte)
```
GetPrev makes the previous line in the file the current line and then returns
it.



#### GetTop

```go
func (traveler *FileTraveler) GetTop() (top []byte)
```
GetTop makes the top line in the file the current line and then returns it.



#### GoTo

```go
func (traveler *FileTraveler) GoTo(n int) (moved bool)
```
GoTo 'n' line in file.



#### HasNext

```go
func (traveler *FileTraveler) HasNext() (hasNext bool)
```



#### HasPrev

```go
func (traveler *FileTraveler) HasPrev() (hasPrev bool)
```



#### IsBottom

```go
func (traveler *FileTraveler) IsBottom() (isBottom bool)
```



#### IsTop

```go
func (traveler *FileTraveler) IsTop() (isTop bool)
```



#### Move

```go
func (traveler *FileTraveler) Move(n int) (moved bool)
```
Move +/- 'n' lines in file.



#### Next

```go
func (traveler *FileTraveler) Next() (next bool)
```
Next makes the next line in file the current line.



#### Position

```go
func (traveler *FileTraveler) Position() (position int)
```
Position within the current file.



#### Prev

```go
func (traveler *FileTraveler) Prev() (prev bool)
```
Prev makes the previous line in file the current line.



#### Top

```go
func (traveler *FileTraveler) Top() (moved bool)
```
Top makes the top line in the file the current line.




