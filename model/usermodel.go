package model

type User struct {
	Id          string `json:"id"`          // User Id
	Name        string `json:"name"`        // User Name
	Dob         string `json:"dob"`         // User Date of Birth
	Address     string `json:"address"`     // User Address
	Description string `json:"description"` // User Description
	CreatedAt   string `json:"createdAt"`   // User Created Date
}
