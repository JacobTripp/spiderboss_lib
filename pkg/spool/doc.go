// Package spool implements a struct to keep track of cable on the spool
// and provides methods to get the revolutions to retract/extend a specific
// amount of cable.
//
// # Spools
//
// You must first configure your spool. This package can handle spools with
// sloped ends and flat ends:
//
//	|\          /|     _          _
//	| \________/ |    | |        | |
//	|  ________  |    | |________| |
//	| /        \ |    |  ________  |
//	|/          \|    | |        | |
//	                  |_|        |_|
//
// And it can handle any size cable. It will make it's calculation base on an
// ideal spool, meaning this package assumes that the spool is always wound and
// unwound fully each layer before acting on the next layer.
package spool // import "github.com/JacobTripp/spiderboss/go/pkg/spool"
