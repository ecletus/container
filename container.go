package container

import (
	"github.com/ecletus/cli"
	"github.com/ecletus/db"
	"github.com/ecletus/plug"
	"github.com/ecletus/sites"
)

type Container struct {
	Plugins    *plug.Plugins
	Options    *plug.Options
	Sites      *sites.SitesRouter
	SingleSite bool
	cli        *cli.CLI
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
	return container.Plugins.Init()
}

func (container *Container) InitDB() (err error) {
	container.dbInit = true
	return container.Plugins.TriggerPlugins(plug.NewPluginEvent(db.E_INIT))
}

func (container *Container) Migrate() (err error) {
	if !container.dbInit {
		err = container.InitDB()
		if err != nil {
			return err
		}
	}
	return container.Plugins.TriggerPlugins(plug.NewPluginEvent(db.E_MIGRATE))
}

func (container *Container) CLI() *cli.CLI {
	if container.cli == nil {
		container.cli = &cli.CLI{Dis: container.Plugins}
	}
	return container.cli
}
