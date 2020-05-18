package register

import (
	fw "github.com/agorago/wego/fw"
	apiregister "github.com/agorago/storeapi/register"
)



func RegisterStoreService(rs fw.RegistrationService,
		serviceToInvoke interface{},
		middlewares []fw.Middleware,
		proxyMiddlewares []fw.Middleware) {
	var sd = apiregister.GetServiceDescriptor(proxyMiddlewares)
	sd.ServiceToInvoke = serviceToInvoke
	rs.RegisterService("storeService", sd)
}
