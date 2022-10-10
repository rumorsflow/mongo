package mongo

import (
	"context"
	"github.com/roadrunner-server/errors"
	"github.com/rumorsflow/contracts/config"
	"github.com/rumorsflow/mongo-ext"
	"go.mongodb.org/mongo-driver/mongo"
)

const PluginName = "mongo"

type Plugin struct {
	cfg *Config
	mdb *mongo.Database
}

func (p *Plugin) Init(cfg config.Configurer) error {
	const op = errors.Op("mongo plugin init")

	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	var err error

	if err = cfg.UnmarshalKey(PluginName, &p.cfg); err != nil {
		return errors.E(op, errors.Init, err)
	}

	p.mdb, err = mongoext.GetDB(context.Background(), p.cfg.URI)
	if err != nil {
		return errors.E(op, errors.Init, err)
	}

	return nil
}

// Name returns user-friendly plugin name
func (p *Plugin) Name() string {
	return PluginName
}

// Provides declares factory methods.
func (p *Plugin) Provides() []any {
	return []any{
		p.ServiceMongoDB,
	}
}

func (p *Plugin) ServiceMongoDB() *mongo.Database {
	return p.mdb
}
