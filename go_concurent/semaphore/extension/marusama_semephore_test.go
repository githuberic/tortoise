package extension

import (
	"context"
	"github.com/marusama/semaphore/v2"
	"testing"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()
	sem := semaphore.New(5) // new semaphore with limit = 5
	sem.Acquire(ctx, 1)     // acquire n with context
}
