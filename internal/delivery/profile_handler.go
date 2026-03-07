package delivery

import (
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"cv-builder-api/pkg/middleware"
	"encoding/json"
	"net/http"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(u usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{u}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	profile, err := h.usecase.GetProfile(userID)

	if err != nil {
		pkg.NotFound(w, "Profile tidak di temukan", err.Error())
		return
	}

	pkg.Success(w, "Berhasil Mendapatkan Profile", profile)
}

func (h *ProfileHandler) UpsertProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	ctxValue := r.Context().Value(middleware.UserIDKey)

	userID, ok := ctxValue.(uint)

	if !ok {
		pkg.BadRequest(w, "Akses di tolak", "Gagal membaca User ID dari token")
		return
	}

	var input struct {
		FullName string `json:"full_name"`
		Phone    string `json:"phone"`
		Summary  string `json:"summary"`
		Address  string `json:"address"`
		PhotoUrl string `json:"photo_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	profile, err := h.usecase.UpsertProfile(
		int(userID),
		input.FullName,
		input.Phone,
		input.Address,
		input.PhotoUrl,
		input.Summary,
	)

	if err != nil {
		pkg.BadRequest(w, "Gagal Menyimpan profile", err.Error())
	}

	pkg.Created(w, "Berhasil membuat data atau mengubah data profile", profile)
}
