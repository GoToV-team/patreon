package handlers

import (
	"patreon/internal/app/handlers/base_handler"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type MainHandler struct {
	router *mux.Router
	base_handler.BaseHandler
}

func NewMainHandler() *MainHandler {
	return &MainHandler{
		BaseHandler: *base_handler.NewBaseHandler(logrus.New()),
	}
}

//func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	h.PrintRequest(r)
//	h.router.ServeHTTP(w, r)
//}
//func (h *MainHandler) JoinHandlers(joinedHandlers []app.Joinable) {
//	h.baseHandler.AddHandlers(joinedHandlers)
//	h.baseHandler.Join(h.router)
//}
//func (h *MainHandler) SetRouter(router *mux.Router) {
//	h.router = router
//}
//func (h *MainHandler) SetLogger(logger *logrus.Logger) {
//	h.log = logger
//}
//func (h *MainHandler) Join(router *mux.Router) {
//	h.baseHandler.Join(router)
//}
