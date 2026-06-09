package main

import (
	"fmt"
	"os"

	"github.com/riyantobudi/bukadulu/internal/config"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

func main() {
	cfg := config.Load()
	db, err := config.InitDB(cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DB error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	config.RunMigrations(db, "migrations/001_init.up.sql")
	userRepo := repository.NewUserRepository(db)

	u, err := userRepo.FindByEmail("fresh@test.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "FindByEmail error: %v\n", err)
	} else {
		fmt.Printf("Found: %+v\n", u)
		fmt.Printf("Hash: %s\n", u.PasswordHash[:20]+"...")
	}

	// Also test scanning all users
	rows, err := db.Query("SELECT id, email, phone, password_hash FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query error: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id, email, phone, hash string
		rows.Scan(&id, &email, &phone, &hash)
		fmt.Printf("DB row: id=%s email=%s phone='%s'\n", id, email, phone)
	}
}
