CREATE TABLE "users" (
                    "id" bigserial PRIMARY KEY,
                    "first_name" varchar NOT NULL,
                    "second_name" varchar NOT NULL,
                    "email" varchar NOT NULL,
                    "password" varchar NOT NULL
);

CREATE TABLE "task_categories" (
                    "id" bigserial PRIMARY KEY,
                    "title" varchar NOT NULL
);

CREATE TABLE "tasks" (
                    "id" bigserial PRIMARY KEY,
                    "category_id" bigint NOT NULL REFERENCES task_categories(id),
                    "description" varchar NOT NULL
);

DROP TABLE "task_answers";
CREATE TABLE "task_answers" (
                    "id" bigserial PRIMARY KEY,
                    "task_id" bigint NOT NULL REFERENCES tasks(id),
                    "text" varchar NOT NULL,
                    "is_right" boolean NOT NULL
);
