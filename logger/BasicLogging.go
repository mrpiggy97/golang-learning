package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func BasicLogging() {
	logger := log.Sample(&zerolog.BasicSampler{N: 200})
	for i := 0; i < 1000; i++ {
		logger.Info().Msg(fmt.Sprintf("%v", i))
	}
}
