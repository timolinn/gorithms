package datastructures

// ARRAY_LENGTH is the length of the
// underlying array
const ARRAY_LENGTH = 100

// HashTable is a data structure type
type HashTable [ARRAY_LENGTH]*LinkedList

type listData struct {
	key   string
	value interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

// Put adds a key value into a hashtable
func (ht *HashTable) Put(key string, value int) *HashTable {
	hash := hash(key)
	index := index(hash)

	if ht[index] == nil {
		ht[index] = NewLinkedList()
		ht[index].Add(&listData{key, value})
	} else {
		node := ht[index].Head
		for {
			if node != nil {
				d := node.Value.(*listData)
				if d.key == key {
					d.value = value
					break
				}
			} else {
				ht[index].Add(&listData{key, value})
				break
			}
			node = node.Next
		}
	}
	return ht
}

// Get return the value of the specified key
func (ht *HashTable) Get(key string) (interface{}, bool) {
	hash := hash(key)
	index := index(hash)
	linkedList := ht[index]

	if linkedList == nil {
		return nil, false
	}

	root := linkedList.Head
	for root != nil {
		d := root.Value.(*listData)
		if d.key == key {
			return d.value, true
		}
		root = root.Next
	}
	return nil, false
}

// using horner's method
func hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}

func index(h int) int {
	return h % ARRAY_LENGTH
}
