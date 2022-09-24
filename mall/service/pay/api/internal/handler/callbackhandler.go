package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero_microservices/mall/service/pay/api/internal/logic"
	"go-zero_microservices/mall/service/pay/api/internal/svc"
	"go-zero_microservices/mall/service/pay/api/internal/types"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbakRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
