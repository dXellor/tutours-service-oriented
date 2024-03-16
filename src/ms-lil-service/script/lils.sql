DROP TABLE IF EXISTS  lils;
DROP SEQUENCE IF EXISTS lil_id_seq;
CREATE SEQUENCE lil_id_seq;

CREATE TABLE IF NOT EXISTS public.lils
(
    id int NOT NULL DEFAULT nextval('lil_id_seq'), -- NOT NULL DEFAULT nextval = SERIAL, auto increment
    name text COLLATE pg_catalog."default" NOT NULL,
    nickname text COLLATE pg_catalog."default" NOT NULL,
    age bigint,
    CONSTRAINT lils_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.lils
    OWNER to postgres;

INSERT INTO public.lils(name, nickname, age)VALUES ('lil1', 'lil1', 2);
INSERT INTO public.lils(name, nickname, age)VALUES ('lil2', 'lil2', 4);
INSERT INTO public.lils(name, nickname, age)VALUES ('lil3', 'lil3', 5);