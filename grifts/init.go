package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/h3rby7/kebap/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
