package host

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddytls"
	"github.com/danlsgiga/caddy-net/caddynet/netserver"
)

func init() {
	caddy.RegisterPlugin("host", caddy.Plugin{
		ServerType: "net",
		Action:     setupHost,
	})
}

func setupHost(c *caddy.Controller) error {
	config := netserver.GetConfig(c)

	// Ignore call to setupHost if the key is not echo or proxy
	if c.Key != "echo" && c.Key != "proxy" {
		return nil
	}

	for c.Next() {
		if !c.NextArg() {
			return c.ArgErr()
		}
		config.Hostname = c.Val()

		if config.TLS == nil {
			config.TLS = &caddytls.Config{Hostname: c.Val()}
		} else {
			config.TLS.Hostname = c.Val()
		}

		if c.NextArg() {
			// only one argument allowed
			return c.ArgErr()
		}
	}

	return nil
}
