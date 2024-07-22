package handler

import (
	"net/http"

	"github.com/nadiannis/libry-api/internal/domain/response"
	"github.com/nadiannis/libry-api/internal/usecase"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BorrowHandler struct {
	usecase usecase.BorrowUsecase
}

func NewBorrowHandler(usecase usecase.BorrowUsecase) BorrowHandler {
	return BorrowHandler{
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
