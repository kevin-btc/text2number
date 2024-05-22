package alphatodigit

import (
	"fmt"
	"os/exec"
	"strings"
)

func Alpha2Digit(sentence string) (string, error) {
    cmd := exec.Command("python", "text_to_num/search_and_replace_by_num.py",sentence)

    output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

    return strings.Trim(string(output), "\"'"), nil
}