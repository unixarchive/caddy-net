package caddynet

import (
	// plug in the server
	_ "github.com/danlsgiga/caddy-net/caddynet/netserver"
	// // plug in the standard directives
	_ "github.com/danlsgiga/caddy-net/caddynet/host"
)
