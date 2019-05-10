package new

import (
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/core/node"
	"github.com/ipfs/go-ipfs/repo"
)

type fxGroup struct {
	opts map[string]func()fx.Option
}

func (g *fxGroup) get() fx.Option {
	opts := make([]fx.Option, 0, len(g.opts))
	for _, opt := range g.opts {
		opts = append(opts, opt())
	}

	return fx.Options(opts...)
}


type settings struct {
	fx fxGroup
}

func Provide(i interface{}) func() fx.Option {
	return func() fx.Option {
		return fx.Provide(i)
	}
}

func defaults() settings {
	out := settings{
		fx: fxGroup{
			map[string]func()fx.Option{},
		},
	}
	out.fx.opts["goprocess"] = Provide(baseProcess)

	out.fx.opts["repo"] = Provide(memRepo)

	out.fx.opts["storage.config"] = Provide(repo.Repo.Config)
	out.fx.opts["storage.datastore"] = Provide(repo.Repo.Datastore)
	out.fx.opts["storage.blockstore.basic"] = Provide(node.BaseBlockstoreCtor(blockstore.DefaultCacheOpts(), false, false))
	out.fx.opts["storage.blockstore.final"] = Provide(node.GcBlockstoreCtor)



	// TEMP: setting global sharding switch here
	//TODO uio.UseHAMTSharding = cfg.Experimental.ShardingEnabled

	return out
}



type Option = func()

func New() {

}
