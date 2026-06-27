package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/riyantobudi/bukadulu/internal/config"
	"github.com/riyantobudi/bukadulu/internal/handler"
	"github.com/riyantobudi/bukadulu/internal/repository"
	"github.com/riyantobudi/bukadulu/internal/service"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg := config.Load()
	db, err := config.InitDB(cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Run migrations
	migrations := []string{
		"migrations/001_init.up.sql",
		"migrations/002_otp_verification.up.sql",
	}
	for _, m := range migrations {
		if err := config.RunMigrations(db, m); err != nil {
			slog.Error("failed to run migration", "file", m, "error", err)
			os.Exit(1)
		}
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	ventureRepo := repository.NewVentureRepository(db)
	ideaRepo := repository.NewIdeaRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	ingredientRepo := repository.NewIngredientRepository(db)
	missionRepo := repository.NewMissionRepository(db)
	evidenceRepo := repository.NewEvidenceRepository(db)
	scoreRepo := repository.NewScoreRepository(db)

	// LLM
	llmCfg := cfg.GetLLMConfig()
	llmProvider := service.ProviderFactory(llmCfg.Provider, llmCfg.BaseURL, llmCfg.APIKey, llmCfg.Model)
	llmSvc := service.NewLLMService(llmProvider)
	slog.Info("LLM service", "provider", llmSvc.ProviderName())

	// Services
	emailCfg := service.EmailConfig{
		Host:     cfg.SMTPHost,
		Port:     cfg.SMTPPort,
		User:     cfg.SMTPUser,
		Password: cfg.SMTPPassword,
		From:     cfg.SMTPFrom,
	}
	emailSvc := service.NewEmailService(emailCfg)
	authSvc := service.NewAuthService(userRepo, emailSvc, cfg.JWTSecret, cfg.JWTExpiry, cfg.OTPExpiry)
	ventureSvc := service.NewVentureService(ventureRepo)
	ideaSvc := service.NewIdeaService(ideaRepo, llmSvc, ventureSvc)
	customerSvc := service.NewCustomerService(customerRepo, ventureSvc)
	menuSvc := service.NewMenuService(menuRepo, llmSvc, ventureSvc)
	costSvc := service.NewCostService(ingredientRepo, ventureSvc)
	missionSvc := service.NewMissionService(missionRepo, llmSvc, ventureSvc)
	evidenceSvc := service.NewEvidenceService(evidenceRepo, missionRepo, llmSvc, ventureSvc)
	scoringSvc := service.NewScoringService(scoreRepo, ventureRepo, ideaRepo, menuRepo, ingredientRepo, missionRepo, evidenceRepo, ventureSvc)
	mentorSvc := service.NewMentorService(ventureRepo, userRepo)

	router := handler.NewRouter(authSvc, ventureSvc, ideaSvc, customerSvc, menuSvc, costSvc, missionSvc, evidenceSvc, scoringSvc, mentorSvc)

	addr := fmt.Sprintf(":%d", cfg.Port)
	slog.Info("server starting", "port", cfg.Port)

	if err := http.ListenAndServe(addr, router); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
