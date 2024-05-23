package alphatodigit

import (
	"fmt"
	"os/exec"
	"strings"
)

func Alpha2Digit(sentence string) (string, error) {
    cmd := exec.Command("search_and_replace_by_num", sentence)

    output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

    return strings.Trim(string(output), "\"'"), nil
}