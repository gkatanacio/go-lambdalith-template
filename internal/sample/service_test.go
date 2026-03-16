package sample_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/go-lambdalith-template/internal/sample"
)

func Test_Service_Hello(t *testing.T) {
	// given
	service := sample.NewService(sample.Config{
		HelloWho: "John Doe",
	})

	// when
	msg := service.Hello(context.Background())

	// then
	assert.Equal(t, "Hello there, John Doe!", msg)
}
