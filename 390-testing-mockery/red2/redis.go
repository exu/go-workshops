package red

import "github.com/mediocregopher/radix.v2/redis"

type storager interface {
	Cmd(string, ...interface{}) *redis.Resp
}
type Handler struct {
	db storager
}

func (h *Handler) Ping() (string, error) {
	res := h.db.Cmd("INCR", "ping:count")
	if res.Err != nil {
		return "", res.Err
	}
	return "pong", nil
}
