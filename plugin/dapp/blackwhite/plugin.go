package blackwhite

import (
	"gitlab.33.cn/chain33/chain33/pluginmgr"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/blackwhite/commands"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/blackwhite/executor"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/blackwhite/rpc"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/blackwhite/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.BlackwhiteX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.BlackwhiteCmd,
		RPC:      rpc.Init,
	})
}