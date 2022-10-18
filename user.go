package restapi

type User struct {
	Id           int    `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

// CreateUserDTO D-data, T-transfer, O-object
type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password "`
}

type Users struct {
	UserList []string
}
