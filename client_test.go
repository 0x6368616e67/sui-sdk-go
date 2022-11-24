package sui

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTransactionNumber(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	num, err := cli.GetTotalTransactionNumber(context.Background())
	assert.Equal(t, err, nil)
	t.Log(num)
	assert.Greater(t, num, uint64(0))
}
