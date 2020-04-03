package main

import (
	"fmt"
	"log"

	"github.com/robjporter/go-library/bolts"
)

func main() {
	// Init
	buckets := []string{"ownerBucket", "sensors"}

	err := bolts.InitBolt("./database.boltdb", buckets)
	if err != nil {
		log.Fatal("Can't init boltDB")
	}

	// Put
	err = bolts.Put([]byte("ownerBucket"), []byte("ownerKey"), []byte("username"))
	fmt.Println("PUT err: ", err)

	// Get owner
	value := bolts.Get([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("GET value: ", string(value))

	// Delete
	err = bolts.Delete([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("DELETE err: ", err)

	// Insert two key/value
	err = bolts.Put([]byte("sensors"), []byte("key1"), []byte("value1"))
	fmt.Println("PUT err: ", err)
	err = bolts.Put([]byte("sensors"), []byte("key2"), []byte("value2"))
	fmt.Println("PUT err: ", err)

	// Get all keys
	keys := bolts.GetAllKeys([]byte("sensors"))
	// keys = [key1, key2]
	fmt.Println("GETALLKEYS keys: ", keys)
	for i := 0; i < len(keys); i++ {
		fmt.Println("KEYS: ", string(keys[i]))
	}

	// Get all key/value pairs
	pairs := bolts.GetAllKeyValues([]byte("sensors"))
	// pairs = [{Key:key1, Value:value1}, {Key: key2, Value:value2}]
	fmt.Println("GETALLKEYVALUES pairs: ", pairs)
	for i := 0; i < len(pairs); i++ {
		fmt.Println("PAIRS: Key - ", string(pairs[i].Key), " | value - ", string(pairs[i].Value))
	}

	// Close
	bolts.Close()
}