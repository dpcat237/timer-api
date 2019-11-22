package model

const (
	SessionBucket = "Session"
	SessionsKey   = "sessions"
)

type Session struct {
	Name      string `json:"name"`
	Time      uint32 `json:"time"`
	CreatedAt uint32 `json:"created_at"`
}

type Sessions []Session
