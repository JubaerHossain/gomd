package routes

import (
	"github.com/JubaerHossain/gomd/gomd"
	c "github.com/JubaerHossain/gomd/services/user/controllers"
)

func UserSetup() {
    v1 := gomd.Router.Group("api/v1")
	v1.GET("users", c.UserIndex())
	v1.POST("users", c.UserCreate())
	v1.GET("users/:userId", c.UserShow())
	v1.PUT("users/:userId", c.UserUpdate())
	v1.DELETE("users/:userId", c.UserDelete())
}
