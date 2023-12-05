package main

import (
	"fmt"
	"runtime"
)

func main() {
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("No es windows es", os)
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es linux")
	case "darwin":
		fmt.Println("esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

}
