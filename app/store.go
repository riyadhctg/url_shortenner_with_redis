package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

// struct is collection of fields whereas interface type is a set of method signature
// Go doesn't have class but you can define methods on types
/**
method vs function: methods are functions with a special receiver argument. For example:
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
ref: https://go.dev/tour
*/

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	//file *os.File
	save chan record

}

type record struct {
	key, URL string
}


const saveQueueLength = 1000

func NewURLStore(filename string) *URLStore {
	// maps are reference type
	// new just allocates memory, not initializes memory; make allocates and initializes memory
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	// f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal("URLStore:", err)
	// }
	// s.file = f
	if err := s.load(filename); err != nil {
		log.Println("Error while loading the URLStore:", err)
	}
	go s.saveloop(filename)
	return s
}

/**
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
ref: https://go.dev/tour
*/
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// if _, ok pattern
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}


// func (s *URLStore) save(key, url string) error {
// 	e := gob.NewEncoder(s.file)
// 	return e.Encode(record{key, url})
// }

func (s *URLStore) saveloop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		// taking a record from the channe
		r := <- s.save 
		// then encoding it
		if err := e.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}

}


func (s *URLStore) Put(url string) string {
	// for {} is infinite loop so there must be something inside the loop to break out of the loop
	// this for loop keeps trying until a unique key is generated by genbKey
	// for existing keys, our Set method will return false, so the loop will continue
	for {
		key := genKey()
		if ok := s.Set(key, url); ok {
			s.save <- record{key, url}
			// if err := s.save(key, url); err != nil {
			// 	log.Println("Error putting data in URLStore:", err)
			// }
			return key
		}

	}
	panic("shouldn't get here")
}


func (s *URLStore) load(filename string) error {
	// if _, err := s.file.Seek(0,0); err != nil {
	// 	return err
	// }
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
