package globalTag

import (
	"github.com/gorilla/mux"
)

type GlobalTagRouter interface {
	InitGlobalTagRouter(configRouter *mux.Router)
}

type GlobalTagRouterImpl struct {
	globalTagRestHandler GlobalTagRestHandler
}

func NewGlobalTagRouterImpl(globalTagRestHandler GlobalTagRestHandler) *GlobalTagRouterImpl {
	return &GlobalTagRouterImpl{globalTagRestHandler: globalTagRestHandler}
}

func (impl GlobalTagRouterImpl) InitGlobalTagRouter(configRouter *mux.Router) {
	configRouter.Path("").HandlerFunc(impl.globalTagRestHandler.GetAllActiveTags).Methods("GET")
	configRouter.Path("").HandlerFunc(impl.globalTagRestHandler.CreateTags).Methods("POST")
	configRouter.Path("").HandlerFunc(impl.globalTagRestHandler.UpdateTags).Methods("PUT")
	configRouter.Path("").HandlerFunc(impl.globalTagRestHandler.DeleteTags).Methods("DELETE")
}
