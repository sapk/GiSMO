package web

import (
    "time"
    "github.com/tylerb/graceful"
    "github.com/labstack/echo"
    mw "github.com/labstack/echo/middleware"
	"github.com/sapk/GiSMO/routers/api"
)


func Start(ip_port string) {
    e := echo.New()

    // Logging middleware
    e.Use(mw.Logger())

    // Recovery middleware
    e.Use(mw.Recover())

    e.Static("/assets/","public/");
    e.Index("public/index.html");
    e.Favicon("public/img/icon.png"); //TODO maybe generate a ico for fallback compatibility
    
    api.RegiterRouteHandler(e.Group("/api"));
    

    //TODO routeurs
    
    // Serve
    graceful.ListenAndServe(e.Server(ip_port), 5*time.Second)
}