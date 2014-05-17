package filetrav

import (
    "fmt"
    "github.com/jmervine/GoT"
    "testing"
)

func stub(start int) FileTraveler {
    l := [][]byte{
        []byte("foo"),
        []byte("bar"),
        []byte("bah"),
        []byte("bin"),
    }

    return FileTraveler{
        lines:    l,
        current:  l[start],
        position: start,
        bottom:   3,
    }
}

func TestPosition(T *testing.T) {
    t := GoT.Go(T)
    traveler := stub(1)

    t.AssertEqual(traveler.Position(), 1)
}

func TestCurrent(T *testing.T) {
    t := GoT.Go(T)
    traveler := stub(1)

    t.AssertDeepEqual(traveler.Current(), []byte("bar"))
}

func TestGoTo(T *testing.T) {
    t := GoT.Go(T)
    traveler := stub(1)

    t.Assert(traveler.GoTo(0))
    t.AssertEqual(traveler.Position(), 0)

    t.Assert(traveler.GoTo(3))
    t.AssertEqual(traveler.Position(), 3)

    t.Refute(traveler.GoTo(-1))
    t.AssertEqual(traveler.Position(), 3)
}

func TestMove(T *testing.T) {
}

func TestIsTop(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.Assert(traveler.IsTop())

    traveler = stub(2)
    t.Refute(traveler.IsTop())
}

func TestIsBottom(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(3)
    t.Assert(traveler.IsBottom())

    traveler = stub(2)
    t.Refute(traveler.IsBottom())
}

func TestTop(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.Refute(traveler.Top())
    t.Assert(traveler.IsTop())

    traveler = stub(3)
    t.Assert(traveler.Top())
    t.Assert(traveler.IsTop())
}

func TestGetTop(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.AssertDeepEqual(traveler.GetTop(), []byte("foo"))

    traveler = stub(3)
    t.AssertDeepEqual(traveler.GetTop(), []byte("foo"))
}

func TestBottom(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.Assert(traveler.Bottom())
    t.Assert(traveler.IsBottom())

    traveler = stub(3)
    t.Refute(traveler.Bottom())
    t.Assert(traveler.IsBottom())
}

func TestGetBottom(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.AssertDeepEqual(traveler.GetBottom(), []byte("bin"))

    traveler = stub(3)
    t.AssertDeepEqual(traveler.GetBottom(), []byte("bin"))
}

func TestHasNext(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.Assert(traveler.HasNext())

    traveler = stub(3)
    t.Refute(traveler.HasNext())
}

func TestNext(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.Assert(traveler.Next())
    t.AssertDeepEqual(traveler.Current(), []byte("bar"))

    traveler = stub(3)
    t.Refute(traveler.Next())
    t.AssertDeepEqual(traveler.Current(), []byte("bin"))
}

func TestGetNext(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(0)
    t.AssertDeepEqual(traveler.GetNext(), []byte("bar"))

    traveler = stub(3)
    t.AssertNil(traveler.GetNext())
}

func TestHasPrev(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(3)
    t.Assert(traveler.HasPrev())

    traveler = stub(0)
    t.Refute(traveler.HasPrev())
}

func TestPrev(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(3)
    t.Assert(traveler.Prev())
    t.AssertDeepEqual(traveler.Current(), []byte("bah"))

    traveler = stub(0)
    t.Refute(traveler.Prev())
    t.AssertDeepEqual(traveler.Current(), []byte("foo"))
}

func TestGetPrev(T *testing.T) {
    t := GoT.Go(T)

    traveler := stub(3)
    t.AssertDeepEqual(traveler.GetPrev(), []byte("bah"))

    traveler = stub(0)
    t.AssertNil(traveler.GetPrev())
}

func Example() {
    traveler, err := ReadFileTraveler("_support/test.txt")
    if err != nil {
        panic(err)
    }

    for ; traveler.HasNext(); traveler.Next() {
        fmt.Printf("%d: %s\n", traveler.Position(), traveler.Current())
    }

    // Output:
    //
    // 0: foo
    // 1: bar
    // 2: bah
    // 3: bin
}
