package store

// Store holds the latest check results
type Store struct {
	results map[string]interface{}
}

func New() *Store {
	return &Store{
		results: make(map[string]interface{}),
	}
}

// UpdateResult saves the latest result for a URL
// INTENTIONAL GAP: This map access is not thread-safe!
// Concurrent writes from the checker and reads from the server will cause a panic eventually.
func (s *Store) UpdateResult(url string, result interface{}) {
	s.results[url] = result
}

func (s *Store) GetAll() map[string]interface{} {
	return s.results
}
