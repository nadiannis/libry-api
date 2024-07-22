package handler

import (
	"net/http"

	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/domain/response"
	"github.com/nadiannis/libry-api/internal/usecase"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BookHandler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "books retrieved successfully",
		Data:    books,
	}

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *BookHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input request.BookRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Title != "", "title", "title is required")
	v.Check(input.Author != "", "author", "author is required")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
		return
	}

	book := h.usecase.Add(&input)

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "book added successfully",
		Data:    book,
	}

	err = utils.WriteJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
