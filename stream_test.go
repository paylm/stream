package stream

import (
	"testing"
)

func Test_Go(t *testing.T) {

	f1 := func(name interface{}) (res interface{}, err error) {
		t.Logf("call f1()->%v\n", name)
		return "f1", nil
	}
	f2 := func(name interface{}) (res interface{}, err error) {
		t.Logf("call f2()->%v\n", name)
		return "f2", nil
	}
	f3 := func(name interface{}) (res interface{}, err error) {
		t.Logf("call f3()->%v\n", name)
		return "f3", nil
	}

	s := InitStream()
	s.Next(f1)
	s.Next(f2)
	s.Next(f3)
	s.Next(f1)
	s.Go("zhang")

	//fmt.Printf("s=>%#v\n", s)
	var trvS *Stream = s
	for {
		if trvS == nil {
			break
		}
		t.Logf("f=>%p", trvS.f)
		trvS = trvS.next
	}
}
