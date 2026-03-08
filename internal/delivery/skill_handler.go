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

type SkillHandler struct {
	usecase usecase.SkillUsecase
}

func NewskillHandler(u usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{u}
}

func (h *SkillHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	skill, err := h.usecase.GetAllByUserID(userID)

	if err != nil {
		pkg.InternalServerError(w, "Internal Server Error", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Mendapatkan Skill", skill)

}

func (h *SkillHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)
	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal mebaca User ID dari token")
		return
	}

	var input struct{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	err := h.usecase.Create(&models.Skills{
		UserID: userID,
	})

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan Skill", err.Error())
		return
	}

	pkg.Created(w, "Berhasil membuat skill", nil)

}

func (h *SkillHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Skill", err.Error())
		return
	}

	ctxValue := r.Context().Value(middleware.UserIDKey)
	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal Membaca User ID dari token")
		return
	}

	var input struct{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	err = h.usecase.Update(userID, uint(id), &models.Skills{})

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan Skill", err.Error())
		return
	}

	pkg.Created(w, "Berhasil membuat skill", nil)

}

func (h *SkillHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)
	paramsID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramsID, 10, 32)

	if err != nil {
		pkg.BadRequest(w, "Gagal Mendapatkan ID Skill", err.Error())
		return
	}

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	err = h.usecase.Delete(userID, uint(id))

	if err != nil {
		pkg.BadRequest(w, "Gagal Menghapus Skill", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Menghapus Skill", nil)
}
