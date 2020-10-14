package stream

/**
 实现类似java stream 功能
**/

type Stream struct {
	next *Stream
	tail *Stream
	f    func(args interface{}) (res interface{}, err error)
}

func InitStream() *Stream {
	s := new(Stream)
	s.tail = nil
	s.f = nil
	return s
}

func newStream(f func(args interface{}) (res interface{}, err error)) *Stream {
	s := new(Stream)
	s.f = f
	return s
}

func (s *Stream) Next(f func(args interface{}) (res interface{}, err error)) {
	if s.f == nil {
		s.f = f
		return
	}
	//other case
	newS := newStream(f)
	tail := s.tail
	if tail == nil {
		tail = s
	}
	tail.next = newS
	s.tail = newS
}

func (s *Stream) Go(args interface{}) (res interface{}, err error) {
	currentS := s
	res = args
	for {
		if currentS == nil {
			return
		}
		res, err = currentS.f(res)
		if err != nil {
			return
		}
		currentS = currentS.next
	}
	return
}
