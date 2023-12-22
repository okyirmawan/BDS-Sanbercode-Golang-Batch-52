package models

import (
	"fmt"
	"strings"
)

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (p Phone) ShowInfo() string {
	colors := strings.Join(p.Colors, ", ")
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.Name, p.Brand, p.Year, colors)
}

type PhoneInfo interface {
	ShowInfo() string
}
