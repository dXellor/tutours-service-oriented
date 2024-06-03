drop table if exists "Users";

create table if not exists "Users"
(
    "Id"                bigint generated by default as identity
        constraint "PK_Users"
            primary key,
    "Username"          text    not null,
    "Password"          text    not null,
    "Role"              integer not null,
    "IsActive"          boolean not null,
    "Email"             text    not null,
    "IsBlocked"         boolean not null,
    "IsEnabled"         boolean not null,
    "VerificationToken" text    not null
) tablespace pg_default;

alter table "Users"
    owner to postgres;

INSERT INTO "Users"(
    "Username", "Password", "Role", "IsActive", "Email", "IsBlocked", "IsEnabled", "VerificationToken")
VALUES
    ('pera', 'pera123', 2, 'true', 'pera@gmail.com', 'false', true, 'token1'),
    ('mika', 'mika123', 2, 'true', 'mika@gmail.com', 'false', true, 'token2'),
    ('ana', 'ana456', 2, 'true', 'ana@gmail.com', 'false', true, 'token3'),
    ('jovan', 'jovanPass', 2, 'true', 'jovan@gmail.com', 'false', true, 'token4'),
    ('marko', 'markoPass', 2, 'true', 'marko@gmail.com', 'false', true, 'token5'),
    ('milica', 'milica567', 2, 'true', 'milica@gmail.com', 'false', true, 'token6'),
    ('stefan', 'stefanPass', 2, 'true', 'stefan@gmail.com', 'false', true, 'token7'),
    ('tamara', 'tamaraPass', 2, 'true', 'tamara@gmail.com', 'false', true, 'token8'),
    ('nikola', 'nikola789', 2, 'true', 'nikola@gmail.com', 'false', true, 'token9'),
    ('lena', 'lenaPass', 2, 'true', 'lena@gmail.com', 'false', true, 'token10'),
    ('petar', 'petarPass', 2, 'true', 'petar@gmail.com', 'false', true, 'token11'),
    ('jelena', 'jelena123', 2, 'true', 'jelena@gmail.com', 'false', true, 'token12'),
    ('miroslav', 'miroslavPass', 2, 'true', 'miroslav@gmail.com', 'false', true, 'token13'),
    ('sofija', 'sofijaPass', 2, 'true', 'sofija@gmail.com', 'false', true, 'token14'),
    ('aleksa', 'aleksaPass', 2, 'true', 'aleksa@gmail.com', 'false', true, 'token15'),
    ('sandra', 'sandraPass', 1, 'true', 'sandra@gmail.com', 'false', true, 'token16'),
    ('kosta', 'kostaPass', 1, 'true', 'kosta@gmail.com', 'false', true, 'token17'),
    ('dragan', 'draganPass', 1, 'true', 'dragan@gmail.com', 'false', true, 'token18'),
    ('dragana', 'draganaPass', 0, 'true', 'dragana@gmail.com', 'false', true, 'token19'),
    ('bogdan', 'bogdanPass', 0, 'true', 'bogdan@gmail.com', 'false', true, 'token20'),
    ('lenka', 'lenkaPass', 0, 'true', 'lenka@gmail.com', 'false', true, 'token21');

drop table if exists "People";

create table "People"
(
    "Id"           bigint generated by default as identity
        constraint "PK_People"
            primary key,
    "UserId"       bigint  not null
        constraint "FK_People_Users_UserId"
            references "Users" ("Id")
            on delete cascade,
    "Name"         text    not null,
    "Surname"      text    not null,
    "Email"        text    not null,
    "ProfileImage" text    not null,
    "Biography"    text    not null,
    "Quote"        text    not null,
    "XP"           integer not null,
    "Level"        integer not null
) tablespace pg_default;

alter table "People"
    owner to postgres;

INSERT INTO "People"(
    "UserId", "Name", "Surname", "Email", "ProfileImage", "Biography", "Quote", "XP", "Level")
VALUES
    (1, 'Pera', 'Stanković', 'pera@gmail.com', 'https://www.svgrepo.com/show/384670/account-avatar-profile-user.svg', 'Passionate traveler exploring the world.', 'Adventure awaits!', 10, 1),
    (2, 'Minja', 'Perović', 'mika@gmail.com', 'https://www.svgrepo.com/show/384674/account-avatar-profile-user-11.svg', 'Lover of new experiences and cultures.', 'Wander often, wonder always.', 10, 1),
    (3, 'Ana', 'Hristić', 'ana@gmail.com', 'https://www.svgrepo.com/show/384671/account-avatar-profile-user-14.svg', 'Adventurous soul seeking thrill in every journey.', 'Explore, dream, discover.', 10, 1),
    (4, 'Jovan', 'Zdravković', 'jovan@gmail.com', 'https://www.svgrepo.com/show/384676/account-avatar-profile-user-6.svg', 'Traveling the world one destination at a time.', 'Collect moments, not things.', 10, 1),
    (5, 'Marko', 'Lazović', 'marko@gmail.com', 'https://www.svgrepo.com/show/384678/account-avatar-profile-user-9.svg', 'Capturing memories and stories through travel.', 'Adventure is worthwhile in itself.', 10, 1),
    (6, 'Milica', 'Savić', 'milica@gmail.com', 'https://www.svgrepo.com/show/384682/account-avatar-profile-user-10.svg', 'Discovering hidden gems and local flavors.', 'Travel far, love wide.', 10, 1),
    (7, 'Stefan', 'Dragović', 'stefan@gmail.com', 'https://www.svgrepo.com/show/384683/account-avatar-profile-user-8.svg', 'Adventure seeker with a passion for photography.', 'Explore the unexplored.', 10, 1),
    (8, 'Tamara', 'Tanacković', 'tamara@gmail.com', 'https://www.svgrepo.com/show/384680/account-avatar-profile-user-4.svg', 'Making memories around the globe.', 'Travel is the only thing you can buy that makes you richer.', 10, 1),
    (9, 'Nikola', 'Praška', 'nikola@gmail.com', 'https://www.svgrepo.com/show/384677/account-avatar-profile-user-12.svg', 'Embracing the journey and the unknown.', 'Adventure awaits, go find it.', 10, 1),
    (10, 'Lena', 'Antonić', 'lena@gmail.com', 'https://www.svgrepo.com/show/384673/account-avatar-profile-user-5.svg', 'Exploring new cultures and making friends worldwide.', 'Adventure is out there!', 10, 1),
    (11, 'Petar', 'Mitić', 'petar@gmail.com', 'https://www.svgrepo.com/show/384672/account-avatar-profile-user-7.svg', 'Seeking adrenaline in every corner of the world.', 'Travel and make memories.', 10, 1),
    (12, 'Jelena', 'Jovanović', 'jelena@gmail.com', 'https://www.svgrepo.com/show/384684/account-avatar-profile-user-15.svg', 'Creating a tapestry of travel experiences.', 'Adventure is calling, answer it.', 10, 1),
    (13, 'Miroslav', 'Jakovljević', 'miroslav@gmail.com', 'https://www.svgrepo.com/show/384675/account-avatar-profile-user-2.svg', 'Chasing sunsets and collecting stories.', 'Travel far, love deep.', 10, 1),
    (14, 'Sofija', 'Kostić', 'sofija@gmail.com', 'https://www.svgrepo.com/show/384671/account-avatar-profile-user-14.svg', 'Connecting with the world through travel.', 'Wander often, wonder always.', 10, 1),
    (15, 'Aleksa', 'Popović', 'aleksa@gmail.com', 'https://www.svgrepo.com/show/384670/account-avatar-profile-user.svg', 'Discovering the beauty of diverse landscapes.', 'Adventure is worthwhile in itself.', 10, 1),
    (16, 'Sandra', 'Maljević', 'sandra@gmail.com', 'https://www.svgrepo.com/show/384680/account-avatar-profile-user-4.svg', 'Exploring cultures and sharing stories.', 'Life is an adventure, make the most of it.', 10, 1),
    (17, 'Kosta', 'Vasiljević', 'kosta@gmail.com', 'https://www.svgrepo.com/show/384678/account-avatar-profile-user-9.svg', 'Adventure lover and thrill seeker.', 'Travel, explore, live.', 10, 1),
    (18, 'Dragan', 'Petrović', 'dragan@gmail.com', 'https://www.svgrepo.com/show/384682/account-avatar-profile-user-10.svg', 'Journeying through life one trip at a time.', 'Adventure is a way of life.', 10, 1),
    (19, 'Dragana', 'Radić', 'dragana@gmail.com', 'https://www.svgrepo.com/show/384677/account-avatar-profile-user-12.svg', 'Finding joy in every destination.', 'The world is a book, and those who do not travel read only a page.', 10, 1),
    (20, 'Bogdan', 'Janković', 'bogdan@gmail.com', 'https://www.svgrepo.com/show/384671/account-avatar-profile-user-14.svg', 'Exploring the wonders of the world.', 'Adventure is out there, find it.', 10, 1),
    (21, 'Lenka', 'Ristić', 'lenka@gmail.com', 'https://www.svgrepo.com/show/384684/account-avatar-profile-user-15.svg', 'Discovering the beauty of the world.', 'Travel far, live well.', 10, 1);
