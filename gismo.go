package main

import (
	"runtime"

	/*"github.com/codegangsta/cli"*/ //TODO

	"github.com/sapk/GiSMO/modules/setting"
	"github.com/sapk/GiSMO/modules/web"
)

const APP_VER = "0.0.1.0001 Alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}

func main() {
	web.Start(":4000")
}

/*
func main() {
    e := echo.New()

    e.Static("/assets/","public/");

    // New permissions middleware
    perm, err := permissionbolt.New()
    if err != nil {
        log.Fatalln(err)
    }

    // Blank slate, no default permissions
    perm.Clear()

    //TODO set perm

    // Set up a middleware handler for Echo, with a custom "permission denied" message.
    permissionHandler := func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c *echo.Context) error {
            // Check if the user has the right admin/user rights
            if perm.Rejected(c.Response().Writer(), c.Request()) {
                // Deny the request
                return echo.NewHTTPError(http.StatusForbidden, "Permission denied!")
            }
            // Continue the chain of middleware
            return next(c)
        }
    }

    // Logging middleware
    e.Use(mw.Logger())

    // Enable the permissions middleware, must come before recovery
    e.Use(permissionHandler)

    // Recovery middleware
    e.Use(mw.Recover())

    // Get the userstate, used in the handlers below
    userstate := perm.UserState()

    e.Get("/login", func(c *echo.Context) error {
        // Headers will be written, for storing a cookie
        userstate.Login(c.Response().Writer(), "bob")
        return c.String(http.StatusOK, fmt.Sprintf("bob is now logged in: %v\n", userstate.IsLoggedIn("bob")))
    })

    // Serve
    e.Run(":3000")

}
*/
