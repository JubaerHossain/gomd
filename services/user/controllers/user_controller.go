package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/JubaerHossain/gomd/services/user/validation"
	"github.com/JubaerHossain/gomd/services/user/services"
	"net/http"
)
func UserIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")
        status := c.DefaultQuery("status", "")

        var filter map[string]interface{} = make(map[string]interface{})
        filter["page"] = page
        filter["limit"] = limit
        filter["status"] = status

        users, paginate := services.AllUser(filter)

        gomd.Res.Code(200).Data(users).Raw(map[string]interface{}{
            "meta": paginate,
        }).Json(c)
	}
}


func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser validation.CreateUserRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createUser); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user := services.CreateAUser(createUser)

		gomd.Res.Code(http.StatusCreated).Message("success").Data(user).Json(c)
	}
}


func UserShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		userId := c.Param("userId")

		user := services.AUser(userId)

		gomd.Res.Code(http.StatusOK).Message("success").Data(user).Json(c)
	}
}


func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser validation.UpdateUserRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")

		if err := c.ShouldBind(&updateUser); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user, err := services.UpdateAUser(userId, updateUser)

		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(user).Json(c)
	}
}


func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")
		err := services.DeleteAUser(userId)

		if !err {
			gomd.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}