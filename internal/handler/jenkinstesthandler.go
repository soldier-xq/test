package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jenkins_test/internal/logic"
	"jenkins_test/internal/svc"
	"jenkins_test/internal/types"
)

func Jenkins_testHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJenkins_testLogic(r.Context(), svcCtx)
		resp, err := l.Jenkins_test(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
