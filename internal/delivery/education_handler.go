package delivery

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"cv-builder-api/pkg/middleware"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type EducationHandler struct {
	usecase usecase.EducationUsecase
}

func NewEducationHandler(u usecase.EducationUsecase) *EducationHandler {
	return &EducationHandler{u}
}

func (h *EducationHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)
	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	education, err := h.usecase.GetAllByUserID(userID)

	if err != nil {
		pkg.InternalServerError(w, "Internal Server Error", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Mendapatkan Education", education)
}

func (h *EducationHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input models.EducationRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	education := input.ToModel(userID)

	err := h.usecase.Create(&education)

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan Education", err.Error())
		return
	}

	pkg.Created(w, "Berhasil membuat Education", education)
}

func (h *EducationHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Work Experience", err.Error())
		return
	}

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input models.EducationRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	education := input.ToModel(userID)

	err = h.usecase.Update(userID, uint(id), &education)

	if err != nil {
		pkg.BadRequest(w, "Gagal Ubah Education", err.Error())
		return
	}

	response := input.ToResponse(uint(id))

	pkg.Success(w, "Berhasil Mengubah Education", response)
}

func (h *EducationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)

	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Education", err.Error())
		return
	}

	userID, ok := ctxValue.(uint)
	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal Membaca User ID dari token")
		return
	}

	err = h.usecase.Delete(userID, uint(id))

	if err != nil {
		pkg.BadRequest(w, "Gagal Menghapus Education", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Menghapus Education", nil)
}
