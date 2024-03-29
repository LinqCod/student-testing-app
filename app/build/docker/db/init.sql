CREATE TABLE IF NOT EXISTS "groups" (
                          "id" bigserial PRIMARY KEY,
                          "title" varchar(10) NOT NULL
);

CREATE TYPE "user_role" AS ENUM ('student', 'teacher');

CREATE TABLE IF NOT EXISTS "users" (
                         "id" bigserial PRIMARY KEY,
                         "full_name" varchar NOT NULL,
                         "email" varchar NOT NULL,
                         "password" varchar NOT NULL,
                         "role" user_role
);

CREATE TABLE IF NOT EXISTS "students" (
                            "id" bigserial PRIMARY KEY,
                            personal_number varchar NOT NULL,
                            group_id bigint NOT NULL REFERENCES groups(id)
) INHERITS(users);

CREATE TABLE IF NOT EXISTS "teachers" (
    "id" bigserial PRIMARY KEY
) INHERITS(users);

CREATE TABLE IF NOT EXISTS "subjects" (
                            "id" bigserial PRIMARY KEY,
                            "title" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "groups_subjects" (
                                   "group_id" bigint NOT NULL REFERENCES groups(id),
                                   "subject_id" bigint NOT NULL REFERENCES subjects(id),
                                   CONSTRAINT groups_subjects_pkey PRIMARY KEY (group_id, subject_id)
);

CREATE TABLE IF NOT EXISTS "teachers_groups" (
                                   "teacher_id" bigint NOT NULL REFERENCES teachers(id),
                                   "group_id" bigint NOT NULL REFERENCES groups(id),
                                   CONSTRAINT teachers_groups_pkey PRIMARY KEY (teacher_id, group_id)
);

CREATE TABLE IF NOT EXISTS "teachers_subjects" (
                                     "teacher_id" bigint NOT NULL REFERENCES teachers(id),
                                     "subject_id" bigint NOT NULL REFERENCES subjects(id),
                                     CONSTRAINT teachers_subjects_pkey PRIMARY KEY (teacher_id, subject_id)
);

CREATE TABLE IF NOT EXISTS "task_categories" (
                                   "id" bigserial PRIMARY KEY,
                                   "title" varchar NOT NULL,
                                   "subject_id" bigint NOT NULL REFERENCES subjects(id)
);

CREATE TABLE IF NOT EXISTS "tasks" (
                         "id" bigserial PRIMARY KEY,
                         "text" varchar NOT NULL,
                         "category_id" bigint NOT NULL REFERENCES task_categories(id)
);

CREATE TABLE IF NOT EXISTS "task_answers" (
                                "id" bigserial PRIMARY KEY,
                                "text" varchar NOT NULL,
                                "is_right" boolean NOT NULL,
                                "task_id" bigint NOT NULL REFERENCES tasks(id)
);
