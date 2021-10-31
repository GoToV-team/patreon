package base_handler

import (
	"net/http"
	"patreon/internal/app"
	"patreon/internal/app/middleware"
	"patreon/internal/app/utilits"
	"strings"

	"github.com/gorilla/mux"

	hf "patreon/internal/app/delivery/http/handlers/base_handler/handler_interfaces"

	"github.com/sirupsen/logrus"
)

const (
	GET     = http.MethodGet
	POST    = http.MethodPost
	PUT     = http.MethodPut
	DELETE  = http.MethodDelete
	OPTIONS = http.MethodOptions
)

type queriesMethods struct {
	method     string
	methodFunc hf.HandlerFunc
	MethodInfo
}

type BaseHandler struct {
	utilitiesMiddleware middleware.UtilitiesMiddleware
	corsMiddleware      middleware.CorsMiddleware
	handlerMethods      map[string]hf.HandlerFunc
	middlewares         []hf.HMiddlewareFunc
	methodsWithQueries  []queriesMethods
	HelpHandlers
}

func NewBaseHandler(log *logrus.Logger, router *mux.Router, config *app.CorsConfig) *BaseHandler {
	h := &BaseHandler{handlerMethods: map[string]hf.HandlerFunc{}, middlewares: []hf.HMiddlewareFunc{},
		utilitiesMiddleware: middleware.NewUtilitiesMiddleware(log),
		corsMiddleware:      middleware.NewCorsMiddleware(config, router),
		HelpHandlers: HelpHandlers{
			Responder: utilits.Responder{LogObject: utilits.NewLogObject(log)},
		},
	}
	h.AddMiddleware(h.corsMiddleware.SetCors)
	h.AddMiddleware(h.utilitiesMiddleware.UpgradeLogger, h.utilitiesMiddleware.CheckPanic)
	return h
}

func (h *BaseHandler) AddMiddleware(middleware ...hf.HMiddlewareFunc) {
	h.middlewares = append(h.middlewares, middleware...)
}

func (h *BaseHandler) AddMethodWithQueries(method string, handlerMethod hf.HandlerFunc,
	middlewares ...hf.HFMiddlewareFunc) *MethodInfo {
	h.methodsWithQueries = append(h.methodsWithQueries, queriesMethods{
		methodFunc: h.applyHFMiddleware(handlerMethod, middlewares...),
		method:     method,
		MethodInfo: MethodInfo{},
	})
	return &h.methodsWithQueries[len(h.methodsWithQueries)-1].MethodInfo
}

func (h *BaseHandler) AddMethod(method string, handlerMethod hf.HandlerFunc, middlewares ...hf.HFMiddlewareFunc) {
	h.handlerMethods[method] = h.applyHFMiddleware(handlerMethod, middlewares...)
}

func (h *BaseHandler) applyQueriesMethods(route *mux.Route) {
	for _, method := range h.methodsWithQueries {
		route.Path("").Queries(method.queries...).HandlerFunc(method.methodFunc).Methods(method.method)
	}
}

func (h *BaseHandler) applyHFMiddleware(handler hf.HandlerFunc,
	middlewares ...hf.HFMiddlewareFunc) hf.HandlerFunc {
	resultHandler := handler
	for index := len(middlewares) - 1; index >= 0; index-- {
		resultHandler = middlewares[index](resultHandler)
	}
	return resultHandler
}

func (h *BaseHandler) applyMiddleware(handler http.Handler) http.Handler {
	resultHandler := handler
	for index := len(h.middlewares) - 1; index >= 0; index-- {
		resultHandler = h.middlewares[index](resultHandler)
	}
	return resultHandler
}

func (h *BaseHandler) getListMethods() []string {
	var useMethods []string
	for key := range h.handlerMethods {
		useMethods = append(useMethods, key)
	}
	useMethods = append(useMethods, http.MethodOptions)
	return useMethods
}

func (h *BaseHandler) Connect(route *mux.Route) {
	route.Handler(h.applyMiddleware(h)).Methods(h.getListMethods()...)
	h.applyQueriesMethods(route)
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.PrintRequest(r)
	ok := true
	var handler hf.HandlerFunc

	handler, ok = h.handlerMethods[r.Method]
	if ok {
		handler(w, r)
	} else {
		h.Log(r).Errorf("Unexpected http method: %s", r.Method)
		w.Header().Set("Allow", strings.Join(h.getListMethods(), ", "))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
