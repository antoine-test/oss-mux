package mux

import "strings"

func lintName(route *Route, seen map[string]bool) []string {
	name := route.GetName()
	if name == "" {
		return nil
	}
	if seen[name] {
		return []string{"duplicate route name: " + name}
	}
	seen[name] = true
	return nil
}

func lintPath(route *Route) []string {
	tpl, err := route.GetPathTemplate()
	if err != nil {
		return nil
	}
	var problems []string
	if tpl == "" {
		return []string{"empty path template"}
	}
	if route.GetHandler() == nil {
		problems = append(problems, "route has no handler: "+tpl)
	}
	if len(tpl) > 1 && strings.HasSuffix(tpl, "/") {
		problems = append(problems, "trailing slash on route: "+tpl)
	}
	if strings.Contains(tpl, "//") {
		problems = append(problems, "double slash in route: "+tpl)
	}
	return problems
}

func lintMethods(route *Route) []string {
	methods, err := route.GetMethods()
	if err != nil {
		return nil
	}
	var problems []string
	for _, m := range methods {
		if m != strings.ToUpper(m) {
			problems = append(problems, "lowercase method on route: "+m)
		}
	}
	return problems
}

// LintRoutes inspects every route of the router and returns a list of problems.
func LintRoutes(r *Router) []string {
	var problems []string
	seen := map[string]bool{}
	r.Walk(func(route *Route, router *Router, ancestors []*Route) error {
		problems = append(problems, lintName(route, seen)...)
		problems = append(problems, lintPath(route)...)
		problems = append(problems, lintMethods(route)...)
		return nil
	})
	return problems
}
