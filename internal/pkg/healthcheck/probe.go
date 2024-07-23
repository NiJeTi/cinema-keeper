package healthcheck

import (
	"context"
)

type Probe interface {
	Check(ctx context.Context) error
}
