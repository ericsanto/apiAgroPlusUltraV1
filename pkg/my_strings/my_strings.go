package mystrings

import (
	"fmt"
	"strings"
)

func TakePartOfAText(content, delimiterStart, delimiterEnd string) (string, error) {

	start := strings.Index(content, delimiterStart)

	if start == -1 {
		return "", fmt.Errorf("substring informada nao existe")
	}

	remainder := content[start+len(delimiterStart):]

	end := strings.Index(remainder, delimiterEnd)

	if end == -1 {
		return "", fmt.Errorf("delimitador de fim não encontrado")
	}

	jsonStr := strings.TrimSpace(remainder[:end])

	return jsonStr, nil
}
