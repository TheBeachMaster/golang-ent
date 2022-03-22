package delivery

import (
	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/ent"
	"com.thebeachmaster/entexample/internal/delivery"
	"github.com/go-redis/redis/v8"
)

type deliveryRegistry struct {
	client *ent.Client
	redis  *redis.Client
	cfg    *config.Config
}

type DeliveryRegistry interface {
	NewDelivery() delivery.Delivery
}

func New(client *ent.Client, redis *redis.Client, cfg *config.Config) DeliveryRegistry {
	return &deliveryRegistry{
		client: client,
		redis:  redis,
		cfg:    cfg,
	}
}

func (d *deliveryRegistry) NewDelivery() delivery.Delivery {
	return delivery.Delivery{
		Songs: d.NewSongDelivery(),
	}
}
