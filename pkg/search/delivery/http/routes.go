package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	"google.golang.org/grpc"
)

func AddSearchRoute(r *gin.Engine, cc *grpc.ClientConn) {
	sc := search.NewSearchServiceClient(cc)
	handler := NewHandler(sc)

	r.GET("/search", handler.Search)
}
