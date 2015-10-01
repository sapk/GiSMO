package api

import (
    "net/http"
    "github.com/labstack/echo"   
)

//RegiterRouteHandler
func RegiterRouteHandler(api *echo.Group) {
    
		api.Get("/", home)
		api.Get("/status", status)
    /*
		e.Post("user/login", auth.SignIn)
		e.Get("user/logout", auth.LogOut)
*/
}

// Home 
func home(ctx *echo.Context) error {
    return ctx.String(http.StatusOK,"Home API");
}
// Status 
func status(ctx *echo.Context) error {
    return ctx.String(http.StatusOK,"OK");
}