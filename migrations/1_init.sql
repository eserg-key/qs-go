
SELECT 'up SQL query';

BEGIN;


CREATE TABLE public.handbooks
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) UNIQUE NOT NULL,
    handbook_name VARCHAR(250) UNIQUE NOT NULL,
    project_code VARCHAR(250) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL,
    updated_at    TIMESTAMPTZ
);

CREATE TABLE public.handbook_type_fields
(
  id SERIAL PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  type VARCHAR(250) NOT NULL
);

CREATE TABLE public.handbook_metadata
(
  id UUID PRIMARY KEY,
  sort INTEGER  NOT NULL,
  handbook_name VARCHAR(250) NOT NULL,
  handbook_type_field_id SERIAL REFERENCES public.handbook_type_fields (id) NOT NULL,
  handbook_field_name VARCHAR(250) NOT NULL,
  handbook_field_rus_name VARCHAR(250) NOT NULL,
  handbook_children_id INTEGER,
  handbook_children_column VARCHAR(250),
  created_at    TIMESTAMPTZ NOT NULL,
  updated_at    TIMESTAMPTZ
);

INSERT INTO public.handbook_type_fields (name, type) VALUES ('Текст', 'text');
INSERT INTO public.handbook_type_fields (name, type) VALUES ('Число', 'int');
INSERT INTO public.handbook_type_fields (name, type) VALUES ('Комментарий', 'comment');
INSERT INTO public.handbook_type_fields (name, type) VALUES ('Города', 'select_city');
INSERT INTO public.handbook_type_fields (name, type) VALUES ('Справочник', 'select_handbook');

SELECT 'down SQL query';

COMMIT;