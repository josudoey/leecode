package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RandomizedSet", func() {
	var this RandomizedSet

	BeforeEach(func() {
		this = Constructor()
	})

	Describe("Insert", func() {
		var (
			val    int
			result bool
		)

		JustBeforeEach(func() {
			result = this.Insert(val)
		})

		Context("happy path", func() {
			BeforeEach(func() {
				val = 1
			})

			It("result matched", func() {
				Expect(result).To(BeTrue())
			})
		})
	})

	Describe("Remove", func() {
		var (
			val    int
			result bool
		)

		JustBeforeEach(func() {
			result = this.Remove(val)
		})

		Context("happy path", func() {
			BeforeEach(func() {
				val = 1
				this.Insert(val)
			})

			It("result matched", func() {
				Expect(result).To(BeTrue())
			})
		})
	})

	Describe("GetRandom", func() {
		var result int

		JustBeforeEach(func() {
			result = this.GetRandom()
		})

		Context("happy path", func() {
			Context("insert(1)", func() {
				BeforeEach(func() {
					this.Insert(1)
				})

				It("result matched", func() {
					Expect(result).To(Equal(1))
				})

				Context("insert(2)", func() {
					BeforeEach(func() {
						this.Insert(2)
					})

					It("result matched", func() {
						Expect(result).To(Or(Equal(1), Equal(2)))
					})

					Context("remove(1)", func() {
						BeforeEach(func() {
							this.Remove(1)
						})

						It("result matched", func() {
							Expect(result).To(Equal(2))
						})
					})
				})
			})
		})
	})
})

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
