package audit

import "time"

type Audit struct {
	CreatedAt  time.Time
	ModifiedAt *time.Time
}
