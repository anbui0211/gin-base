package initialize

import (
	"gin-base/internal/config"
)

func mySQL() {
	config.ConnectDBEcommerce()
}
