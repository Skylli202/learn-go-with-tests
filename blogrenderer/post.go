package blogrenderer

import (
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func (p Post) SanitiseTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
