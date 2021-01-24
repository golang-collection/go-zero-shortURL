package handler

import (
	"net/http"

	"go-zero-shortURL/api/internal/logic"
	"go-zero-shortURL/api/internal/svc"
	"go-zero-shortURL/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ExpandHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), ctx)
		resp, err := l.Expand(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
