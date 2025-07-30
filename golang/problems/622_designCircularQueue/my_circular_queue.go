package code

// ref https://leetcode.com/problems/design-circular-queue/

// Example 1:
// Input
// ["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// Output
// [null, true, true, true, false, 3, true, true, true, 4]

// Explanation
// MyCircularQueue myCircularQueue = new MyCircularQueue(3);
// myCircularQueue.enQueue(1); // return True
// myCircularQueue.enQueue(2); // return True
// myCircularQueue.enQueue(3); // return True
// myCircularQueue.enQueue(4); // return False
// myCircularQueue.Rear();     // return 3
// myCircularQueue.isFull();   // return True
// myCircularQueue.deQueue();  // return True
// myCircularQueue.enQueue(4); // return True
// myCircularQueue.Rear();     // return 4

type MyCircularQueue struct {
	sizeLimit    int
	queue        []int
	currentIndex int
	currentSize  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		sizeLimit: k,
		queue:     make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currentSize == this.sizeLimit {
		return false
	}

	this.queue[this.currentIndex] = value
	this.currentIndex = (this.currentIndex + 1) % this.sizeLimit
	this.currentSize++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.currentSize == 0 {
		return false
	}

	this.currentSize -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.currentSize == 0 {
		return -1
	}

	offset := (this.sizeLimit - this.currentSize)
	frontIndex := (this.currentIndex + offset) % this.sizeLimit
	return this.queue[frontIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.currentSize == 0 {
		return -1
	}

	offset := this.sizeLimit - 1
	rearIndex := (this.currentIndex + offset) % this.sizeLimit
	return this.queue[rearIndex]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.currentSize != 0 {
		return false
	}

	return true
}

func (this *MyCircularQueue) IsFull() bool {
	if this.currentSize != this.sizeLimit {
		return false
	}

	return true
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
