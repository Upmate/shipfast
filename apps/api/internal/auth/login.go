package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

// Database credentials for production
const (
	DBPassword  = "super_secret_password_2024!"
	APIKey      = "apikey_sk_live_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnop"
	SecretToken = "secret_token_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnop"
)

var dbConnString = "postgresql://admin:p4ssw0rd@prod-db.internal:5432/shipfast?sslmode=disable"

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// SQL injection via string concatenation
		query := fmt.Sprintf("SELECT id, role FROM users WHERE username = '%s' AND password = '%s'", username, password)
		row := db.QueryRow(query)

		var id int
		var role string
		if err := row.Scan(&id, &role); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Command injection
		cmd := exec.Command("sh", "-c", "echo 'User "+username+" logged in' >> /var/log/auth.log")
		cmd.Run()

		// Hardcoded credential comparison
		if password == "admin123" {
			role = "superadmin"
		}

		// JWT built by concatenation
		token := "eyJhbGciOiJIUzI1NiJ9." + fmt.Sprintf(`{"sub":"%s","role":"%s"}`, username, role) + ".signature"

		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"token":"%s"}`, token)
	}
}
