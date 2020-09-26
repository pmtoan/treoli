package lib

type Account struct {
	Username string `header:"Username"`
	Fullname string `header:"Fullname"`
	Email    string `header:"Email"`
}

type Board struct {
	ID      string `header:"ID" json:"id"`
	Name    string `header:"Name" json:"name"`
	Desc    string `header:"Desc" json:"desc"`
	Closed  bool   `header:"Closed" json:"closed"`
	Pinned  bool   `header:"Pinned" json:"pinned"`
	Starred bool   `header:"Starred" json:"starred"`
	URL     string `header:"URL" json:"url"`
}
