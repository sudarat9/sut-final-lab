package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"` // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string `valid:"matches((^[L][0-9]{7}$)|(^[M][0-9]{7}$)|(^[H][0-9]{7}))~CustomerID not ture"` // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}
