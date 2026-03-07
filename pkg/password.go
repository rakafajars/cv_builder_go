package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword() menerima:
	// 1. []byte(password): Password diubah jadi byte array.
	// 2. bcrypt.DefaultCost: Tingkat "kerumitan" hashing (default: 10).
	//    Semakin tinggi cost, semakin lama proses hashing, tapi semakin aman.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Ubah hasil hash (byte array) kembali ke string, lalu kembalikan.
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	// pengecekan password, jika password dan hash cocok maka akan mengembalikan nilai true
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
