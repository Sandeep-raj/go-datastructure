package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Design Lru Cache

1. Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

2. Discards the least recently used items first. This algorithm requires keeping track of what
   was used when, which is expensive if one wants to make sure the algorithm always discards
   the least recently used item. General implementations of this technique require keeping
   "age bits" for cache-lines and track the "Least Recently Used" cache-line based on age-bits.
   In such an implementation, every time a cache-line is used, the age of all other cache-lines changes

3. Implement the LRUCache class:
      3.1 LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
      3.2 int get(int key) Return the value of the key if the key exists, otherwise return -1.
      3.3 void put(int key, int value) Update the value of the key if the key exists. Otherwise,
          add the key-value pair to the cache. If the number of keys exceeds the capacity from this
          operation, evict the least recently used key.

4. Could you do get and put in O(1) time complexity?

Sample Input
5 2
0 5 2
0 6 8
0 5 3
0 4 7
1 5

Sample Output
3
*/

type lrucache struct {
	max     int
	lrulist *utils.Dll
	hash    map[int]*utils.Dll_node
}

func initlru(size int) *lrucache {
	return &lrucache{
		max:     size,
		lrulist: utils.InitDll(),
		hash:    make(map[int]*utils.Dll_node),
	}
}

func (lru *lrucache) get(key int) int {
	if lru.hash[key] == nil {
		return -1
	} else {
		node := lru.hash[key]
		val := node.Data

		lru.lrulist.DllRemoveNode(node)
		lru.lrulist.DllInsertRear(node)

		return val.(int)
	}
}

func (lru *lrucache) put(key int, value int) {
	if lru.hash[key] != nil {
		node := lru.hash[key]
		node.Data = value

		lru.lrulist.DllRemoveNode(node)
		lru.lrulist.DllInsertRear(node)
	} else {
		node := utils.InitDllNode(key, value)
		lru.hash[key] = node
		lru.lrulist.DllInsertRear(node)

		if lru.lrulist.Count > lru.max {
			evictnode, _ := lru.lrulist.DllRemoveFront()
			delete(lru.hash, evictnode.Key)
		}
	}
}

func TestLRUCache() {
	lru := initlru(3)

	lru.put(1, 1)
	lru.put(2, 2)
	lru.put(2, 1)
	lru.put(3, 3)
	lru.put(1, 2)
	log.Print(lru.get(2))
	lru.lrulist.PrintDll()
	log.Print("---------------------------------------------------")

	lru.put(4, 1)
	lru.put(5, 1)
	log.Print(lru.get(4))
	lru.lrulist.PrintDll()
	log.Print("---------------------------------------------------")

	lru.put(1, 1)
	lru.put(6, 1)
	lru.put(7, 1)

	lru.lrulist.PrintDll()
	log.Print("---------------------------------------------------")

}
