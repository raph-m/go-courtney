module github.com/dave/courtney

go 1.12

require (
	github.com/dave/astrid v0.0.0-20170323122508-8c2895878b14
	github.com/dave/brenda v1.1.0
	github.com/dave/patsy v0.0.0-20170606133301-2245ba804d71
	github.com/pkg/errors v0.9.1
	golang.org/x/tools v0.0.0-20201211185031-d93e913c1a58
)

replace github.com/dave/patsy latest => github.com/rubensayshi/patsy courtney-gomod-prep
