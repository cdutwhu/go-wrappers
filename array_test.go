package wrappers

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {

	rst := IArrMake("Strs", 3, "hello")
	rst111 := []string(rst.(Strs))
	fPln(rst111)

	// arr := I32s{1, 2, 3, 4, 5}

	arr1 := Strs{"a", "b", "c", "c", "a", "D", "b", "A"}
	rst11 := IArrFoldRep(arr1, "[i]").([]string)
	fPln(rst11)

	arr := F64s{1, 2, 3, 4, 5.5, 1, 6, 7, 4}
	fPln(arr)

	arr.Set(4, 5.777)
	fPln(arr)

	// fPln(IArrEleIn(float64(3), arr)) // removed, use u.XIn

	ok1, index, rst1 := IArrSearchOne(arr, func(i int, a interface{}) (bool, interface{}) { return a == float64(4), "junk" })
	if ok1 {
		fPln(ok1, index, rst1.(float64))
	} else {
		fPln("no")
	}

	ok, indices, rst := IArrSearch(arr, func(i int, a interface{}) (bool, interface{}) { return a == float64(1), "junk" })
	if ok {
		fPln(ok, indices, rst.([]float64))
	} else {
		fPln("no")
	}
}

func TestInsert(t *testing.T) {
	// arr := I32s{1, 2, 3, 4, 5}
	arr := Strs{"a", "b", "c", "d"}
	fPln(IArrInsert(arr, INSERT_BEFORE, func(i int, a interface{}) (bool, interface{}) { return i == 0 || i == 1 || a == "d", a.(string) + "1" }))
	fPln(IArrInsert(arr, INSERT_AFTER, func(i int, a interface{}) (bool, interface{}) { return i == 0 || i == 1 || a == "d", a.(string) + "2" }))
}

func TestRemove(t *testing.T) {
	arr := I32s{1, 2, 3, 4, 5}
	// arr := Strs{"a", "b", "c", "d"}
	fPln(IArrRemove(arr, func(i int, a interface{}) (bool, interface{}) { return a == 4 || i == 0, "Junk" }))
}

func TestReplace(t *testing.T) {
	arr := I32s{1, 2, 3, 4, 5}
	// arr := Strs{"a", "b", "c", "d"}
	fPln(IArrReplace(arr, func(i int, a interface{}) (bool, interface{}) { return a == 4 || i == 0, a.(int) + 100 }))
}

func TestRmRep(t *testing.T) {
	arr := Strs{"abc", "::", "abc", "de", "de"}
	// arr := I32s{ 2, 4, 5, 6, 3, 2, 5, 2, 4 }
	fPln(IArrRmRep(arr))
}

func TestContain(t *testing.T) {
	arr := Strs{"::", "abc", "de", "mn", " "}
	fPln(IArrCtns(arr, "abc", "mn", "de", "::", " "))
}

func TestSeqContain(t *testing.T) {
	arr := Strs{"::", "abc", "de", "mn", ""}
	fPln(IArrSeqCtns(arr, "::", "de", "mn"))
}

func TestOrder(t *testing.T) {
	arr := Strs{":::", "abcd", "zt", "de", "mn", "A", ""}
	FunSortLess = func(left, right interface{}) bool { return len(left.(string)) < len(right.(string)) }
	sort.Sort(arr)
	fPln(arr)
	arr1 := I32s{2, 4, 5, 6, 3, 2, 5, 2, 4}
	FunSortLess = func(left, right interface{}) bool { return left.(int) < right.(int) }
	sort.Sort(arr1)
	fPln(arr1)
}

func TestSameEle(t *testing.T) {
	arr := Strs{"abc", "abc", "abc"}
	fPln(IArrIsSameEle(arr))
}

func TestInterSecANDUnion(t *testing.T) {
	arr1 := Strs{"::", "abcd", "def", "mn", "A", "C"}
	arr2 := Strs{"abcd", "def", "::", "A", "B", "D"}
	r := IArrIntersect(arr1, arr2)
	fPln(r.([]string))
	r = IArrUnion(arr1, arr2)
	fPln(r.([]string))
}

func TestIArrStrJoinEx(t *testing.T) {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"1", "2", "3"}
	fPln(IArrStrJoinEx(Strs(s1), Strs(s2), "#", " ~ "))
}
