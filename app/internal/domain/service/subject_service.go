package service

import (
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
	"github.com/linqcod/student-testing-app/app/internal/domain/repository"
)

type SubjectService interface {
	GetSubjectsByGroupId(groupId int64) ([]entity.Subject, error)
}

type SubjectServiceImpl struct {
	subjectRepo repository.SubjectRepository
}

func NewSubjectService(subjectRepo repository.SubjectRepository) SubjectService {
	return &SubjectServiceImpl{
		subjectRepo: subjectRepo,
	}
}

func (s SubjectServiceImpl) GetSubjectsByGroupId(groupId int64) ([]entity.Subject, error) {
	subjects, err := s.subjectRepo.GetSubjectsByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
