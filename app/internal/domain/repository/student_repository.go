package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	SaveStudentQuery       = "INSERT INTO students (full_name, email, password, role, personal_number, group_id) VALUES ($1, $2, $3, $4, $5, $6)"
	GetStudentDetailsQuery = `
		SELECT groups.id, groups.title, students.id, students.full_name, students.email, students.password, students.role, students.personal_number 
		FROM groups 
		INNER JOIN students 
		ON students.group_id=groups.id 
		WHERE students.id = $1 
		LIMIT 1;
	`
	GetStudentByEmailQuery = `
		SELECT groups.id, groups.title, students.id, students.full_name, students.email, students.password, students.role, students.personal_number 
		FROM groups 
		INNER JOIN students 
		ON students.group_id=groups.id 
		WHERE students.email = $1 
		LIMIT 1;
	`
	UpdateStudentQuery        = "UPDATE students SET full_name = $2, email = $3, password = $4, personal_number = $5, group_id = $6 WHERE id = $1"
	GetStudentGroupTitleQuery = "SELECT title FROM groups WHERE id = $1 LIMIT 1"
)

type StudentRepository interface {
	SaveStudent(student *entity.Student) (*entity.Student, error)
	GetStudentDetails(studentId int64) (*entity.Student, error)
	UpdateStudent(student *entity.Student) (*entity.Student, error)
	GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error)
}

type StudentRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewStudentRepository(ctx context.Context, db *sql.DB) StudentRepository {
	return &StudentRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (r StudentRepositoryImpl) SaveStudent(student *entity.Student) (*entity.Student, error) {
	err := r.db.QueryRowContext(
		r.ctx,
		SaveStudentQuery,
		student.FullName,
		student.Email,
		student.Password,
		student.Role,
		student.PersonalNumber,
		student.Group.Id,
	).Err()
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (r StudentRepositoryImpl) GetStudentDetails(studentId int64) (*entity.Student, error) {
	var student entity.Student
	if err := r.db.QueryRowContext(r.ctx, GetStudentDetailsQuery, studentId).Scan(
		&student.Group.Id,
		&student.Group.Title,
		&student.Id,
		&student.FullName,
		&student.Email,
		&student.Password,
		&student.Role,
		&student.PersonalNumber,
	); err != nil {
		return nil, err
	}

	return &student, nil
}

func (r StudentRepositoryImpl) UpdateStudent(student *entity.Student) (*entity.Student, error) {
	err := r.db.QueryRowContext(
		r.ctx,
		UpdateStudentQuery,
		student.Id,
		student.FullName,
		student.Email,
		student.Password,
		student.PersonalNumber,
		student.Group.Id,
	).Err()
	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRowContext(r.ctx, GetStudentGroupTitleQuery, student.Group.Id).
		Scan(&student.Group.Title); err != nil {
		return nil, err
	}

	return student, nil
}

func (r StudentRepositoryImpl) GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error) {
	var student entity.Student
	if err := r.db.QueryRowContext(r.ctx, GetStudentByEmailQuery, login.Email).Scan(
		&student.Group.Id,
		&student.Group.Title,
		&student.Id,
		&student.FullName,
		&student.Email,
		&student.Password,
		&student.Role,
		&student.PersonalNumber,
	); err != nil {
		return nil, err
	}

	return &student, nil
}
