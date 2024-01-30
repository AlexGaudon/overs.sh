package settings

import "time"

const MAX_FILE_SIZE = 1024 * 1024 * 25

var (
	DeadlineTimeout = 30 * time.Second
	IdleTimeout     = 15 * time.Minute
)
