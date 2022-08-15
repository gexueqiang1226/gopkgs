package gopkgs

import "fmt"

func Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}