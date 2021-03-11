package models

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FunCSS struct {
	UUID   uuid.UUID
	ID     int
	CSSHex string
	Name   string
	Author string
}

type FunCSSDB interface {
	Page(page int, limit int, name string) ([]*FunCSS, error)

	ByID(id int) (*FunCSS, error)

	Create(funCSS *FunCSS) error
	Update(funCSS *FunCSS) error
	Delete(funCSS *FunCSS) error
}

type funCSSPGX struct {
	pool pgxpool.Pool
}

func (fc *funCSSPGX) Page(page int, limit int, name string) ([]*FunCSS, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (fc *funCSSPGX) Create(funCSS FunCSS) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil
	}
	_, err = fc.pool.Exec(context.Background(), "INSERT INTO funcss(uuid, id, csshex, name, author) VALUES "+
		"	($1,  $2,        $3,            $4,          $5)",
		uuid, funCSS.ID, funCSS.CSSHex, funCSS.Name, funCSS.Author)
	return err
}

func (fc *funCSSPGX) ByID(id int) (*FunCSS, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (fc *funCSSPGX) Update(funCSS *FunCSS) error {
	return fmt.Errorf("not implemented yet")
}

func (fc *funCSSPGX) Delete(funCSS *FunCSS) error {
	return fmt.Errorf("not implemented yet")
}
