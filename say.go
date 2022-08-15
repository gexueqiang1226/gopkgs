package gopkgs

import (
	"errors"
	"fmt"
)

func Say(name string, lang string) (string, error) {
	switch lang {
	case "zh":
		return fmt.Sprintf("你好, %s!!!", name), nil
		break
	case "en":
		return fmt.Sprintf("Hi, %s!!!", name), nil
	}
	return "", errors.New("unknown language")
}
