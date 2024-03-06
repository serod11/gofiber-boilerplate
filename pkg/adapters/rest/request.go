package rest

import "github.com/serod11/gofiber-boilerplate/pkg/domain/command"

type BookRequest struct {
	Name string `json:"name"`
	Page uint   `json:"page"`
}

func (req BookRequest) ToCommand() command.CreateBookCommand {
	return command.CreateBookCommand{
		Name: req.Name,
		Page: req.Page,
	}
}
