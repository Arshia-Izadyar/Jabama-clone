package dto

type RegisterByUsername struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password" binding:"required"`
}

type LoginByUserName struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id           int                    `json:"id"`
	Username     string                 `json:"username"`
	Email        string                 `json:"email,omitempty"`
	FirstName    string                 `json:"firstName,omitempty"`
	LastName     string                 `json:"lastName,omitempty"`
	Password     string                 `json:"password"`
	Activated    bool                   `json:"activated"`
	UserWishList []UserWishListResponse `json:"userWishList"`
}

type OtpRequest struct {
	PhoneNumber string `json:"phone_number" binding:"phone"`
}

type OtpDto struct {
	Value string
	Used  bool
}

type RegisterLoginByPhone struct {
	PhoneNumber string `json:"phone_number"`
	Otp         string `json:"otp"`
}
