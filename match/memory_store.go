package match

import (
	"github.com/jmartin82/mmock/definition"
)

//MemoryStore stores all received request and their matches in memory until the last reset
type MemoryStore struct {
	matches []definition.Match
}

//Save store a match informtion
func (mrs *MemoryStore) Save(req definition.Match) {
	mrs.matches = append(mrs.matches, req)
}

//Reset clean the request stored in memory
func (mrs *MemoryStore) Reset() {
	mrs.matches = make([]definition.Match, 0, 100)
}

//GetAll return current matches (positive and negative) in memory
func (mrs *MemoryStore) GetAll() []definition.Match {
	r := make([]definition.Match, len(mrs.matches))
	copy(r, mrs.matches)
	return r
}

//NewMemoryStore is the MemoryStore contructor
func NewMemoryStore() *MemoryStore {
	reqs := make([]definition.Match, 0, 100)
	return &MemoryStore{matches: reqs}

}
