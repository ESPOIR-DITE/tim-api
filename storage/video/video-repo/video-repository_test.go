package video_repo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateVideoTable(t *testing.T) {
	result := CreateVideoTable()
	assert.True(t, result)
}
