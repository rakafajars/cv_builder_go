package delivery

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"cv-builder-api/pkg/middleware"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type WorkExperienceHandler struct {
	usecase usecase.WorkExperienceUsecase
}

func NewWorkExperienceHandler(u usecase.WorkExperienceUsecase) *WorkExperienceHandler {
	return &WorkExperienceHandler{u}
}

func (h *WorkExperienceHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
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

	pkg.Success(w, "Berhasil Mendapatkan Work Experience", workExperience)
}

func (h *WorkExperienceHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input struct {
		CompanyName string     `json:"company_name"`
		Position    string     `json:"position"`
		StartDate   time.Time  `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
		IsCurrent   bool       `json:"is_current"`
		Description string     `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	workExperience := models.WorkExperience{
		UserID:      userID,
		Position:    input.Position,
		CompanyName: input.CompanyName,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		IsCurrent:   input.IsCurrent,
		Description: input.Description,
	}

	err := h.usecase.Create(
		&workExperience,
	)

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan Work Experience", err.Error())
		return
	}

	pkg.Created(w, "Berhasil membuat data Work Experience", workExperience)
}

func (h *WorkExperienceHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		CompanyName string     `json:"company_name"`
		Position    string     `json:"position"`
		StartDate   time.Time  `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
		IsCurrent   bool       `json:"is_current"`
		Description string     `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	workExperience := models.WorkExperience{
		UserID:      userID,
		Position:    input.Position,
		CompanyName: input.CompanyName,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		IsCurrent:   input.IsCurrent,
		Description: input.Description,
	}

	err = h.usecase.Update(
		userID, uint(id),
		&workExperience,
	)

	if err != nil {
		pkg.BadRequest(w, "Gagal Ubah Work Experience", err.Error())
		return
	}

	pkg.Success(w, "Berhasil mengubah data Work Experience", workExperience)
}

func (h *WorkExperienceHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.usecase.Delete(userID, uint(id))

	if err != nil {
		pkg.BadRequest(w, "Gagal Menghapus Work Experience", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Menghapus Work Experience", nil)

}
