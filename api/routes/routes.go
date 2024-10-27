package routes

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func InitRoutes(db gorm.DB) {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test routes")
	})

	UserRoutes(db)
}
