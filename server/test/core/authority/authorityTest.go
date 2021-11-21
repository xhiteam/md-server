package authority

import (
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/service/core"
	"testing"
)


func TestCreateContextLink(t *testing.T){
	err, contextLink := core.AuthorityServiceApp.CreateContextLink(request.CreateContextLinkReq{
		UID:        request.UID{UserId: 1},
		ContextId:  4,
		Permission: "read",
	})
	if err != nil {
		t.Error(err)
	}
    t.Log(contextLink)
}
func TestGetContextByLink(t *testing.T){
	err,contextinfo:=core.AuthorityServiceApp.GetContextByLink(request.GetContextByLinkReq{
		UID:         request.UID{UserId: 2},
		ContextLink: "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(contextinfo)
}


