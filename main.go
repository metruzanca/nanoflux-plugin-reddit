// Command nanoflux-plugin-reddit serves the Reddit plugin over gRPC so
// nanoflux can load it as an external plugin.
//
// Build it and drop the binary into the host's plugins directory
// (NF_PLUGINS_DIR, ./plugins by default):
//
//	GOOS=linux CGO_ENABLED=0 go build -o ../nanoflux-plugin-reddit .
package main

import (
	goplugin "github.com/hashicorp/go-plugin"

	reddit "nanoflux-plugin-reddit/reddit"

	"github.com/metruzanca/nanoflux/pluginapi"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: pluginapi.Handshake,
		Plugins:         pluginapi.PluginSet(&reddit.Plugin{}),
		GRPCServer:      goplugin.DefaultGRPCServer,
	})
}
