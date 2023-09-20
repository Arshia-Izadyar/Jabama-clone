package common

import (
	"regexp"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

var pnRegex = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`
var log = logger.NewLogger(config.GetConfig())

func ValidateNumber(pn string) bool {
	matched, err := regexp.Match(pnRegex, []byte(pn))
	if err != nil {
		log.Error(logger.General, logger.MobileValidation, err, nil)

	}
	return matched
}
