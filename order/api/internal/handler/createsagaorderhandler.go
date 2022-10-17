package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-dtm/order/api/internal/logic"
	"go-zero-dtm/order/api/internal/svc"
	"go-zero-dtm/order/api/internal/types"
)

func createSagaOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateSagaOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateSagaOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
