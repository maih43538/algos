#### Challenge: Implement a Set of unique elements without using the built-in array or hash data structures


Hash table - an array coupled with a hash function

##### Requirements

* Hash table must handle collisions
* There are operations to insert and delete
* Hash table is resized when load balance is too high or too low

##### Hashing Algorithm

* Given an input, a unique output will be returned
* A good one will have items evenly distributed across buckets

##### Hash Table Design

* Implement table using a slice of buckets
* A bucket is a linked list structure
* An element's hash value is used to index the slice of buckets

##### Hash Table Methods

##### Insertion
1. Compute a hash code against the given key. Note that two different keys may have the same hash code.
2. The hash code is used to specify the index of a bucket in the slice
		hash(key) % array_length
3. At this index, there is a linked list of keys and values. Store the key and value in this index.
4. TODO: Check for resizing.

##### Retrieval
1. Compute the hash code from the key and then compute the index from the hash code. Access element
    at index, if bucket is linked, traverse through the linked list for the value with this key.

##### Deletion
1. Compute index from key.
2. Access bucket or traverse linked list until element is found.
3. TODO: Check for resizing.

##### TODO: Growing
1. Allocate a new table with twice the size
2. Evenly evacuate key/value pairs from the old bucket
3. Copy it over to the new bucket and delete from old bucket
4. Replace original table with resized table