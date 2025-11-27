package utils

import ("golang.org/x/crypto/bcrypt"
		"fmt")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// func CheckPassword(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func CheckPassword(password, hashed string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
//     return err == nil
// }

func CheckPassword(password, hashed string) bool {
    fmt.Println("DEBUG CHECK:", password, hashed)
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
    if err != nil {
        fmt.Println("COMPARE ERROR:", err)
        return false
    }
    return true
}
