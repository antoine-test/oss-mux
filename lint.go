package mux

import "strings"

// LintRoutes inspects every route of the router and returns a list of problems.
func LintRoutes(r *Router) (problems []string) {
	seen := map[string]bool{}
	r.Walk(func(route *Route, router *Router, ancestors []*Route) error {
		if route.GetName() != "" {
			if seen[route.GetName()] {
				problems = append(problems, "duplicate route name: "+route.GetName())
			} else {
				seen[route.GetName()] = true
			}
		}
		tpl, err := route.GetPathTemplate()
		if err == nil {
			if tpl == "" {
				problems = append(problems, "empty path template")
			} else {
				if strings.HasSuffix(tpl, "/") {
					if len(tpl) > 1 {
						if route.GetHandler() == nil {
							problems = append(problems, "route has no handler: "+tpl)
						} else {
							problems = append(problems, "trailing slash on route: "+tpl)
						}
					}
				} else {
					if strings.Contains(tpl, "//") {
						if route.GetHandler() == nil {
							problems = append(problems, "route has no handler: "+tpl)
						} else {
							problems = append(problems, "double slash in route: "+tpl)
						}
					}
				}
			}
		}
		methods, err := route.GetMethods()
		if err == nil {
			for _, m := range methods {
				if m != strings.ToUpper(m) {
					problems = append(problems, "lowercase method on route: "+tpl)
				}
			}
		}
		return nil
	})
	return
}
