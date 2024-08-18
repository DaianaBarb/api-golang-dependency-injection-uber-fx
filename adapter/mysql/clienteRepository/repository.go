package repository

import (
	"database/sql"
	model "golang-uber-fx/core/domain"
)

type Irepository interface {
	SaveCliente(cliente *model.Cliente)
	DeleteCliente(cpf string)
	FindCliente(cpf string)
}

type Repository struct {
	db *sql.DB
}

// DeleteCliente implements Irepository.
func (r *Repository) DeleteCliente(cpf string) {
	print(" deletando clinte " + cpf)
}

// SaveCliente implements Irepository.
func (r *Repository) SaveCliente(cliente *model.Cliente) {
	print(" salvando cliente " + cliente.Cpf)
}

func (r *Repository) FindCliente(cpf string) {
	print(" encontrando cliente  " + cpf)
}

func NewRepository(db *sql.DB) Irepository {
	return &Repository{
		db: db,
	}

}
