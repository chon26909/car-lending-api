package entity

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	Id                  uint      `gorm:"primaryKey"`
	LicensePlate        string    `gorm:"type:varchar(20);not null;unique"` // ทะเบียนรถ
	Brand               string    `gorm:"type:varchar(50);not null"`        // ยี่ห้อ
	Model               string    `gorm:"type:varchar(50);not null"`        // รุ่น
	Year                int       `gorm:"not null"`                         // ปีที่ผลิต
	Color               string    `gorm:"type:varchar(30)"`                 // สี
	Category            string    `gorm:"type:varchar(30);not null"`        // ประเภท (Sedan, SUV)
	VIN                 string    `gorm:"type:varchar(50);unique"`          // หมายเลขประจำตัวรถ
	Status              string    `gorm:"type:varchar(20);not null"`        // สถานะ (Available, Rented)
	Mileage             int       `gorm:"not null"`                         // ระยะทางสะสม
	FuelLevel           string    `gorm:"type:varchar(20)"`                 // ระดับน้ำมัน
	RentalPricePerDay   float64   `gorm:"not null"`                         // ราคาเช่าต่อวัน
	RentalPricePerWeek  float64   `gorm:"not null"`                         // ราคาเช่าต่อสัปดาห์
	RentalPricePerMonth float64   `gorm:"not null"`                         // ราคาเช่าต่อเดือน
	LateFee             float64   `gorm:"not null"`                         // ค่าปรับคืนรถช้า
	Features            string    `gorm:"type:text"`                        // คุณสมบัติพิเศษ
	Image               string    `gorm:"type:varchar(255)"`                // รูปภาพรถ
	InsuranceDetails    string    `gorm:"type:varchar(255)"`                // ข้อมูลประกัน
	RegistrationExpiry  time.Time // วันหมดอายุการจดทะเบียน
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"` // Soft Delete
}
