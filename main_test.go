package main

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Log("Start main()")
	main()
	t.Log("we jusfinished main()")
}
