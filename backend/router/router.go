package router

import (
	"context"
	"net/http"
	"regexp"
	"sync"

	"golang.org/x/time/rate"
)

type route struct {
	regex   *regexp.Regexp
	handler http.HandlerFunc
	methods map[string]bool
}

type Router struct {
	routes      []route
	limiter     *rate.Limiter
	limiterLock sync.Mutex
	indexPath   string
	apiPrefix   string
}

func (router *Router) SetRateLimiter(rps float64, burst int) {
	router.limiterLock.Lock()
	defer router.limiterLock.Unlock()
	router.limiter = rate.NewLimiter(rate.Limit(rps), burst)
}

func (router *Router) NewRoute(regexpString string, handler http.HandlerFunc, methods ...string) {
	if router.apiPrefix != "" {
		regexpString = "/" + router.apiPrefix + regexpString
	}
	regex := regexp.MustCompile("^" + regexpString + "$")
	methodMap := make(map[string]bool)
	for _, method := range methods {
		methodMap[method] = true
	}
	router.routes = append(router.routes, route{
		regex,
		handler,
		methodMap,
	})
}

func (router *Router) SetIndexPath(path string) {
	router.indexPath = path
}

func (router *Router) SetAPIPrefix(prefix string) {
	router.apiPrefix = prefix
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.limiterLock.Lock()
	limiter := router.limiter
	router.limiterLock.Unlock()

	if limiter != nil && !limiter.Allow() {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	if !regexp.MustCompile("^/" + router.apiPrefix).MatchString(r.URL.Path) {
		http.ServeFile(w, r, router.indexPath)
		return
	}

	for _, v := range router.routes {
		matches := v.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if !v.methods[r.Method] {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			matchMap := make(map[string]string)
			groupNames := v.regex.SubexpNames()
			for i, name := range groupNames {
				if i != 0 && name != "" {
					matchMap[name] = matches[i]
				}
			}
			ctx := context.WithValue(r.Context(), "routeParams", matchMap)
			v.handler(w, r.WithContext(ctx))
			return
		}
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func StartServer(addr string, certFile string, keyFile string, handler http.Handler) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}
