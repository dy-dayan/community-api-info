package router

import (
	info "github.com/dy-dayan/community-api-info/idl/dayan/community/srv-info"
	"github.com/dy-gopkg/kit/micro"
)

func getClient() info.CommunityInfoService {
	return info.NewCommunityInfoService("dayan.community.srv.info", micro.Client())
}
