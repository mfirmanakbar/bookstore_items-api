package controllers

import (
	"encoding/json"
	"github.com/mfirmanakbar/bookstore_items-api/src/domain/items"
	"github.com/mfirmanakbar/bookstore_items-api/src/services"
	"github.com/mfirmanakbar/bookstore_items-api/src/utils/http_utils"
	"github.com/mfirmanakbar/bookstore_oauth-go/oauth"
	"github.com/mfirmanakbar/bookstore_utils-go/rest_errors"
	"io/ioutil"
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
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
