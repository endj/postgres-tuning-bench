package state

import "sync"

type PostRequest struct {
	post    string
	replyTo uint64
}

type ThreadRequest struct {
	title string
	post  string
	id    uint64
}

type Thread struct {
	title string
	id    uint64
	posts []Post
}

type Post struct {
	post    string
	id      uint64
	replyTo uint64
}

type MutexMap struct {
	boardMap map[uint64]Thread
	mux      sync.Mutex
}

func GetThreadById(id uint64, s *MutexMap) *Thread {
	s.mux.Lock()
	val, exists := s.boardMap[id]
	s.mux.Unlock()
	if exists {
		return &val
	} else {
		return nil
	}
}

func CreateThread(request *ThreadRequest, s *MutexMap) {
	s.mux.Lock()
}

func CreateMap() MutexMap {
	return MutexMap{boardMap: make(map[uint64]Thread)}
}
