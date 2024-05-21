package handler

import (
	"errors"
	_ "exam/api/docs"
	"exam/api/models"
	"exam/pkg/check"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/Customers [post]
// @Summary		Creates a Customers
// @Description	This api creates a Customers and returns its id
// @Tags		Customers
// @Accept		json
// @Produce		json
// @Param		Customers body models.Customers true "Customers"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomers(c *gin.Context) {
	Customers := models.Customers{}

	if err := c.ShouldBindJSON(&Customers); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateAge(Customers.Age,Customers.Birthday); err != nil {
		handleResponse(c, h.Log, "error while validating Customers age, year: "+strconv.Itoa(Customers.Age), http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidateGmail(Customers.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	for _, phone := range Customers.Phone {
		if !check.ValidatePhone(phone) {
			handleResponse(c, h.Log, "error while validation phone", http.StatusBadRequest, errors.New("wrong phone"))
		return
		}
	}

	id, err := h.Service.Customers().Create(c.Request.Context(), Customers)
	if err != nil {
		handleResponse(c, h.Log, "error while creating Customers", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/Customers/updateCustomers/{id} [PUT]
// @Summary		Update a Customers
// @Description	This API updates a Customers
// @Tags		Customers
// @Accept		json
// @Produce		json
// @Param		id path string true "Customers ID"
// @Param		Customers body models.Customers true "Customers object to update"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) UpdateCustomers(c *gin.Context) {

	id := c.Param("id")

	Customers := models.Customers{}

	if err := c.ShouldBindJSON(&Customers); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateAge(Customers.Age,Customers.Birthday); err != nil {
		handleResponse(c, h.Log, "error while validating Customers age, year: "+strconv.Itoa(Customers.Age), http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidateGmail(Customers.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	for _, phone := range Customers.Phone {
		if !check.ValidatePhone(phone) {
			handleResponse(c, h.Log, "error while validation phone", http.StatusBadRequest, errors.New("wrong phone"))
		return
		}
	}

	err := h.Service.Customers().UpdateCustomers(c.Request.Context(), id,Customers)
	if err != nil {
		handleResponse(c, h.Log, "error while updating Customers", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, nil)
}

// @Router		/Customers [get]
// @Summary		Get a Customers
// @Description	This API returns all Customerss
// @Tags		Customers
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCustomerss(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.Customers().GetAllCustomers(c.Request.Context(), models.GetAllCustomersRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all Customerss", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "Select all successful", http.StatusOK, resp)
}

// @Router		/Customers/{id} [get]
// @Summary		Get by id a Customers
// @Description	This API get by id a Customers
// @Tags		Customers
// @Produce		json
// @Param		id path string true "Customers Id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Service.Customers().GetByIdCustomers(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while get by id Customers %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Select by id successful", http.StatusOK, data)
}

// @Router		/Customers/deleteCustomers/{id} [delete]
// @Summary		Delete a Customers
// @Description	This API delete a Customers
// @Tags		Customers
// @Produce		json
// @Param		id path string true "Customers ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteCustomers(c *gin.Context) {

	Id := c.Param("id")

	err := h.Service.Customers().DeleteCustomers(c.Request.Context(), Id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting  Customers", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Customers deleted successfully", http.StatusOK, err)
}