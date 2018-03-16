package red

import "github.com/mediocregopher/radix.v2/redis"

type Handler struct {
	db *redis.Client
}

func (h *Handler) Ping() (string, error) {
	res := h.db.Cmd("INCR", "ping:count")
	if res.Err != nil {
		return "", res.Err
	}
	return "pong", nil
}
