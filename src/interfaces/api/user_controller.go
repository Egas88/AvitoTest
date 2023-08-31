package controllers

import (
	"github.com/Egas88/AvitoTest/src/domain"
	"github.com/Egas88/AvitoTest/src/interfaces/database"
	"github.com/Egas88/AvitoTest/src/usecase"
	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := domain.User{Name: c.Param("Name")}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []domain.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) GetUserSegmentsById(id string) (segments []string, err error) {
	res, err := controller.Interactor.GetSegmentsByUserId(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (controller *UserController) UpdateUserSegments(SegmentsToAdd []string, SegmentsToDelete []string, idUser string) error {
	err := controller.Interactor.UpdateUserSegments(
		SegmentsToAdd,
		SegmentsToDelete,
		idUser,
	)
	if err != nil {
		return err
	}
	return nil
}

func (controller *UserController) Delete(id string) error {
	err := controller.Interactor.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
