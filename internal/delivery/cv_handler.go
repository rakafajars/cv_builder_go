package delivery

import (
	"cv-builder-api/internal/usecase"
	"cv-builder-api/pkg"
	"cv-builder-api/pkg/middleware"
	"html/template"
	"net/http"
	"path/filepath"
)

type CVHandler struct {
	usecase usecase.CVUsecase
}

func NewCVHandler(u usecase.CVUsecase) *CVHandler {
	return &CVHandler{u}
}

func (h *CVHandler) GenerateCV(w http.ResponseWriter, r *http.Request) {

	ctxValue := r.Context().Value(middleware.UserIDKey)

	var userID uint
	switch v := ctxValue.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	default:
		pkg.BadRequest(w, "Akses ditolak", "Gagal membaca User ID dari token")
		return
	}

	cvData, err := h.usecase.GetCVData(userID)

	if err != nil {
		pkg.InternalServerError(w, "Gagal mendapatkan data CV", err.Error())
		return
	}

	// Opsional: Jika user memanggil dengan query ?format=html, render template
	format := r.URL.Query().Get("format")
	if format == "html" {
		tmplPath := filepath.Join("templates", "cv_template.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			pkg.InternalServerError(w, "Gagal memuat template CV", err.Error())
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err = tmpl.Execute(w, cvData)
		if err != nil {
			// Jika gagal eksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	pkg.Success(w, "Berhasil men-generate CV utuh", cvData)
}
