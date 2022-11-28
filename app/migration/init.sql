CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "first_name" varchar NOT NULL,
                         "second_name" varchar NOT NULL,
                         "email" varchar NOT NULL,
                         "password" varchar NOT NULL
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("second_name");

CREATE INDEX ON "users" ("email");