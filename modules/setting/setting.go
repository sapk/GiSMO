package setting

import (
    "runtime"
)

var (
	// App settings.
	AppVer    string
    
	IsWindows bool
)

func init() {
	IsWindows = runtime.GOOS == "windows"
}