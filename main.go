package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/kirtiPhegade/packer/numbers"
	"github.com/kirtiPhegade/packer/strings"
	"github.com/kirtiPhegade/packer/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}