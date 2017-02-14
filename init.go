package archive

import (
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.OnInit(func() {
		log = logger.New().WithField("pkg", "archive")
	})
}
