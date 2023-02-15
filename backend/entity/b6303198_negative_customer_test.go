package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativCustomer(t *testing.T) {
	g := NewGomegaWithT(t)
	Customer := Customer{
		Name:       "sudaret",
		Email:      "sudarat@gmail.com",
		CustomerID: "M12469", // รหัสลูกค้าขึ้นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Customer)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("CustomerID not ture"))
}
