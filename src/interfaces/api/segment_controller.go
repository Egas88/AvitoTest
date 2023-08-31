package controllers

import (
	"github.com/Egas88/AvitoTest/src/domain"
	"github.com/Egas88/AvitoTest/src/interfaces/database"
	"github.com/Egas88/AvitoTest/src/usecase"
	"github.com/labstack/echo"
)

type SegmentController struct {
	Interactor usecase.SegmentInteractor
}

func NewSegmentController(sqlHandler database.SqlHandler) *SegmentController {
	return &SegmentController{
		Interactor: usecase.SegmentInteractor{
			SegmentRepository: &database.SegmentRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SegmentController) Create(c echo.Context) {
	s := domain.Segment{Name: c.Param("name")}
	c.Bind(&s)
	controller.Interactor.Add(s)
	createdSegments := controller.Interactor.GetInfo()
	c.JSON(201, createdSegments)
	return
}

func (controller *SegmentController) GetSegment() []domain.Segment {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *SegmentController) Delete(id string) error {
	return controller.Interactor.Delete(id)
}

func (controller *SegmentController) DeleteByName(name string) error {
	return controller.Interactor.DeleteByName(name)
}
