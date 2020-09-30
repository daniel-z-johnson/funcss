package models

type FunCSS struct {
	ID     int64
	CSSHex string
	Name   string
}

type FunCSSDB interface {
	Page(page int, limit int, name string) ([]*FunCSS, error)

	Create(funCSS *FunCSS) error
	Update(funCSS *FunCSS) error
	Delete(funCSS *FunCSS) error
}
