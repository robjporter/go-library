package main

import (
	"fmt"
	"log"
)

func main() {
	// Init
	buckets := []string{"ownerBucket", "sensors"}

	err := boltdbboilerplate.InitBolt("./database.boltdb", buckets)
	if err != nil {
		log.Fatal("Can't init boltDB")
	}

	// Put
	err = boltdbboilerplate.Put([]byte("ownerBucket"), []byte("ownerKey"), []byte("username"))
	fmt.Println("PUT err: ", err)

	// Get owner
	value := boltdbboilerplate.Get([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("GET value: ", value)

	// Delete
	err = boltdbboilerplate.Delete([]byte("ownerBucket"), []byte("ownerKey"))
	fmt.Println("DELETE err: ", err)

	// Insert two key/value
	err = boltdbboilerplate.Put([]byte("sensors"), []byte("key1"), []byte("value1"))
	fmt.Println("PUT err: ", err)
	err = boltdbboilerplate.Put([]byte("sensors"), []byte("key2"), []byte("value2"))
	fmt.Println("PUT err: ", err)

	// Get all keys
	keys := boltdbboilerplate.GetAllKeys([]byte("sensors"))
	// keys = [key1, key2]
	fmt.Println("GETALLKEYS keys: ", keys)

	// Get all key/value pairs
	pairs := boltdbboilerplate.GetAllKeyValues([]byte("sensors"))
	// pairs = [{Key:key1, Value:value1}, {Key: key2, Value:value2}]
	fmt.Println("GETALLKEYVALUES pairs: ", pairs)

	// Close
	boltdbboilerplate.Close()
}