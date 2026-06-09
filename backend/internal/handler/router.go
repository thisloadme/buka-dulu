package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/middleware"
	"github.com/riyantobudi/bukadulu/internal/service"
)

func NewRouter(
	authSvc *service.AuthService,
	ventureSvc *service.VentureService,
	ideaSvc *service.IdeaService,
	customerSvc *service.CustomerService,
	menuSvc *service.MenuService,
	costSvc *service.CostService,
	missionSvc *service.MissionService,
	evidenceSvc *service.EvidenceService,
	scoringSvc *service.ScoringService,
	mentorSvc *service.MentorService,
) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CORS)
	r.Use(middleware.Recovery)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "version": "1.0.0"})
	})

	// Auth
	ah := NewAuthHandler(authSvc)
	r.Post("/api/v1/auth/register", ah.Register)
	r.Post("/api/v1/auth/login", ah.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authSvc))

		// Sprint 1
		vh := NewVentureHandler(ventureSvc)
		r.Post("/api/v1/ventures", vh.Create)
		r.Get("/api/v1/ventures", vh.List)
		r.Get("/api/v1/ventures/{id}", vh.GetByID)
		r.Put("/api/v1/ventures/{id}", vh.Update)

		ih := NewIdeaHandler(ideaSvc)
		r.Post("/api/v1/ventures/{id}/idea", ih.Capture)
		r.Get("/api/v1/ventures/{id}/idea", ih.Get)
		r.Put("/api/v1/ventures/{id}/idea", ih.Update)
		r.Post("/api/v1/ventures/{id}/idea/process", ih.Process)
		r.Post("/api/v1/ventures/{id}/idea/confirm", ih.Confirm)

		// Sprint 2
		ch := NewCustomerHandler(customerSvc)
		r.Post("/api/v1/ventures/{id}/customer", ch.Create)
		r.Get("/api/v1/ventures/{id}/customer", ch.Get)
		r.Post("/api/v1/ventures/{id}/customer/confirm", ch.Confirm)

		mh := NewMenuHandler(menuSvc)
		r.Post("/api/v1/ventures/{id}/menus", mh.Create)
		r.Get("/api/v1/ventures/{id}/menus", mh.List)
		r.Put("/api/v1/ventures/{id}/menus/{menuId}", mh.Update)
		r.Delete("/api/v1/ventures/{id}/menus/{menuId}", mh.Delete)
		r.Post("/api/v1/ventures/{id}/menus/focus", mh.Focus)

		cosh := NewCostHandler(costSvc)
		r.Post("/api/v1/ventures/{id}/ingredients", cosh.AddIngredient)
		r.Get("/api/v1/ventures/{id}/ingredients", cosh.ListIngredients)
		r.Delete("/api/v1/ventures/{id}/ingredients/{ingredientId}", cosh.DeleteIngredient)
		r.Post("/api/v1/ventures/{id}/packaging", cosh.AddPackaging)
		r.Get("/api/v1/ventures/{id}/packaging", cosh.ListPackaging)
		r.Post("/api/v1/ventures/{id}/cost/calculate/{menuId}", cosh.Calculate)
		r.Get("/api/v1/ventures/{id}/cost/summary/{menuId}", cosh.GetSummary)
		r.Get("/api/v1/ventures/{id}/cost/summaries", cosh.GetAllSummaries)
		r.Post("/api/v1/ventures/{id}/cost/confirm", cosh.Confirm)

		// Sprint 3
		msh := NewMissionHandler(missionSvc)
		r.Get("/api/v1/ventures/{id}/missions", msh.List)
		r.Post("/api/v1/ventures/{id}/missions/generate", msh.Generate)
		r.Post("/api/v1/ventures/{id}/missions", msh.Create)
		r.Post("/api/v1/ventures/{id}/missions/{missionId}/accept", msh.Accept)

		eh := NewEvidenceHandler(evidenceSvc)
		r.Post("/api/v1/ventures/{id}/evidence", eh.Upload)
		r.Get("/api/v1/ventures/{id}/evidence/mission/{missionId}", eh.ListByMission)
		r.Get("/api/v1/ventures/{id}/evidence/{evidenceId}", eh.GetWithReview)
		r.Post("/api/v1/ventures/{id}/evidence/{evidenceId}/review", eh.Review)
		r.Post("/api/v1/ventures/{id}/evidence/{evidenceId}/override", eh.Override)

		// Sprint 4: Scoring
		sch := NewScoringHandler(scoringSvc)
		r.Get("/api/v1/ventures/{id}/score", sch.GetScore)
		r.Post("/api/v1/ventures/{id}/score/calculate", sch.Calculate)
		r.Post("/api/v1/ventures/{id}/score/decision", sch.GenerateDecision)
		r.Get("/api/v1/ventures/{id}/score/decision", sch.GetDecision)

		// Sprint 4: Mentor
		mt := NewMentorHandler(mentorSvc)
		r.Get("/api/v1/mentor/mentees", mt.ListMentees)
		r.Post("/api/v1/ventures/{id}/mentor/comments", mt.AddComment)
	})

	return r
}
