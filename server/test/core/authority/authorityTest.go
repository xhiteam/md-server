package authority

import (
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/service/core"
	"github.com/mitchellh/go-testing-interface"
)

func TestGetLink(t *testing.T) {
	core.AuthorityServiceApp.CreateContextLink(request.CreateContextLinkReq{
		UID:        request.UID{UserId: 1},
		ContextId:  0,
		Permission: "",
	})
	core.AuthorityServiceApp.GetContextByLink(request.GetContextByLinkReq{
		UID:         request.UID{
			UserId: 1,
		},
		ContextLink: "",
	})
}