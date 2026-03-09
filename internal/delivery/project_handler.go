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

type ProjectHandler struct {
	usecase usecase.ProjectUsecase
}

func NewProjectHandler(u usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{u}
}

func (h *ProjectHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	workExperience, err := h.usecase.GetAllByUserID(userID)

	if err != nil {
		pkg.InternalServerError(w, "Internal Server Error", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Mendapatkan Project", workExperience)
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input models.ProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	project := input.ToModel(userID)

	err := h.usecase.Create(&project)

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan Project", err.Error())
		return
	}

	pkg.Created(w, "Berhasil membuat data Project", project)
}

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Project", err.Error())
		return
	}

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input models.ProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	project := input.ToModel(userID)

	err = h.usecase.Update(userID, uint(id), &project)

	if err != nil {
		pkg.BadRequest(w, "Gagal Ubah Project", err.Error())
		return
	}

	response := input.ToResponse(uint(id))

	pkg.Success(w, "Berhasil mengubah data Project", response)
}

func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)

	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Project", err.Error())
		return
	}

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	err = h.usecase.Delete(userID, uint(id))

	if err != nil {
		pkg.BadRequest(w, "Gagal Menghapus Project", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Menghapus Project", nil)
}
