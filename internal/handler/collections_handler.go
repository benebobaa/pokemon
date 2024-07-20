package handler

import (
	"fmt"
	"pokemon_solid/internal/domain/request"
	"pokemon_solid/internal/usecase"
	"pokemon_solid/internal/util"

	"github.com/benebobaa/valo"
)

type CollectionsHandlerImpl struct {
	CollectionsUsecase usecase.CollectionsUsecase
}
type CollectionsHandler interface {
	Handler
	Catch(request request.CatchRequest)
	Release(request request.ReleaseRequest)
	FindAllByUserID(userId int)
}

func NewCollectionsHandler(collectionsUsecase usecase.CollectionsUsecase) CollectionsHandler {
	return CollectionsHandlerImpl{
		CollectionsUsecase: collectionsUsecase,
	}
}

func (c CollectionsHandlerImpl) Catch(request request.CatchRequest) {

	err := valo.Validate(request)

	if err != nil {
		fmt.Println("Error catch: ", err.Error())
	}

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

func (c CollectionsHandlerImpl) Release(request request.ReleaseRequest) {
	err := c.CollectionsUsecase.ReleaseCollection(request)

	if err != nil {
		fmt.Println("Release error: ", err.Error())
		return
	}

	fmt.Println("Success release collection id: ", request.CollectionID)
}

// FindAllByUserID implements CollectionsHandler.
func (c CollectionsHandlerImpl) FindAllByUserID(userId int) {
	collections := c.CollectionsUsecase.FindAllByUserID(userId)

	util.TableCollection(collections)
}
