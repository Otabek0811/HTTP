package handler

import (
	"app/config"
	"app/models"
	"app/pkg/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handler) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateProduct(w, r)

	case "GET":

		var (
			method = r.URL.Query().Get("method")
		)

		if method == "GET_LIST" {
			h.GetListProduct(w, r)

		} else if method == "GET" {
			h.GetProductByID(w, r)
		}
	case "PUT":
		h.UpdateProduct(w, r)
	case "DELETE":
		h.DeleteProduct(w, r)
	}

}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.CreateProduct

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Product().CreateProduct(&product)
	if err != nil {
		h.handlerResponse(w, "error while storage Product create:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	resp, err := h.strg.Product().GetProductByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage Product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "Success", http.StatusOK, resp)

}

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while give product id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	resp, err := h.strg.Product().GetProductByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)

}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var upPorduct models.UpdateProduct

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &upPorduct)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Product().UpdateProduct(&upPorduct)
	if err != nil {
		h.handlerResponse(w, "error while storage product upadate:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	resp, err := h.strg.Product().GetProductByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage Product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)

}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")
	fmt.Println(id)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while delete --> give product id: invalid uuid", http.StatusBadRequest, nil)
		return
	}
	_, err := h.strg.Product().GetProductByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage Product get by id:"+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = h.strg.Product().DeleteProduct(&models.ProductPrimaryKey{Id: id})

	if err != nil {
		h.handlerResponse(w, "error while delete product:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "Success", http.StatusOK, nil)

}

func (h *handler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	var (
		offsetStr       = r.URL.Query().Get("offset")
		limitStr        = r.URL.Query().Get("limit")
		search          = r.URL.Query().Get("search")
		offset    int   = config.Load().DefaultOffset
		limit     int   = config.Load().DefaultLimit
		err       error = nil
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			h.handlerResponse(w, "error while offset: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			h.handlerResponse(w, "error while limit: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
	}

	resp, err := h.strg.Product().GetListProduct(&models.ProductGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		h.handlerResponse(w, "error while storage product get list:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}
