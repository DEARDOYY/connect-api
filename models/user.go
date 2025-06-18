package models

// User โครงสร้างข้อมูลผู้ใช้งาน
type User struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	Role         string `json:"role" bson:"role"`
	Phone        string `json:"phone" bson:"phone"`
	Is_active    bool   `json:"is_active" bson:"is_active"`
	Avatar       string `json:"avatar" bson:"avatar"`
	Created_date string `json:"created_date" bson:"created_date"`
	Created_user string `json:"created_user" bson:"created_user"`
	Updated_date string `json:"updated_date" bson:"updated_date"`
	Updated_user string `json:"updated_user" bson:"updated_user"`
}
