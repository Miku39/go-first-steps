module helloworld

go 1.19

require (
	github.com/myuser/calculator v0.0.0
	github.com/myuser/geometry v0.0.0
	rsc.io/quote v1.5.2
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/rs/zerolog v1.28.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
)

replace github.com/myuser/calculator => ../calculator
replace github.com/myuser/geometry => ../geometry
