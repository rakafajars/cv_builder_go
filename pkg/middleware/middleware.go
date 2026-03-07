package middleware

import (
	"context"
	"cv-builder-api/pkg"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Langkah 1: Ambil Header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				pkg.Unauthorized(w, "Akses Ditolak", "Header Authorization tidak ditemukan")
				return
			}

			// Langkah 2: Pisahkan kata "Bearer" dan tokennya
			// Contoh: "Bearer eyJhbGci..." menjadi ["Bearer", "eyJhbGci..."]
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {

				pkg.Unauthorized(w, "Akses di tolak", "Format Token Harus Bearer <token>")
				return
			}

			tokenString := parts[1]

			claims, err := pkg.ValidateToken(tokenString, secretKey)
			if err != nil {
				pkg.Unauthorized(w, "Akses Ditolak", "Token tidak valid atau sudah kedaluwarsa")
				return
			}

			// Langkah 4: Keajaiban Golang Context
			// Kita masukkan UserID dari dalam token ke dalam "keranjang" Request
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)

			// Buat request baru yang sudah membawa keranjang (context) tersebut
			reqWithCtx := r.WithContext(ctx)

			// Langkah 5: Persilakan masuk ke Handler selanjutnya
			next.ServeHTTP(w, reqWithCtx)
		})
	}
}
