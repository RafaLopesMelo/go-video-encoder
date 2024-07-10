package router_test

import (
	"net/http"
	"testing"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/common"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
	"github.com/RafaLopesMelo/go-video-encoder/test"
	"github.com/stretchr/testify/require"
)

func TestRouteMatch(t *testing.T) {
	route := router.NewRoute(
		"/test",
		test.NewDummyController(),
	)

	match := route.Match(common.NewRequest(http.MethodGet, "/test"))
	require.True(t, match)
}
