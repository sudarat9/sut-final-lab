package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	Customer := Customer{
		Name:       "", // ต้องไม่เป็นค่าว่าง
		Email:      "sudarat@gmail.com",
		CustomerID: "Mserwtyd",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Customer)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Name not blank"))
}

// package entity

// import (
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestPositive(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	Customer := Customer{
// 		Name:       "sudaret", // ต้องไม่เป็นค่าว่าง
// 		Email:      "sudarat@gmail.com",
// 		CustomerID: "Mserwtyd",
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(Customer)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 	g.Expect(ok).To(BeTrue())

// 	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 	g.Expect(err).To(BeNil())

// 	// err.Error ต้องมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal("ok"))
// }

// package entity

// import (
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestNegativCcustomer(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	Customer := Customer{
// 		Name:       "sudaret", // ต้องไม่เป็นค่าว่าง
// 		Email:      "sudarat@gmail.com",
// 		CustomerID: "Mseyd",
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(Customer)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 	g.Expect(ok).To(BeTrue())

// 	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 	g.Expect(err).To(BeNil())

// 	// err.Error ต้องมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal("ok"))
// }
