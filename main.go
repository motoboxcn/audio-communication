package audiocommunication

import "github.com/patsnapops/noop/log"

func init() {
	log.Default().WithLevel(log.DebugLevel).Init()
}
