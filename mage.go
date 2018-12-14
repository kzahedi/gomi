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
func MacOSSharedLibrary() {
	fmt.Println("Building shared library for macOS")
	parent, _ := os.Getwd()
	os.Chdir("apps/libgomi/")
	env := make(map[string]string)
	env["GOOS"] = "darwin"
	env["GOARCH"] = "amd64"
	sh.RunWith(env, "go", "build", "-o", "libgomi.so", "-buildmode=c-shared", "main.go")
	sh.Run("mv", "libgomi.so", "../../bin/macos")
	os.Chdir(parent)
}

// Builds a cpp example that uses the shared library
func MacOSCExample() {
	fmt.Println("Building C++ example, using the shared library for macOS")
	parent, _ := os.Getwd()
	os.Chdir("apps/cpp/")
	sh.Run("g++", "main.cpp", "-o", "main", "-I../../apps/libgomi", "-L../../apps/libgomi", "-lgomi")
	// sh.Run("mv", "libgomi.so", "bin/macos")
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

// Build for shared library for linux
// func LinuxSharedLibrary() {
// 	fmt.Println("Building shared library for linux")
// 	env := make(map[string]string)
// 	env["GOOS"] = "linux"
// 	env["GOARCH"] = "amd64"
// 	sh.RunWith(env, "go", "build", "-o", "libgomi.so", "-buildmode=c-shared", "apps/gomi/main.go")
// 	sh.Run("mv", "libgomi.so", "bin/linux")
// }

// Build all targets
func All() {
	// mg.Deps(MacOS, Windows, Linux)
	MacOS() // macOS and linux generate the same file name
	// MacOSSharedLibrary()
	// MacOSCExample()
	Windows()
	Linux()
	// LinuxSharedLibrary()
}

var Default = All
