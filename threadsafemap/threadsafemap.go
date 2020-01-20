package threadsafemap

import (
	"sync"
	"errors"
)

/* write a Go interface that wraps the maps and has mutexes that protects
reading & writing to the map */
type ThreadSafeWrapper interface {
	/* specify set of method signatures */
	// empty interface - can hold any data and work with any type
	Read(string) interface{}
	Write(string, interface{})
	Exist(string) interface{}
	Delete(string)
}

/* the primary job of an interface is to provide only method signatures
consisting of the method name, input arguments and return types
it is up to a type (e.g. struct type) to declare methods and implement them */

/* create a struct with a map & required mutexes */
type ThreadSafeMap struct {
	mutex      *sync.Mutex // <- this mutex protects the map below
	/* map[string]interface{} is a  map whose keys are strings and
	values are any type */
	threadsafe map[string]interface{}
}


var ErrNotFound = errors.New("key not found")

/* New() will instantiate an instance of the threadsafe map struct that we can use */
// create an instance of threadsafemap and return it
// New instance of threadsafemap
func New (inputThreadsafemap map[string]interface{}) ThreadSafeMap {
	newMutex := sync.Mutex{}
	// threadsafe := make(map[string]interface{}) /* threadsafe declared and not used */
	return ThreadSafeMap{mutex: &newMutex, threadsafe: inputThreadsafemap}
}

/* add a thread safe read function */
// Read() is reading the value
func (r *ThreadSafeMap) Read(key string) (interface{}, error) { 
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.threadsafe[key]; !ok {
		return "", ErrNotFound 
	}	
	return r.threadsafe[key], nil
}

/* add a thread safe write function */
// Write() is writing a new key/value pair to the map
func (w *ThreadSafeMap) Write(key string, value interface{}) interface{} {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.threadsafe[key] = value
	return ""
}

/* add exists function - external could ask do you have this key */
func (e *ThreadSafeMap) Exists(key string) (interface{}, bool) { // return just bool?
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if _, found := e.threadsafe[key]; !found {
		return ErrNotFound, false
	}
	return e.threadsafe[key], true
}

/* add a delete function to delete a key-value pair */
func (d *ThreadSafeMap) Delete(key string) (error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _, found := d.threadsafe[key]; !found {
		return ErrNotFound
	}
	delete(d.threadsafe, key)
	return nil
}
