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
	return nil, fmt.Errorf("Not implemented yet")
}

func (fc *funCSSPGX) Create(FunCSS) error {
	_, err := fc.pool.Exec(context.Background(), "")
	return err
}

//func (fc *funCSSPGX) ByID(id int) (*FunCSS, error) {
//
//}
