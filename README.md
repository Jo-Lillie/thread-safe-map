# thread-safe-map
## Create thread safe maps.

In Golang there can be a lot of threads running at the same time reading maps in parallel.  Since Go's maps aren't completely thread safe, this causes data races. 

This package implements a go interface that wraps the maps and has mutexes to protect reading and writing to the map.

(_threadsafemap_test.go_)
Calls the functions and tests against dummy data.

to test this run:
```
go test -v 
```
