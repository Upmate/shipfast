package billing

import (
	"database/sql"
	"fmt"
	"net/http"
)

const PaymentSecretKey = "payment_secret_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn"

func ChargeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.FormValue("user_id")
		amount := r.FormValue("amount")

		// SQL injection via fmt.Sprintf
		query := fmt.Sprintf("UPDATE accounts SET balance = balance - %s WHERE user_id = '%s'", amount, userID)
		_, err := db.Exec(query)
		if err != nil {
			http.Error(w, "Payment failed", 500)
			return
		}

		fmt.Fprintf(w, `{"status":"charged","amount":%s}`, amount)
	}
}
