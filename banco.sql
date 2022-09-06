CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;

$$ LANGUAGE plpgsql;

CREATE TABLE "equipes" (
  "id_equipe" bigserial PRIMARY KEY,
  "nome_equipe" varchar(80) NOT NULL,
 "created_at" timestamp default now() NOT NULL,
 "updated_at" timestamp default now() not null
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "equipes"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "pessoas" (
  "id_pessoa" bigserial PRIMARY KEY,
  "nome_pessoa" varchar(80) NOT NULL,
  "funcao_pessoa" varchar(80) NOT NULL,
  "equipe_id" bigint,
  "favoritar" int NOT NULL DEFAULT 0,
  "data_contratacao" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp not null default (now()),
  constraint fk_equipe
	foreign key(equipe_id)
		references equipes(id_equipe) on delete set null
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "pessoas"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "projetos" (
  "id_projeto" bigserial PRIMARY KEY,
  "nome_projeto" varchar(80) NOT NULL,
  "descricao_projeto" text NOT NULL,
  "equipe_id" bigint,
  "status" varchar(80) NOT NULL DEFAULT 'Em planejamento',
  "data_inicio" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "data_conclusao" timestamp,
  constraint fk_equipe
	foreign key(equipe_id)
		references equipes(id_equipe) on delete set null
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "projetos"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "tasks" (
  "id_task" bigserial PRIMARY KEY,
  "descricao_task" text NOT NULL,
  "comentario" text,
  "pessoa_id" bigint,
  "projeto_id" bigint NOT NULL,
  "status" varchar(80) NOT NULL DEFAULT 'Em planejamento',
  "nivel" varchar(80) NOT NULL,
  "created_at" timestamp not null default (now()),
  "updated_at" timestamp not null default (now()),
  constraint fk_pessoa
	foreign key(pessoa_id)
		references pessoas(id_pessoa) on delete set null,
  constraint fk_projeto
	foreign key(projeto_id)
		references projetos(id_projeto) on delete cascade
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "tasks"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "user" (
  "id_user" bigserial PRIMARY KEY,
  "nome_user" varchar(80) not null,
  "email" varchar(80) not null unique,
  "senha" varchar(80) not null,
  "created_at" timestamp not null default (now()),
  "updated_at" timestamp not null default (now())
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE INDEX ON "equipes" ("id_equipe");

CREATE INDEX ON "pessoas" ("id_pessoa");

CREATE INDEX ON "pessoas" ("nome_pessoa");

CREATE INDEX ON "pessoas" ("equipe_id");

CREATE INDEX ON "projetos" ("id_projeto");

CREATE INDEX ON "projetos" ("nome_projeto");

CREATE INDEX ON "projetos" ("equipe_id");

CREATE INDEX ON "tasks" ("id_task");

CREATE INDEX ON "tasks" ("pessoa_id");

CREATE INDEX ON "tasks" ("projeto_id");

COMMENT ON COLUMN "pessoas"."equipe_id" IS 'pode ser nulo';

COMMENT ON COLUMN "projetos"."data_conclusao" IS 'será preenchido automaticamente quando um projeto for marcado como concluído';