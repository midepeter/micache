package store

type MemStore struct {
	map[string]interface{}
}

func (m MemStore) Put(data interface{}) {
	m.["first"]  = []string{"sample of files"}
}
