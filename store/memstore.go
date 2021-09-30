package store

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
)

type MemStore struct {
	sync.Mutex
}

//Save to file save all contents contained in the cache into a file so it can be read from it later
func (m MemStore) SaveToFile(path string) error {
	db, err := bolt.Open(path, os.ModePerm, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	start := time.Now()
	m.Mutex.Lock()
	bulkInputs := make([]*Inputs, len(input))
	i := 0
	for _, v := range Inputs {
		bulkInputs[i] = v
		i++
	}
	m.Mutex.Unlock()
	if Debug {
		log.Printf("Unlocked after %s", time.Since(start))
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_ = tx.DeleteBucket([]byte("inputs"))
		bucket, err := tx.CreateBucket([]byte("inputs"))
		if err != nil {
			return err
		}
		for _, bulkInput := range bulkInputs {
			buffer := bytes.Buffer{}
			err := gob.NewEncoder(&buffer).Encode(bulkInput)
			if err != nil {
				continue
			}
			bucket.Put([]byte(bulkInput.Key), buffer.Bytes())
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

//The read function is used to read contentes stored in the cache file
func (MemStore) ReadFromFile() {

}

func (MemStore) DeleteFromFile() {

}

func (MemStore) DeleteAllFromFile() {

}
