package container

import (
	"context"

	"github.com/ecletus/plug"
	"github.com/ecletus/sites"

	"github.com/ecletus/db"
)

type Container struct {
	Plugins    *plug.Plugins
	Options    *plug.Options
	Sites      *sites.SitesRouter
	SingleSite bool
	dbInit     bool
}

func New(plugins *plug.Plugins) *Container {
	c := &Container{
		Options: plugins.Options(),
		Plugins: plugins,
	}
	return c
}

func (container *Container) Init() (err error) {
	if err = container.Plugins.ProvideOptions(); err != nil {
		return
	}
	if err = container.Plugins.Init(); err != nil {
		return
	}
	return container.InitDB(context.Background())
}

func (container *Container) InitDB(ctx context.Context) (err error) {
	if container.dbInit {
		return
	}
	container.dbInit = true
	return container.Plugins.TriggerPlugins(plug.NewPluginEvent(db.E_INIT, ctx))
}

func (container *Container) Migrate(ctx context.Context) (err error) {
	if !container.dbInit {
		err = container.InitDB(ctx)
		if err != nil {
			return err
		}
	}
	return container.Plugins.TriggerPlugins(plug.NewPluginEvent(db.E_MIGRATE, ctx))
}