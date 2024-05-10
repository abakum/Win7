package main

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"

	"golang.org/x/sys/windows"
)

func main() {
	log.SetFlags(log.Lshortfile)
	info, ok := debug.ReadBuildInfo()
	s := "go"
	if ok {
		s = fmt.Sprintf("%s", info.GoVersion)
	}
	majorVersion, minorVersion, buildNumber := windows.RtlGetNtVersionNumbers()
	log.Printf("Is %s work with Windows %d.%d.%d on %s", s, majorVersion, minorVersion, buildNumber, runtime.GOARCH)
}
