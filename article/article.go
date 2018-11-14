package article

//Article Articulo
type Article struct {
	Articleid string `json:"articleid" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

//New Crea un nuevo articulo.
func New() *Article {
	return &Article{
		Articleid: "",
		Quantity:  0,
	}
}
