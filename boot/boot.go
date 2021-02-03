package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	//v := g.View()
	//c := g.Config()
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetFileServerEnabled(false)

	s.SetPort(8081)
}
