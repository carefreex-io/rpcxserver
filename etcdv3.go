//+build etcdV3

package rpcxserver

import (
	"github.com/carefreeskyio/logger"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
)

func AddRegistryPlugin(s *server.Server, options *ServerOption) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: options.Network + "@" + options.ServerIp + ":" + options.Port,
		EtcdServers:    options.RegistryAddr,
		BasePath:       options.BasePath,
		UpdateInterval: options.UpdateInterval,
	}
	if err := r.Start(); err != nil {
		logger.Fatal(err)
	}
	s.Plugins.Add(r)
}
