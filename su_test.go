package su

import (
	"testing"
)

func TestAppend(t *testing.T) {
	arr := New[int]()
	arr.Append(0)
	arr.Append(1)
	arr.Append(2, 3, 4)
	arr.Append([]int{5, 6}...)
	for i := 0; i < arr.Len(); i++ {
		if i != arr.Get(i) {
			t.Error("invalid element value")
		}
	}
}

func TestPrepend(t *testing.T) {
	arr := New([]int{1, 2, 3})
	arr.Prepend(0)
	for i := 0; i < arr.Len(); i++ {
		if i != arr.Get(i) {
			t.Error("invalid element value")
		}
	}
}

func TestRemove(t *testing.T) {
	arr := New([]int{1, 2})
	removedValue := arr.Remove(0)
	if removedValue != 1 {
		t.Error("failed remove")
	}
	if arr.Get(0) != 2 {
		t.Error("invalid slice")
	}
}

func TestRemoveRange(t *testing.T) {
	arr := New([]int{1, 2, 3, 4, 5, 6, 7, 8})
	arr.RemoveRange(0, 1)
	if arr.Get(0) != 3 {
		t.Error("failed remove range")
	}
	arr.RemoveRange(4, 5)
	for i := 0; i < arr.Len(); i++ {
		switch i {
		case 0:
			if arr.Get(i) != 3 {
				t.Error("invalid slice")
			}
		case 1:
			if arr.Get(i) != 4 {
				t.Error("invalid slice")
			}
		case 2:
			if arr.Get(i) != 5 {
				t.Error("invalid slice")
			}
		case 3:
			if arr.Get(i) != 6 {
				t.Error("invalid slice")
			}
		}
	}
}

func TestLen(t *testing.T) {
	arr := New([]int{1, 2, 3})
	if arr.Len() != 3 {
		t.Error("invalid length")
	}
}

func TestIsEmpty(t *testing.T) {
	arr := New[int]()
	if arr.IsEmpty() == false {
		t.Error("invalid IsEmpty")
	}
}

func TestPtr(t *testing.T) {
	arr := New([]int{1, 2, 3})
	ptr := arr.Ptr(2)
	if *ptr != 3 {
		t.Error("invalid ptr")
	}
}

func TestGet(t *testing.T) {
	arr := New([]int{1, 2, 3})
	value := arr.Get(2)
	if value != 3 {
		t.Error("invalid get")
	}
}

func TestSet(t *testing.T) {
	arr := New([]int{1, 2, 3})
	arr.Set(1, 10)
	value := arr.Get(1)
	if value != 10 {
		t.Error("invalid set")
	}
}

func TestGetSlice(t *testing.T) {
	arr := New([]int{1, 2, 3})
	s := arr.GetSlice()
	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Error("invalid GetSlice")
	}
}

func TestJoin(t *testing.T) {
	arr := New([]int{1, 2, 3})
	join1 := arr.Join()
	if join1 != "1,2,3" {
		t.Error("invalid join")
	}
	join2 := arr.Join(", ")
	if join2 != "1, 2, 3" {
		t.Error("invalid join")
	}
}

func TestReverse(t *testing.T) {
	arr := New([]int{1, 2, 3})
	arr.Reverse()
	if arr.Get(0) != 3 || arr.Get(1) != 2 || arr.Get(2) != 1 {
		t.Error("invalid reverse")
	}
}

func TestNew(t *testing.T) {
	arr := New([]int{1, 2}, []int{4, 5})
	if arr.Get(0) != 1 ||
		arr.Get(1) != 2 ||
		arr.Get(2) != 4 ||
		arr.Get(3) != 5 {
		t.Error("invalid New")
	}
}

func TestRemoveAll(t *testing.T) {
	arr := New([]int{1, 2, 3})
	arr.RemoveAll()
	if arr.Len() != 0 {
		t.Error("RemoveAll")
	}
}

func TestString(t *testing.T) {
	arr := New([]int{1, 2, 3})
	str := arr.String()
	if str != "[1, 2, 3]" {
		t.Error("invalid String")
	}
}

func TestSort(t *testing.T) {
	arr := New([]int{3, 2, 1})

	arr.Sort(func(i, j int) bool {
		return arr.Get(i) < arr.Get(j)
	})
}

func TestTrueForAll(t *testing.T) {
	arr := New([]int{3, 2, 1})
	if arr.TrueForAll(func(value int) bool {
		return value < 5
	}) != true {
		t.Error("error TrueForAll")
	}

	if arr.TrueForAll(func(value int) bool {
		return value < 2
	}) != false {
		t.Error("error TrueForAll")
	}
}

func TestFileter(t *testing.T) {
	arr := New([]int{4, 3, 2, 1})
	filtered := arr.Filter(func(value int) bool {
		return value%2 == 0
	})

	if filtered.Get(0) != 4 || filtered.Get(1) != 2 || filtered.Len() != 2 {
		t.Error("error Filter")
	}
}

func TestIndexOf(t *testing.T) {
	arr := New([]int{1, 2, 3, 4})
	if arr.IndexOf(func(value int) bool {
		return value == 4
	}) != 3 {
		t.Error("error IndexOf")
	}
	if arr.IndexOf(func(value int) bool {
		return value == 5
	}) != -1 {
		t.Error("error IndexOf")
	}
}

func TestSome(t *testing.T) {
	arr := New([]int{1, 3, 4})
	if arr.Some(func(value int) bool {
		return value > 3
	}) != true {
		t.Error("error Some")
	}

	if arr.Some(func(value int) bool {
		return value > 4
	}) != false {
		t.Error("error Some")
	}
}

func TestMap(t *testing.T) {
	arr := New([]int{1, 2, 3})
	arr.Map(func(value int) int {
		return value * 2
	})
	if arr.Get(0) != 2 ||
		arr.Get(1) != 4 ||
		arr.Get(2) != 6 {
		t.Error("error Map")
	}
}

func TestIterator(t *testing.T) {
	arr := New([]int{1, 2, 3})
	itr := arr.Iterator()
	for itr.MoveNext() {
		i, v := itr.Current()
		switch i {
		case 0:
			if v != 1 {
				t.Error("error 1")
			}
		case 1:
			if v != 2 {
				t.Error("error 2")
			}
		case 2:
			if v != 3 {
				t.Error("error 3")
			}
		}
	}
	if i, _ := itr.Current(); i != -1 {
		t.Error("error itr len check")
	}
	itr.Reset()
	for itr.MoveNext() {
		itr.Remove()
	}
	if arr.Len() != 0 {
		t.Error("error itr Remove")
	}
}
