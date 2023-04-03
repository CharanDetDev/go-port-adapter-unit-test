package model

type MakeNewRedisCache struct {
	Key    string      `json:"key" validate:"required"`
	Data   interface{} `json:"data" validate:"required"`
	Expire string      `json:"expire"`
}
