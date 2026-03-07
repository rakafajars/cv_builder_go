package delivery

import (
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	user, err := h.usecase.Register(input.Email, input.Password)
	if err != nil {
		pkg.BadRequest(w, "Registrasi Gagal!", err.Error())
		return
	}

	responseData := map[string]any{
		"id":    user.ID,
		"email": user.Email,
	}

	pkg.Created(w, "Registrasi Berhasil", responseData)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.BadRequest(w, "Gagal Parsing Data", err.Error())
		return
	}

	if input.Email == "" || input.Password == "" {
		pkg.BadRequest(w, "Email dan password wajib diisi", "email or password cannot be empty")
		return
	}

	token, err := h.usecase.Login(input.Email, input.Password)
	if err != nil {
		pkg.Unauthorized(w, "Login Gagal", err.Error())
		return
	}

	pkg.Success(w, "Login Berhasil", map[string]any{
		"token": token,
	})
}
