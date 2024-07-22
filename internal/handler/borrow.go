package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/domain/response"
	"github.com/nadiannis/libry-api/internal/usecase"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BorrowHandler struct {
	usecase usecase.IBorrowUsecase
}

func NewBorrowHandler(usecase usecase.IBorrowUsecase) IBorrowHandler {
	return &BorrowHandler{
		usecase: usecase,
	}
}

func (h *BorrowHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	borrows := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "borrowed books retrieved successfully",
		Data:    borrows,
	}

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *BorrowHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	var input request.BorrowRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.UserID != "", "user_id", "user_id is required")
	v.Check(input.BookID != "", "book_id", "book_id is required")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
		return
	}

	borrowedBook, err := h.usecase.Borrow(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUserNotFound) || errors.Is(err, utils.ErrBookNotFound):
			utils.NotFoundResponse(w, r, err)
		case errors.Is(err, utils.ErrBookCurrentlyBorrowed):
			utils.BadRequestResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "book is borrowed successfully",
		Data:    borrowedBook,
	}

	err = utils.WriteJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
