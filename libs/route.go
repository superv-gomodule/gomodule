package libs

func (c *Controller) GET(route Route) Route {
	route.Method = GET
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) POST(route Route) Route {
	route.Method = POST
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) PUT(route Route) Route {
	route.Method = PUT
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) DELETE(route Route) Route {
	route.Method = DELETE
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) PATCH(route Route) Route {
	route.Method = PATCH
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) OPTIONS(route Route) Route {
	route.Method = OPTIONS
	c.Routes = append(c.Routes, route)
	return route
}

func (c *Controller) HEAD(route Route) Route {
	route.Method = HEAD
	c.Routes = append(c.Routes, route)
	return route
}
