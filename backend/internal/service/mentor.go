package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type MentorService struct {
	ventureRepo *repository.VentureRepository
	userRepo    *repository.UserRepository
}

func NewMentorService(ventureRepo *repository.VentureRepository, userRepo *repository.UserRepository) *MentorService {
	return &MentorService{ventureRepo: ventureRepo, userRepo: userRepo}
}

type MenteeSummary struct {
	Founder *domain.User    `json:"founder"`
	Venture *domain.Venture `json:"venture"`
}

func (s *MentorService) ListMentees(mentorID string) ([]*MenteeSummary, error) {
	ventures, err := s.ventureRepo.FindByOwner(mentorID)
	if err != nil {
		return nil, err
	}
	if ventures == nil {
		return []*MenteeSummary{}, nil
	}

	var mentees []*MenteeSummary
	for _, v := range ventures {
		if v.OwnerUserID == mentorID {
			continue
		}
		founder, err := s.userRepo.FindByID(v.OwnerUserID)
		if err != nil {
			continue
		}
		mentees = append(mentees, &MenteeSummary{Founder: founder, Venture: v})
	}
	return mentees, nil
}

func (s *MentorService) AddComment(mentorID, ventureID, missionID, evidenceID, content string) (*domain.MentorComment, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	return &domain.MentorComment{
		ID: uuid.New().String(), MentorID: mentorID, VentureID: ventureID,
		MissionID: missionID, EvidenceID: evidenceID, Content: content,
		IsDeleted: false, CreatedAt: now,
	}, nil
}
