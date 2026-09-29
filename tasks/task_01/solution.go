package main

import(
	"strings"
	"fmt"
)

func greet(name string) string {
	clean_name := strings.TrimSpace(name)

	if clean_name == "" {
		clean_name = "World"
	}
	return fmt.Sprintf("Hello, %s!", clean_name)
}
