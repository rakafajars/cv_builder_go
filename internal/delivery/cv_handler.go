package delivery

import (
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"cv-builder-api/pkg/middleware"
	"net/http"
)

type CVHandler struct {
	usecase usecase.CVUsecase
}

func NewCVHandler(u usecase.CVUsecase) *CVHandler {
	return &CVHandler{u}
}

func (h *CVHandler) GenerateCV(w http.ResponseWriter, r *http.Request) {

	ctxValue := r.Context().Value(middleware.UserIDKey)
	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses ditolak", "Gagal membaca User ID dari token")
		return
	}

	cvData, err := h.usecase.GetCVData(userID)

	if err != nil {
		pkg.InternalServerError(w, "Gagal mendapatkan data CV", err.Error())
		return
	}

	pkg.Success(w, "Berhasil men-generate CV utuh", cvData)
}
