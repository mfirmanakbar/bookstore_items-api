package controllers

import (
	"fmt"
	"github.com/mfirmanakbar/bookstore_items-api/src/domain/items"
	"github.com/mfirmanakbar/bookstore_items-api/src/services"
	"github.com/mfirmanakbar/bookstore_oauth-go/oauth"
	"net/http"
)

/*
in Golang we can't create same function name at same directory
even we create on different files like items_controller and users_controller has same func like 'func create()'
so the solution is using interface
*/

var (
	// 4. this the variable should we defined if we are using interface way
	// the purpose is when we need to call all controller functions here,
	// we can use like `ItemsController.Create()`
	// we can use like `usersController.Create()`
	ItemsController itemsControllerInterface = &itemsController{}
)

// 1. this is the interface
type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// 2. this is struct that implement the interface
type itemsController struct{}

// 3. this is out function that implement the struct
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// /*
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return json error
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: return json error
		return
	}

	//TODO 201 created
	fmt.Println(result)

	// */
	//panic("implement me")
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
