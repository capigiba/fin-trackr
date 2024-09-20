package user

import (
	"fintrack/internal/domain/entity"
	"fintrack/internal/pkg/json"
	"fintrack/internal/pkg/reason"
	"fintrack/internal/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for managing users
type UserHandler struct {
	service *user.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user into the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entity.User true "User registration details"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.RegisterUser(user.UserName, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}
	json.WriteJSON(c, http.StatusOK, reason.UserRegistered.Message(), nil)
}

// Login godoc
// @Summary User login
// @Description Allows a user to log in and obtain a token
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body object true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful, returns token"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := json.ReadJSON(c, &req); err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(req.UserName, req.Password)
	if err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return
	}

	json.WriteJSON(c, http.StatusOK, reason.UserLogged.Message(), gin.H{"token": token})
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Fetches details of a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User "User details"
// @Failure 400 {object} map[string]string "Invalid User ID"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		json.ErrorJSON(c, reason.InvalidUserID.Message(), http.StatusBadRequest)
		return
	}

	json.WriteJSON(c, http.StatusOK, reason.UserRetrieved.Message(), user)
}

// UpdatedUser godoc
// @Summary Update user information
// @Description Updates the information of an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entity.User true "Updated user details"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {object} map[string]string "Invalid User ID or input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdatedUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		json.ErrorJSON(c, reason.InvalidUserID.Message(), http.StatusBadRequest)
		return
	}

	var user entity.User
	if err := json.ReadJSON(c, &user); err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateUser(id, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Status)
	if err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	json.WriteJSON(c, http.StatusOK, reason.UserUpdated.Message(), nil)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Deletes a user from the system by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} map[string]string "Invalid User ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		json.ErrorJSON(c, reason.InvalidUserID.Message(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		json.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	json.WriteJSON(c, http.StatusOK, reason.UserDeleted.Message(), nil)
}
