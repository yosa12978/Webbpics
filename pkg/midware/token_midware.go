package midware

import (
	"net/http"
	"strings"

	"github.com/yosa12978/webbpics/pkg/helpers"
	"github.com/yosa12978/webbpics/pkg/services"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token_raw := r.Header.Get("Authorization")
		if !strings.HasPrefix(token_raw, "Bearer ") {
			helpers.WriteJson(w, 401, helpers.StatusMessage{Status_code: 401, Detail: "Unauthorized"})
			return
		}
		token := strings.Replace(token_raw, "Bearer ", "", 1)
		user, err := services.NewUserService().GetUserByToken(token)
		if err != nil {
			helpers.WriteJson(w, 401, helpers.StatusMessage{Status_code: 401, Detail: "Unauthorized"})
			return
		}
		if !user.Is_admin {
			helpers.WriteJson(w, 403, helpers.StatusMessage{Status_code: 403, Detail: "Forbidden"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
