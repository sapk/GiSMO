package user

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo" 
    "github.com/xyproto/pinterface" 
)

//RegiterRouteHandler
func RegiterRouteHandler(user *echo.Group, userstate pinterface.IUserState) {
    user.Post("/login", func(c *echo.Context) error {
        username := c.Form("username") //TODO clean var
        if (userstate.CorrectPassword(username , c.Form("password"))){
            userstate.Login(c.Response().Writer(), username)
            return c.String(http.StatusOK, fmt.Sprintf("%v is now logged in: %v\n", username, userstate.IsLoggedIn(username))) //TODO use JSON
        }else{
            return c.String(http.StatusUnauthorized, fmt.Sprintf("%v is not logged in: %v\n", username, userstate.IsLoggedIn(username))) //TODO use JSON        
        }
    })
    user.Get("/logout", func(c *echo.Context) error {
        username := userstate.Username(c.Request()) //TODO check for error
        userstate.Logout(username)
        userstate.ClearCookie(c.Response().Writer())
        return c.String(http.StatusOK, fmt.Sprintf("%v is now logged out: %v\n",username , !userstate.IsLoggedIn("bob")))     
    })
    user.Get("/_current/status", func(c *echo.Context) error {
        username := userstate.Username(c.Request()) //TODO check for error
        if (userstate.UserRights(c.Request())){
            return c.String(http.StatusOK, fmt.Sprintf("User (%v) is logged !\n",username)) //TODO use JSON
        }else{
            return c.String(http.StatusOK, fmt.Sprintf("User (%v) is not logged !\n",username)) //TODO use JSON
        }
    })
}