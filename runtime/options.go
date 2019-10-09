package runtime

import (
	"github.com/spf13/viper"
	"github.com/wagoodman/dive/dive"
)

type Options struct {
	Ci         bool
	Image      string
	Source     dive.ImageSource
	ExportFile string
	CiConfig   *viper.Viper
	BuildArgs  []string
}
