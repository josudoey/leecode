package code

import (
	"math/rand"
)

// ref https://leetcode.com/problems/insert-delete-getrandom-o1/

// Example 1:
// Input
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// Output
// [null, true, false, true, 2, true, false, 2]

// Explanation
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
// randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
// randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
// randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
// randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
// randomizedSet.insert(2); // 2 was already in the set, so return false.
// randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.

type RandomizedSet struct {
	items    []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{indexMap: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexMap[val]
	if ok {
		return false
	}

	this.items = append(this.items, val)
	this.indexMap[val] = len(this.items) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	toBeRemovedIndex, ok := this.indexMap[val]
	if !ok {
		return false
	}

	lastItemIndex := len(this.items) - 1
	lastItem := this.items[lastItemIndex]
	this.items[toBeRemovedIndex] = lastItem
	this.indexMap[lastItem] = toBeRemovedIndex

	this.items = this.items[:lastItemIndex]
	delete(this.indexMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	count := len(this.items)
	index := rand.Intn(count)
	return this.items[index]
}
