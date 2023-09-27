package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	DOB      string `json:"dob" bson:"dob"`
	Role     []string `json:"role" bson:"role"`
	Status   string `json:"status" bson:"status"`
}

type DBResponse struct{
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	DOB      string `json:"dob" bson:"dob"`
	Role     []string `json:"role" bson:"role"`
	Status   string `json:"status" bson:"status"`
}

type Role struct{
	Role string `json:"role" bson:"role"`
	Access string `json:"access" bson:"access"`
	Responsibility string `json:"responsibility" bson:"responsibility"`
}

type RDBResponse struct{
	Role string `json:"role" bson:"role"`
	Access string `json:"access" bson:"access"`
	Responsibility string `json:"responsibility" bson:"responsibility"`
}
