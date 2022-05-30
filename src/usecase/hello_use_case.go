package usecase

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type helloUseCase struct {
}

func NewHelloUseCase() domain.HelloUseCase {
	newHelloUseCase := helloUseCase{}

	return &newHelloUseCase
}

func (u *helloUseCase) Execute() entity.Hello {
	newHello := entity.Hello{
		Message: "hello world",
	}

	return newHello
}
