package main

import (
	"fmt"
	"log"

	"github.com/robjporter/go-library/xbolts"
)

func main() {
	// Init
	buckets := []string{"ownerBucket", "sensors"}

	err := xbolts.InitBolt("./database.boltdb", buckets)
	if err != nil {
		log.Fatal("Can't init boltDB")
	}

	// Put
	err = xbolts.Put([]byte("ownerBucket"), []byte("ownerKey"), []byte("username"))
	fmt.Println("PUT err: ", err)

	// Get owner
	value := xbolts.Get([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("GET value: ", string(value))

	// Delete
	err = xbolts.Delete([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("DELETE err: ", err)

	// Insert two key/value
	err = xbolts.Put([]byte("sensors"), []byte("key1"), []byte("value1"))
	fmt.Println("PUT err: ", err)
	err = xbolts.Put([]byte("sensors"), []byte("key2"), []byte("value2"))
	fmt.Println("PUT err: ", err)

	// Get all keys
	keys := xbolts.GetAllKeys([]byte("sensors"))
	// keys = [key1, key2]
	fmt.Println("GETALLKEYS keys: ", keys)
	for i := 0; i < len(keys); i++ {
		fmt.Println("KEYS: ", string(keys[i]))
	}

	// Get all key/value pairs
	pairs := xbolts.GetAllKeyValues([]byte("sensors"))
	// pairs = [{Key:key1, Value:value1}, {Key: key2, Value:value2}]
	fmt.Println("GETALLKEYVALUES pairs: ", pairs)
	for i := 0; i < len(pairs); i++ {
		fmt.Println("PAIRS: Key - ", string(pairs[i].Key), " | value - ", string(pairs[i].Value))
	}

	// Close
	xbolts.Close()
}