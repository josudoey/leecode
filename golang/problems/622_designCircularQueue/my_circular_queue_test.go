package code

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("MyCircularQueue",
	func(ops []string, args [][]int, expected interface{}) {
		var (
			this   MyCircularQueue
			result []interface{}
		)

		result = make([]interface{}, len(ops))
		for i, op := range ops {
			switch op {
			case "MyCircularQueue":
				k := args[i][0]
				this = Constructor(k)
				result[i] = nil
			case "enQueue":
				value := args[i][0]
				result[i] = this.EnQueue(value)
			case "deQueue":
				result[i] = this.DeQueue()
			case "Front":
				result[i] = this.Front()
			case "Rear":
				result[i] = this.Rear()
			case "IsEmpty":
				result[i] = this.IsEmpty()
			case "isFull":
				result[i] = this.IsFull()
			}
		}

		log.Printf("%v", result)
		Expect(result).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]string{"MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"},
		[][]int{{3}, {1}, {2}, {3}, {4}, {}, {}, {}, {4}, {}},
		[]interface{}{nil, true, true, true, false, 3, true, true, true, 4},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
