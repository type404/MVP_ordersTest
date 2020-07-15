package controller

import (
	"log"
	"net/http"

	"api/model"

	"api/httputil"

	"github.com/gin-gonic/gin"
)

// ShowCUSTOMER godoc
// @Summary Show an CUSTOMER
// @Description get CUSTOMER by ID
// @Tags CUSTOMERs
// @Accept  json
// @Produce  json
// @Param id path string true "CUSTOMER identifier"
// @Success 200 {object} model.CUSTOMERResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /customers/{id} [get]
func (c *Controller) ShowCUSTOMER(ctx *gin.Context) {
	log.Println("Show CUSTOMER called...")
	id := ctx.Param("id")
	log.Println("Using CUSTOMER id:", id)
	customer, err := model.CUSTOMERGetByID(c.MongoDBCollCustomer, id)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

// ListCUSTOMERs godoc
// @Summary List CUSTOMERs
// @Description List all CUSTOMERs
// @Tags Infrastructure to Vehicle Information Messages (CUSTOMERs)
// @Accept  json
// @Produce  json
// @Param geohash query string false "List CUSTOMERs for a geohash"
// @Success 200 {object} model.CUSTOMERsResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /customers [get]
func (c *Controller) ListCUSTOMERs(ctx *gin.Context) {
	log.Println("Listing CUSTOMERs...")
	r, err := model.CUSTOMERsGetAll(c.MongoDBCollCustomer, c.MongoFindLimit)
	if err != nil {
		log.Println("Error listing CUSTOMERs, err:", err)
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, r)
}

// AddCUSTOMER godoc
// @Summary Add an CUSTOMER
// @Description add by json customer
// @Tags Infrastructure to Vehicle Information Messages (CUSTOMERs)
// @Accept  json
// @Produce  json
// @Param CUSTOMER body model.CUSTOMER true "Add CUSTOMER"
// @Success 200 {object} model.CUSTOMERResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /customers [post]
func (c *Controller) AddCUSTOMER(ctx *gin.Context) {
	log.Println("Adding CUSTOMER...")
	var customer model.CUSTOMER
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Println("CUSTOMER request body:", customer)
	r, err := model.CUSTOMERInsert(c.MongoDBCollCustomer, customer)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, r)
}

// DeleteCUSTOMER godoc
// @Summary Delete an CUSTOMER
// @Description Delete by CUSTOMER ID
// @Tags Infrastructure to Vehicle Information Messages (CUSTOMERs)
// @Accept  json
// @Produce  json
// @Param  id path string true "CUSTOMER ID" Format(int64)
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /customers/{id} [delete]
func (c *Controller) DeleteCUSTOMER(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("Deleting CUSTOMER with id:", id)
	if err := model.CUSTOMERDeleteByID(c.MongoDBCollCustomer, id); err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// DeleteCUSTOMERs godoc
// @Summary Delete all CUSTOMERs
// @Description Delete all CUSTOMERs
// @Tags Infrastructure to Vehicle Information Messages (CUSTOMERs)
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /customers [delete]
func (c *Controller) DeleteCUSTOMERs(ctx *gin.Context) {
	if err := model.CUSTOMERDeleteAll(c.MongoDBCollCustomer); err != nil {
		log.Println("Error deleting all CUSTOMERs, err:", err)
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
