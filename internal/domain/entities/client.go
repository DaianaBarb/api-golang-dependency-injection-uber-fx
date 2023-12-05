package entities

import "time"

type Client struct {
	Id        int
	Name      string
	Active    bool
	Telefone  string
	CreatedAt time.Time
}
