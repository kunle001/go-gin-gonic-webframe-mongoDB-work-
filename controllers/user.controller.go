package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunle001/gogingonic/models"
	"github.com/kunle001/gogingonic/services"
	"github.com/kunle001/gogingonic/utils"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := u.UserService.CreateUser(&user)

	if err != nil {
		utils.SendError(ctx, 400, err)
		return
	}

	utils.SendSuccess(ctx, 200, "user created")
}

func (u *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := u.UserService.GetUser(&username)

	if err != nil {
		utils.SendError(ctx, 400, err)
		return
	}

	ctx.JSON(200, user)
}

func (u *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendError(ctx, http.StatusBadGateway, err)
		return
	}

	if err := u.UserService.UpdateUser(&user); err != nil {
		utils.SendError(ctx, 400, err)
		return
	}

	utils.SendSuccess(ctx, 200, "data updated")
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	if err := u.UserService.DeleteUser(&username); err != nil {
		utils.SendError(ctx, 400, err)
		return
	}
	utils.SendSuccess(ctx, 200, "user deleted")
}

func (uc *UserController) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/", uc.GetAll)
	userroute.GET("/:name", uc.GetUser)
	userroute.PATCH("/:name", uc.UpdateUser)
	userroute.DELETE("/:name", uc.DeleteUser)
}
