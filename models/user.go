package models

// User โครงสร้างข้อมูลผู้ใช้งาน
type User struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
