package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/svenjeppsson/travis-test/dal"
)

func main() {
	NewApp(dal.NewSQLDataAcesssLayer())

}
