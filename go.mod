module gophers.dev/cmds/stock

require (
	github.com/davecgh/go-spew v1.1.1
	gophers.dev/pkgs/yahoo-finance v0.0.0
)

replace gophers.dev/pkgs/yahoo-finance => ../../pkgs/yahoo-finance

go 1.16
