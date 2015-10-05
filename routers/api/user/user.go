package user

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo" 
    "github.com/xyproto/pinterface" 
)
/*
type User struct {
    FirstName string
    LastName string
}
*/
//RegiterRouteHandler
func RegiterRouteHandler(user *echo.Group, userstate pinterface.IUserState) {
    user.Post("/_login", func(c *echo.Context) error {
        username := c.Form("username") //TODO clean var
        if (userstate.CorrectPassword(username , c.Form("password"))){
            userstate.Login(c.Response().Writer(), username)
        return c.String(http.StatusOK, userState(true,userstate.IsAdmin(username),username))
        }else{
        return c.String(http.StatusOK, userState(false,false,""))        
        }
    })
    user.Get("/_logout", func(c *echo.Context) error {
        username := userstate.Username(c.Request()) //TODO check for error
        userstate.Logout(username)
        userstate.ClearCookie(c.Response().Writer())
        return c.String(http.StatusOK, userState(false,false,""))  
    })
    user.Get("/_current", func(c *echo.Context) error {
        username := userstate.Username(c.Request())
        return c.String(http.StatusOK, userState(userstate.UserRights(c.Request()),userstate.AdminRights(c.Request()),username))
    })
}

func userState(isLogged, isAdmin bool, username string) string {
    return fmt.Sprintf("{\"isLogged\":%v,\"isAdmin\":%v,\"username\":\"%v\"}",isLogged,isAdmin,username)
}