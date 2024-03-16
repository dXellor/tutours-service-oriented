DROP TABLE IF EXISTS "Encounters";

CREATE TABLE IF NOT EXISTS "Encounters"
(
    "Id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    "UserId" integer NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "Description" text COLLATE pg_catalog."default" NOT NULL,
    "Latitude" double precision NOT NULL,
    "Longitude" double precision NOT NULL,
    "Xp" integer NOT NULL,
    "Status" integer NOT NULL,
    "Type" integer NOT NULL,
    "Range" double precision NOT NULL,
    "Image" text COLLATE pg_catalog."default",
    "PeopleCount" integer,
    "ApprovalStatus" integer NOT NULL DEFAULT 0,
    "ImageLatitude" double precision,
    "ImageLongitude" double precision,
    CONSTRAINT "PK_Encounters" PRIMARY KEY ("Id")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS "Encounters"
    OWNER to postgres;

INSERT INTO "Encounters"("UserId","Name","Description","Latitude","Longitude","Xp","Status","Type","Range","Image","PeopleCount","ApprovalStatus","ImageLatitude","ImageLongitude") 
VALUES 
(19, 'yes', 'yes', '0', '0', 2, 1, 1, '1', NULL, NULL, 1, NULL, NULL),
(19, 'Train? Bus? Find it ', 'Find this image', '45.26495528894582', '19.829776374241465', 150, 1, 1, '1', NULL, 0, 1, '45.26483069286267', '19.829197272945347'),
(19, 'Find a cat', 'Find a cat and pet it', '45.256125909748846', '19.844571164364595', 20, 1, 2, '1', NULL, 0, 1, '0', '0'),
(19, 'Shopping Spree', 'Buy something in Panda', '45.25958524334332', '19.83216331284152', 20, 1, 2, '1', NULL, 0, 1, '0', '0'),
(19, 'Play in the park', 'Play in the park and have fun', '45.259147646137286', '19.833865482175163', 10, 1, 2, '1', NULL, 0, 1, '0', '0'),
(19, 'Friends Night Out', 'Go out with friends, requires 5+ people', '45.25547471700954', '19.843383009550617', 150, 1, 0, '1', NULL, 5, 1, '0', '0'),
(19, 'Bridge Fun', 'Go over the bridge and don''t jump', '45.234236388477065', '19.849983595808702', 10, 1, 2, '1', NULL, 0, 1, '0', '0'),
(19, 'Ice Magic', 'Go ice skating', '45.25541933237368', '19.851234425399287', 10, 1, 2, '1', NULL, 0, 1, '0', '0'),
(19, 'Friend Meeting', 'Meet 10+ people here', '45.2631098125524', '19.83646162360244', 100, 1, 0, '1', NULL, 10, 1, '0', '0'),
(19, 'Travel?', 'Meet some friends here and go on a bus or train ride.', '45.26539871706191', '19.830793644526153', 150, 1, 0, '1', NULL, 10, 1, '0', '0'),
(19, 'School Time', 'Go to Elektro', '45.25014657699679', '19.834380594231572', 200, 1, 0, '1', NULL, 5, 1, '0', '0');
