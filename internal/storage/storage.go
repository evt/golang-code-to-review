package storage

// Store is the memory key value store
type Store struct {
	values map[string]string
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Load(key string) string {
	return s.values[key]
}

func (s *Store) Remove(key string) {
	s.values[key] = ""
}

func (s *Store) Exist(key string) bool {
	if s.values[key] == "" {
		return true
	}
	return false
}


