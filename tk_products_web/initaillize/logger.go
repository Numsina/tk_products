package initaillize

import "github.com/Numsina/tk_products/tk_products_srv/logger"

var l *logger.Logger

func InitLogger() *logger.Logger {
	if l == nil {
		return logger.NewLogger()
	}
	return l
}
