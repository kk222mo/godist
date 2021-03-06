package portforwarder

import (
	"regexp"
	"strings"
)

func ParseSSDPResponse(response string) map[string]string {
	res := make(map[string]string)
	lines := strings.Split(response, "\r\n")
	for _, line := range lines {
		params := strings.Split(line, ": ")
		if len(params) >= 2 && params[0] != "" {
			res[params[0]] = strings.Join(params[1:], ": ")
		}
	}
	return res
}

func ParseUPNPInterfaces(data string, prefix string, interfaces *map[string]bool) {
	r := regexp.MustCompile(`<controlURL>([0-9a-zA-Z/: ]+)</controlURL>`)
	matches := r.FindAllStringSubmatch(data, -1)
	for _, v := range matches {
		(*interfaces)[prefix+v[1]] = true
	}
}
