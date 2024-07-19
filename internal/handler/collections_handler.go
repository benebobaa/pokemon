package handler

import (
	"fmt"
	"pokemon_solid/internal/domain/request"
	"pokemon_solid/internal/usecase"
	"pokemon_solid/internal/util"
)

type CollectionsHandlerImpl struct {
	CollectionsUsecase usecase.CollectionsUsecase
}

type CollectionsHandler interface {
	Catch(request request.CatchRequest)
	Release(id int)
	FindAll()
}

func NewCollectionsHandler(collectionsUsecase usecase.CollectionsUsecase) CollectionsHandler {
	return CollectionsHandlerImpl{
		CollectionsUsecase: collectionsUsecase,
	}
}

func (c CollectionsHandlerImpl) Catch(request request.CatchRequest) {
	message, err := c.CollectionsUsecase.TryCatch(request)
	if err != nil {
		fmt.Println("Catch error: ", err)
		return
	}

	fmt.Println("Catch success: ", message)
}

func (c CollectionsHandlerImpl) FindAll() {
	collections := c.CollectionsUsecase.FindAll()

	util.TableCollection(collections)
}

func (c CollectionsHandlerImpl) Release(id int) {
	panic("unimplemented")
}
