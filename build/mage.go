//+build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// Build for macOS
func MacOS() {
	fmt.Println("Building gomi executable for macOS")
	parent, _ := os.Getwd()
	os.Chdir("apps/gomi/")
	env := make(map[string]string)
	env["GOOS"] = "darwin"
	env["GOARCH"] = "amd64"
	sh.RunWith(env, "go", "build")
	sh.Run("mv", "gomi", "../../bin/macos")
	os.Chdir(parent)
}

// Build shared C library for macOS
func MacOSCLib() {
	fmt.Println("Building shard library for macOS")
	parent, _ := os.Getwd()
	// os.Chdir("apps/gomi/")
	env := make(map[string]string)
	env["GOOS"] = "darwin"
	env["GOARCH"] = "amd64"
	sh.RunWith(env, "go", "build", "-buildmode=c-shared", "*.go")
	// sh.Run("mv", "gomi", "../../bin/macos")
	os.Chdir(parent)
}

// Build for Windows
func Windows() {
	fmt.Println("Building for Windows")
	parent, _ := os.Getwd()
	os.Chdir("apps/gomi/")
	env := make(map[string]string)
	env["GOOS"] = "windows"
	env["GOARCH"] = "amd64"
	sh.RunWith(env, "go", "build")
	sh.Run("mv", "gomi.exe", "../../bin/windows")
	os.Chdir(parent)
}

// Build for Linux
func Linux() {
	fmt.Println("Building for linux")
	parent, _ := os.Getwd()
	os.Chdir("apps/gomi/")
	env := make(map[string]string)
	env["GOOS"] = "linux"
	env["GOARCH"] = "amd64"
	sh.RunWith(env, "go", "build")
	sh.Run("mv", "gomi", "../../bin/linux")
	os.Chdir(parent)
}

// Build all targets
func All() {
	// mg.Deps(MacOS, Windows, Linux)
	MacOS() // macOS and linux generate the same file name
	MacOSCLib()
	Windows()
	Linux()
}

var Default = All
