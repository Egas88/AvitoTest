package infrastructure

import (
	"net/http"
	"strings"

	controllers "github.com/Egas88/AvitoTest/src/interfaces/api"
	"github.com/labstack/echo"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())
	segmentController := controllers.NewSegmentController(NewSqlHandler())

	e.GET("/user_segments/:id", func(ctx echo.Context) error {
		segments, err := userController.GetUserSegmentsById(ctx.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to get user segments")
		}
		ctx.Bind(&segments)
		return ctx.JSON(http.StatusOK, segments)
	})

	e.GET("/users", func(ctx echo.Context) error {
		users := userController.GetUser()
		ctx.Bind(&users)
		return ctx.JSON(http.StatusOK, users)
	})

	e.POST("/users/:name", func(c echo.Context) error {
		userController.Create(c)
		return c.JSON(http.StatusOK, "created")
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		err := userController.Delete(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete user")
		}
		return c.JSON(http.StatusOK, "deleted")
	})

	e.PUT("/users/:SegmentsToAdd/:SegmentsToDelete/:UserId", func(c echo.Context) error {
		id := c.Param("UserId")
		segmentsToAdd := strings.Split(c.Param("SegmentsToAdd"), " ")
		segmentsToDelete := strings.Split(c.Param("SegmentsToDelete"), " ")
		err := userController.UpdateUserSegments(segmentsToAdd, segmentsToDelete, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to update user segments")
		}
		return c.JSON(http.StatusOK, "updated")
	})

	e.GET("/segments", func(ctx echo.Context) error {
		segments := segmentController.GetSegment()
		ctx.Bind(&segments)
		return ctx.JSON(http.StatusOK, segments)
	})

	e.POST("/segments/:name", func(c echo.Context) error {
		segmentController.Create(c)
		return c.JSON(http.StatusOK, "created")
	})

	e.DELETE("/segments/DeleteById/:id", func(c echo.Context) error {
		id := c.Param("id")
		err := segmentController.Delete(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete segment")
		}
		return c.JSON(http.StatusOK, "deleted")
	})

	e.DELETE("/segments/:name", func(c echo.Context) error {
		name := c.Param("name")
		err := segmentController.DeleteByName(name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete segment")
		}
		return c.JSON(http.StatusOK, "deleted")
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
