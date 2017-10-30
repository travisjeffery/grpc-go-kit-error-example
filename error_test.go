package example

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	t.Run("WrapErr", func(tt *testing.T) {
		e := WrapErr(errors.New("is it borked?"), "it is borked")
		require.Equal(tt, "it is borked: is it borked?", e.Error())
	})

	t.Run("E", func(tt *testing.T) {
		e := E("it is borked", errors.New("is it borked?"), 404)
		require.Equal(tt, "it is borked: is it borked?", e.Error())
		require.Equal(tt, int32(404), e.(*Error).Code)
	})
}
