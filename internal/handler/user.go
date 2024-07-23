package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/domain/response"
	"github.com/nadiannis/libry-api/internal/usecase"
	"github.com/nadiannis/libry-api/internal/utils"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "users retrieved successfully",
		Data:    users,
	}

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input request.UserRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Username != "", "username", "username is required")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := h.usecase.Add(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUserAlreadyExists):
			utils.BadRequestResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "user added successfully",
		Data:    user,
	}

	err = utils.WriteJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
