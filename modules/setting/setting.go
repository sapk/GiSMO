package setting

import (
    "runtime"
)

var (
	// App settings.
	AppVer    string
	ApiVer    string
    
	IsWindows bool
)

func init() {
	IsWindows = runtime.GOOS == "windows"
}