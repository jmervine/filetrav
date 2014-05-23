package filetrav

import (
    "fmt"
    . "github.com/jmervine/GoT"
    "regexp"
    "testing"
)

var traveler, _ = ReadFileTraveler("_support/test.txt")

func reload() {
    traveler.GoTo(0)
}

func stub(start int) FileTraveler {
    l := [][]byte{
        []byte("foo"),
        []byte("bar"),
        []byte("bah"),
        []byte("bin"),
        []byte(""),
    }

    return FileTraveler{
        lines:    l,
        current:  l[start],
        position: start,
        bottom:   3,
    }
}

func TestPosition(T *testing.T) {
    Go(T).AssertEqual(traveler.Position(), 0)
}

func TestCurrent(T *testing.T) {
    Go(T).AssertDeepEqual(traveler.Current(), []byte("foo"))
}

func TestGoTo(T *testing.T) {
    Go(T).Assert(traveler.GoTo(1))
    Go(T).AssertEqual(traveler.Position(), 1)

    Go(T).Assert(traveler.GoTo(0))
    Go(T).AssertEqual(traveler.Position(), 0)

    Go(T).Assert(traveler.GoTo(3))
    Go(T).AssertEqual(traveler.Position(), 3)

    Go(T).Refute(traveler.GoTo(-1))
    Go(T).AssertEqual(traveler.Position(), 3)
}

func TestMove(T *testing.T) {
    reload()

    traveler.Move(1)
    Go(T).AssertEqual(traveler.Position(), 1)

    traveler.Move(-1)
    Go(T).AssertEqual(traveler.Position(), 0)

    traveler.Move(-1)
    Go(T).AssertEqual(traveler.Position(), 0)
}

func TestIsTop(T *testing.T) {
    reload()
    Go(T).Assert(traveler.IsTop())

    traveler.Move(1)
    Go(T).Refute(traveler.IsTop())
}

func TestIsBottom(T *testing.T) {
    reload()
    Go(T).Refute(traveler.IsBottom())

    traveler.GoTo(traveler.bottom)
    Go(T).Assert(traveler.IsBottom())
}

func TestTop(T *testing.T) {
    reload()

    Go(T).Refute(traveler.Top())
    Go(T).Assert(traveler.IsTop())

    traveler.GoTo(traveler.bottom)
    Go(T).Assert(traveler.Top())
    Go(T).Assert(traveler.IsTop())
}

func TestGet(T *testing.T) {
    reload()
    Go(T).AssertDeepEqual(traveler.Get(1), []byte("bar"))
    Go(T).AssertDeepEqual(traveler.Get(3), []byte("bin"))
}

func TestGetTop(T *testing.T) {
    reload()
    Go(T).AssertDeepEqual(traveler.GetTop(), []byte("foo"))

    traveler.GoTo(traveler.bottom)
    Go(T).AssertDeepEqual(traveler.GetTop(), []byte("foo"))
}

func TestBottom(T *testing.T) {
    reload()

    Go(T).Assert(traveler.Bottom())
    Go(T).Assert(traveler.IsBottom())

    traveler.GoTo(traveler.bottom)
    Go(T).Refute(traveler.Bottom())
    Go(T).Assert(traveler.IsBottom())
}

func TestGetBottom(T *testing.T) {
    reload()

    Go(T).AssertDeepEqual(traveler.GetBottom(), []byte(""))

    traveler.GoTo(traveler.bottom)
    Go(T).AssertDeepEqual(traveler.GetBottom(), []byte(""))
}

func TestHasNext(T *testing.T) {
    reload()

    Go(T).Assert(traveler.HasNext())

    traveler.GoTo(traveler.bottom)
    Go(T).Refute(traveler.HasNext())

    for traveler.Top(); traveler.HasNext(); traveler.Next() {
    }

    Go(T).AssertEqual(traveler.Position(), 4)
}

func TestNext(T *testing.T) {
    reload()

    Go(T).Assert(traveler.Next())
    Go(T).AssertDeepEqual(traveler.Current(), []byte("bar"))

    traveler.GoTo(traveler.bottom)
    Go(T).Refute(traveler.Next())
    Go(T).AssertDeepEqual(traveler.Current(), []byte(""))
}

func TestGetNext(T *testing.T) {
    reload()

    Go(T).AssertDeepEqual(traveler.GetNext(), []byte("bar"))

    traveler.GoTo(traveler.bottom)
    Go(T).AssertNil(traveler.GetNext())
}

func TestHasPrev(T *testing.T) {
    reload()

    Go(T).Refute(traveler.HasPrev())

    traveler.GoTo(traveler.bottom)
    Go(T).Assert(traveler.HasPrev())
}

func TestPrev(T *testing.T) {
    reload()

    Go(T).Refute(traveler.Prev())
    Go(T).AssertDeepEqual(traveler.Current(), []byte("foo"))

    traveler.GoTo(traveler.bottom)
    Go(T).Assert(traveler.Prev())
    Go(T).AssertDeepEqual(traveler.Current(), []byte("bin"))
}

func TestGetPrev(T *testing.T) {
    reload()

    Go(T).AssertNil(traveler.GetPrev())

    traveler.GoTo(traveler.bottom)
    Go(T).AssertDeepEqual(traveler.GetPrev(), []byte("bin"))
}

func TestLength(T *testing.T) {
    Go(T).AssertEqual(traveler.Length(), 5)
}

func TestCurrentLength(T *testing.T) {
    reload()
    Go(T).AssertEqual(traveler.CurrentLength(), 3)
}

func TestFind(T *testing.T) {
    traveler.GoTo(0)
    rx := regexp.MustCompile("^foo$")
    Go(T).AssertLength(traveler.Find(rx), 1)
    Go(T).AssertDeepEqual(traveler.Find(rx), []int{0})

    traveler.GoTo(traveler.bottom)
    rx = regexp.MustCompile("^b.+$")
    Go(T).AssertLength(traveler.Find(rx), 3)
    Go(T).AssertDeepEqual(traveler.Find(rx), []int{1, 2, 3})
}

func TestFindString(T *testing.T) {
    traveler.GoTo(0)
    Go(T).AssertLength(traveler.FindString("^foo$"), 1)
    Go(T).AssertDeepEqual(traveler.FindString("^foo$"), []int{0})

    traveler.GoTo(traveler.bottom)
    Go(T).AssertLength(traveler.FindString("^b.+$"), 3)
    Go(T).AssertDeepEqual(traveler.FindString("^b.+$"), []int{1, 2, 3})
}

func TestForEach(T *testing.T) {
    reload()

    var ls [][]byte
    traveler.ForEach(func(p int, l []byte) {
        ls = append(ls, l)
    })

    Go(T).AssertLength(ls, 5)
    Go(T).AssertDeepEqual(ls[len(ls)-1], []byte(""), "Last recorded should be last line of test file.")

    Go(T).AssertEqual(traveler.Position(), 0, "Position should not change.")
}

func TestForRange(T *testing.T) {
    reload()

    var ls [][]byte
    traveler.ForRange(1, 3, func(p int, l []byte) {
        ls = append(ls, l)
    })

    Go(T).AssertLength(ls, 3)
    Go(T).AssertDeepEqual(ls[2], []byte("bin"))
    Go(T).AssertEqual(traveler.Position(), 0)
}

func Example() {
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
}

func ExampleFileTraveler_Top() {
    traveler, _ := ReadFileTraveler("_support/test.txt")

    traveler.Top()
    if traveler.IsTop() {
        fmt.Println("top")
    }

    // Output: top
}

func ExampleFileTraveler_Move() {
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
}

func ExampleFileTraveler_FindString() {
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
}

func ExampleFileTraveler_ForEach() {
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
}

func ExampleFileTraveler_ForRange() {
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
}
