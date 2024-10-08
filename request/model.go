package request

import (
	"io"

	"github.com/wesleyfantinel/cohere-go/model"
)

type Model struct {
	Model model.Model `json:"-"`
}

func (c *Model) Path() (string, error) {
	return "/models/" + c.Model.String(), nil
}

// nolint: nilnil
func (c *Model) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *Model) ContentType() string {
	return ""
}
