package auth

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const CONTEXT_KEY_USER = "key_user"
const CONTEXT_KEY_ORG_ID = "org_id"

func AuthUser(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgIdStr := mux.Vars(r)["oid"]
		if len(orgIdStr) == 0 {
			http.Error(w, "param org_id not found", http.StatusBadRequest)
			return
		}

		orgId, err := strconv.ParseInt(orgIdStr, 10, 64)
		if err != nil {
			http.Error(w, "param org_id invalid", http.StatusBadRequest)
			return
		}

		userId, _ := Me(r)

		is, err := IsMember(userId, orgId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !is {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), CONTEXT_KEY_USER, userId))
		r = r.WithContext(context.WithValue(r.Context(), CONTEXT_KEY_ORG_ID, orgId))
		handler(w, r)
	}
}

func IsMember(myid, orgid int64) (bool, error) {
	var res IsMemberResponse
	if err := CallUIC("Authority.IsMember", IsMemberRequest{
		UserID: myid,
		OrgID:  orgid,
	}, &res); err != nil {
		return false, err
	}

	return res.Is, nil
}
