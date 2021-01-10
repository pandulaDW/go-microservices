package helpers

import (
	"io/ioutil"
	"os"
	"strings"
)

// DotEnv will read the .env file in the root path and
// will set the environment variables
func DotEnv() {
	content, _ := ioutil.ReadFile(".env")

	for _, line := range strings.Split(string(content), "\n") {
		envVar := strings.Split(line, "=")
		os.Setenv(strings.TrimSpace(envVar[0]), strings.TrimSpace(envVar[1]))
	}
}
