package controllers

import (
	"local/utils"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func Init() {
	if logger == nil {
		logger = utils.NewLogger("controler")
	}
}
