drop table if exists "Tours";

create table if not exists "Tours"
(
    "Id" bigint generated by default as identity (start 1 increment 1) primary key,
    "UserId" integer not null,
    "Name" text not null,
    "Description" text not null,
    "Price" double precision not null,
    "Duration" integer,
    "Distance" double precision,
    "Difficulty" integer,
    "TransportType" integer,
    "Status" integer,
    "StatusUpdateTime" timestamp with time zone not null,
    "Tags" text[]
) tablespace pg_default;

alter table "Tours"
    owner to postgres;


insert into "Tours"(
    "UserId", "Name", "Description", "Price", "Duration", "Distance", "Difficulty", "TransportType", "Status", "StatusUpdateTime", "Tags")
VALUES
    (1, 'Zlatibor Nature Escape', 'Discover the natural beauty of Zlatibor.', 1500, 37, 7, 2, 3, 1, '2024-02-16', ARRAY['nature', 'escape', 'Zlatibor']),
    (2, 'Beograd Cultural Delight', 'Immerse yourself in the rich culture of Belgrade.', 1800, 9, 1, 2, 3, 1, '2024-01-26', ARRAY['culture', 'city', 'Belgrade']),
    (3, 'Fruška Gora Monastery Tour', 'Explore the historic monasteries of Fruška gora.', 1200, 10, 5, 2, 3, 1, '2024-03-18', ARRAY['history', 'monastery', 'national park']),
    (4, 'Novi Sad Tour', 'Explore the rich history of the city through its landmarks.', 1500, 60, 6, 2, 3, 1, '2024-01-21', ARRAY['city', 'history', 'culture']),
    (5, 'Fruška gora', 'Take a walkand enjoy the natural beauty of the national park.', 1200, 10, 5, 2, 3, 1, '2024-02-11', ARRAY['mountain', 'national park']),
    (6, 'Ovčar-Kablar Gorge', 'Explore the breathtaking views of Ovčar-Kablar Gorge.', 1800, 32, 8, 2, 3, 1, '2024-02-16', ARRAY['nature', 'adventure']),
    (7, 'Subotica Heritage Tour', 'Discover the architectural wonders of Subotica.', 1600, 32, 8, 2, 3, 1, '2024-01-13', ARRAY['heritage', 'architecture', 'Subotica']),
    (8, 'Kragujevac History Walk', 'Walk through the history of Kragujevac.', 1400, 8, 2, 2, 3, 1, '2024-03-19', ARRAY['history', 'walk', 'Kragujevac']),
    (16, 'Uvac Nature Expedition', 'Experience the stunning nature of Uvac.', 1800, 16, 2, 2, 3, 1, '2024-03-24', ARRAY['nature', 'expedition', 'Uvac']);
