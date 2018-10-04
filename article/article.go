package article

type Article struct {
	Articleid string `json:"articleid" validate:"required"`
	Quantity  string `json:"quantity" validate:"required"`
}

func New() *Article {
	return &Article{
		Articleid: "",
		Quantity:  "",
	}
}
