package main

// func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
// 	customers := []Customer{}

// 	h.DB.Find(&customers)

// 	return c.JSON(http.StatusOK, customers)
// }

// func (h *CustomerHandler) GetCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) SaveCustomer(c echo.Context) error {
// 	customer := Customer{}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := h.DB.Delete(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }
