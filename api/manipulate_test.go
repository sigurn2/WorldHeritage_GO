package api

import (
	"testing"

	"github.com/sigurn2/WorldHeritage_GO/data"
)

func TestManipulate(t *testing.T) {
	AddCol(data.Attribute{
		Name:   "GONGXIFACAI",
		Values: nil,
	})
}
