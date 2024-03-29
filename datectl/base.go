package datectl

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type DateCTL struct {
}

var debug = false

type Message struct {
	Message string `json:"message"`
}

func New(inst *DateCTL) *DateCTL {
	return inst
}

func cleanCommand(resp string, cmd *exec.Cmd, err error, debug ...bool) string {
	outAsString := strings.TrimRight(resp, "\n")
	if len(debug) > 0 {
		if debug[0] {
			log.Infof("cmd %s", cmd.String())
			log.Infof("path %s", cmd.Path)
			if err != nil {
				log.Errorf("err:%s", err.Error())
			} else {
				log.Infof("resp:%s", outAsString)
			}
		}
	}
	return outAsString
}
