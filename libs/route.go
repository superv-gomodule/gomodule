package libs

func (c *Controller) GET(route Route) Route {
	route.Method = GET
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) POST(route Route) Route {
	route.Method = POST
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) PUT(route Route) Route {
	route.Method = PUT
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) DELETE(route Route) Route {
	route.Method = DELETE
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) PATCH(route Route) Route {
	route.Method = PATCH
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) OPTIONS(route Route) Route {
	route.Method = OPTIONS
	c.routes = append(c.routes, route)
	return route
}

func (c *Controller) HEAD(route Route) Route {
	route.Method = HEAD
	c.routes = append(c.routes, route)
	return route
}
