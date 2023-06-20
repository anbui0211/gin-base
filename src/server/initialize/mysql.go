package initialize

import "gin-base/src/database"

func mySQL() {
	database.ConnectDBEcommerce()
}
