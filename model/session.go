package model

const (
	SessionBucket = "Session"
	SessionsKey   = "sessions"
)

type Session struct {
	Name      string `json:"name" validate:"required"`
	Time      uint32 `json:"time" validate:"required,gte=0"`
	CreatedAt int64  `json:"created_at"`
}

type Sessions []Session
