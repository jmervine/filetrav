
# filetrav

[![GoDoc](https://godoc.org/github.com/jmervine/filetrav?status.png)](https://godoc.org/github.com/jmervine/filetrav) [![Build Status](https://travis-ci.org/jmervine/filetrav.svg?branch=master)](https://travis-ci.org/jmervine/filetrav)

## [Documentation](https://godoc.org/github.com/jmervine/filetrav)

```go
import "github.com/jmervine/filetrav"
```
Package filetrav: Travel around files easily.

This may exist elsewhere, there's probably an easier way to do this. But it was
a fun bit of code to write and I find it useful.

Warning: This is not fully tested as of yet.

```go
    // Example:
	if traveler, err := ReadFileTraveler("_support/test.txt"); err == nil {
	    // Overly complex for example purposes.
	
	    if traveler.IsTop() {
	        fmt.Println("starting at the top")
	    }
	
	    fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	
	    if traveler.Next() {
	        fmt.Println("moving to the next line")
	        fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	    }
	
	    if traveler.Move(1) {
	        fmt.Println("moving down a line")
	        fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	    }
	
	    if traveler.Move(-1) {
	        fmt.Println("moving back up a line")
	        fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	    }
	
	    if traveler.Bottom() {
	        fmt.Println("moving to the bottom line")
	        fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	    }
	
	    if traveler.IsBottom() {
	        fmt.Println("finishing at the bottom")
	    }
	}
	for traveler.Top(); traveler.HasNext(); traveler.Next() {
	}
	
	// Output:
	//
	// starting at the top
	// 0: "foo"
	// moving to the next line
	// 1: "bar"
	// moving down a line
	// 2: "bah"
	// moving back up a line
	// 1: "bar"
	// moving to the bottom line
	// 4: ""
	// finishing at the bottom

```

### Types

#### FileTraveler
```go
type FileTraveler struct {
    // contains filtered or unexported fields
}
```


#### NewFileTraveler
```go
func NewFileTraveler(data []byte) (traveler FileTraveler, err error)
```
> NewFileTraveler creates a new FileTraveler from a []byte, like the one that
> would be read in from ioutil.ReadFile.

```go

```
#### ReadFileTraveler
```go
func ReadFileTraveler(file string) (traveler FileTraveler, err error)
```
> ReadFileTraveler creates a new FileTraveler from a path to a file on disk.

```go

```
#### Bottom
```go
func (traveler *FileTraveler) Bottom() (moved bool)
```
> Bottom makes the bottom line in the file the current line.



#### Current
```go
func (traveler *FileTraveler) Current() (current []byte)
```
> Current returns the current line



#### CurrentLength
```go
func (traveler *FileTraveler) CurrentLength() (length int)
```
> CurrentLength returns the number of bytes in the current line.



#### Find
```go
func (traveler *FileTraveler) Find(rx *regexp.Regexp) (matches []int)
```
> Find takes a compiled regexp and starting at the top of the file returns the
> line number of all lines that match.



#### FindString
```go
func (traveler *FileTraveler) FindString(pattern string) (matches []int)
```
> FindString takes a regexp as a string and starting at the top of the file
> returns the line number of all lines that match.


```go
    // Example:
	traveler, err := ReadFileTraveler("_support/test.txt")
	if err != nil {
	    panic(err)
	}
	
	for _, i := range traveler.FindString("^b.+$") {
	    if traveler.GoTo(i) {
	        fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	    }
	}
	
	// Output:
	//
	// 1: "bar"
	// 2: "bah"
	// 3: "bin"

```

#### ForEach
```go
func (traveler *FileTraveler) ForEach(act func(position int, line []byte))
```
> ForEach iterates over the file from top to bottom, calling the passed method
> on each line. It leaves the current line unchanged when complete.


```go
    // Example:
	traveler, err := ReadFileTraveler("_support/test.txt")
	if err != nil {
	    panic(err)
	}
	
	traveler.ForEach(func(pos int, line []byte) {
	    fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	})
	
	// Output:
	//
	// 0: "foo"
	// 1: "bar"
	// 2: "bah"
	// 3: "bin"
	// 4: ""

```

#### ForRange
```go
func (traveler *FileTraveler) ForRange(start int, end int, act func(position int, line []byte))
```
> ForRange iterates over the file from 'start' to 'end', calling the passed
> method on each line. It leaves the current line unchanged when complete.


```go
    // Example:
	traveler, err := ReadFileTraveler("_support/test.txt")
	if err != nil {
	    panic(err)
	}
	
	traveler.ForRange(1, 3, func(pos int, line []byte) {
	    fmt.Printf("%d: %q\n", traveler.Position(), traveler.Current())
	})
	
	// Output:
	//
	// 1: "bar"
	// 2: "bah"
	// 3: "bin"

```

#### Get
```go
func (traveler *FileTraveler) Get(n int) (current []byte)
```
> Get line 'n' in file and return it.



#### GetBottom
```go
func (traveler *FileTraveler) GetBottom() (bottom []byte)
```
> GetBottom makes the bottom line in the file the current line and then
> returns it.



#### GetNext
```go
func (traveler *FileTraveler) GetNext() (next []byte)
```
> GetNext makes the next line in the file the current line and then returns
> it.



#### GetPrev
```go
func (traveler *FileTraveler) GetPrev() (prev []byte)
```
> GetPrev makes the previous line in the file the current line and then
> returns it.



#### GetTop
```go
func (traveler *FileTraveler) GetTop() (top []byte)
```
> GetTop makes the top line in the file the current line and then returns it.



#### GoTo
```go
func (traveler *FileTraveler) GoTo(n int) (moved bool)
```
> GoTo 'n' line in file.



#### HasNext
```go
func (traveler *FileTraveler) HasNext() (hasNext bool)
```
> HasNext checks to see if there is a line after the current line.



#### HasPrev
```go
func (traveler *FileTraveler) HasPrev() (hasPrev bool)
```
> HasPrev checks to see if there is a line before the current line.



#### IsBottom
```go
func (traveler *FileTraveler) IsBottom() (isBottom bool)
```
> IsBottom checks to see if this current line is the last line.



#### IsTop
```go
func (traveler *FileTraveler) IsTop() (isTop bool)
```
> IsTop checks to see if this current line is the first line.



#### Length
```go
func (traveler *FileTraveler) Length() (length int)
```
> Length returns the number of lines in the file.



#### Move
```go
func (traveler *FileTraveler) Move(n int) (moved bool)
```
> Move +/- 'n' lines in file.


```go
    // Example:
	traveler, _ := ReadFileTraveler("_support/test.txt")
	
	if traveler.Move(1) {
	    fmt.Printf("%q\n", traveler.Current())
	}
	
	if traveler.Move(-1) {
	    fmt.Printf("%q\n", traveler.Current())
	}
	
	if traveler.Bottom() {
	    fmt.Printf("%q\n", traveler.Current())
	}
	
	if traveler.Move(-1) {
	    fmt.Printf("%q\n", traveler.Current())
	}
	
	// Output:
	// "bar"
	// "foo"
	// ""
	// "bin"

```

#### Next
```go
func (traveler *FileTraveler) Next() (next bool)
```
> Next makes the next line in file the current line.



#### Position
```go
func (traveler *FileTraveler) Position() (position int)
```
> Position within the current file.



#### Prev
```go
func (traveler *FileTraveler) Prev() (prev bool)
```
> Prev makes the previous line in file the current line.



#### Top
```go
func (traveler *FileTraveler) Top() (moved bool)
```
> Top makes the top line in the file the current line.


```go
    // Example:
	traveler, _ := ReadFileTraveler("_support/test.txt")
	
	traveler.Top()
	if traveler.IsTop() {
	    fmt.Println("top")
	}
	
	// Output: top

```


