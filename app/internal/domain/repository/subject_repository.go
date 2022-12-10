package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	GetSubjectByGroupIdQuery = `
		SELECT subjects.id, subjects.title
		FROM groups_subjects
		INNER JOIN subjects 
		ON subjects.id=groups_subjects.subject_id 
		WHERE groups_subjects.group_id = $1 
	`
)

type SubjectRepository interface {
	GetSubjectsByGroupId(groupId int64) ([]entity.Subject, error)
}

type SubjectRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewSubjectRepository(ctx context.Context, db *sql.DB) SubjectRepository {
	return &SubjectRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s SubjectRepositoryImpl) GetSubjectsByGroupId(groupId int64) ([]entity.Subject, error) {
	rows, err := s.db.QueryContext(s.ctx, GetSubjectByGroupIdQuery, groupId)
	if err != nil {
		return nil, err
	}

	var subjects []entity.Subject
	for rows.Next() {
		var subject entity.Subject
		if err = rows.Scan(&subject.Id, &subject.Title); err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}
