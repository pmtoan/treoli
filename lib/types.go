package lib

type Account struct {
	Username string `header:"Username"`
	Fullname string `header:"Fullname"`
	Email    string `header:"Email"`
}
