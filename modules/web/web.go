package web

import (
    "log"
    "net/http"
    "time"
    "github.com/tylerb/graceful"
    "github.com/labstack/echo"
    mw "github.com/labstack/echo/middleware"
    "github.com/xyproto/permissionbolt"
    
	"github.com/sapk/GiSMO/routers/api"
)

var (
	perm *permissionbolt.Permissions
)

func Start(ip_port string) {
    e := echo.New()
    // init permissionbolt middleware
    perm = initPerm()
    // Get the userstate, used in the handlers below
    //userstate := perm.UserState()

    // Logging middleware
    e.Use(mw.Logger())
    // Enable the permissions middleware, must come before recovery
    e.Use(permissionHandler)
    // Recovery middleware
    e.Use(mw.Recover())

    e.Static("/assets/","public/");
    e.Index("public/index.html");
    e.Favicon("public/img/icon.png"); //TODO maybe generate a ico for fallback compatibility
    
    api.RegiterRouteHandler(e.Group("/api"),perm.UserState());
    
    // Serve
    graceful.ListenAndServe(e.Server(ip_port), 5*time.Second)
}

func initPerm() *permissionbolt.Permissions{
    p, err := permissionbolt.NewWithConf("data/user.db")
    if err != nil {
        log.Fatalln(err)
    }
    p.Clear();
    p.AddAdminPath("/api/admin");
    //p.AddUserPath("/api/user");//TODO manage login
    p.AddPublicPath("/")
    
    //Add default admin user 
    userstate := p.UserState()
    //userstate.RemoveUser("admin")
    if(!userstate.HasUser("admin")){
        userstate.AddUser("admin", "admin", "admin@localhost")
        userstate.MarkConfirmed("admin")
        userstate.SetAdminStatus("admin")
        log.Printf("User admin was created: %v\n", userstate.HasUser("admin"));
    }
    return p;
}

// Set up a middleware handler for Echo, with a custom "permission denied" message.
func permissionHandler(next echo.HandlerFunc) echo.HandlerFunc {
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