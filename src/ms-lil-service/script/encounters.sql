
--
-- Name: encounters; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA encounters;


ALTER SCHEMA encounters OWNER TO postgres;

--
-- Name: EncounterCompletions; Type: TABLE; Schema: encounters; Owner: postgres
--

CREATE TABLE encounters."EncounterCompletions" (
    "Id" bigint NOT NULL,
    "UserId" bigint NOT NULL,
    "LastUpdatedAt" timestamp with time zone NOT NULL,
    "EncounterId" bigint NOT NULL,
    "Xp" integer NOT NULL,
    "Status" integer NOT NULL
);


ALTER TABLE encounters."EncounterCompletions" OWNER TO postgres;

--
-- Name: EncounterCompletions_Id_seq; Type: SEQUENCE; Schema: encounters; Owner: postgres
--

ALTER TABLE encounters."EncounterCompletions" ALTER COLUMN "Id" ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME encounters."EncounterCompletions_Id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: Encounters; Type: TABLE; Schema: encounters; Owner: postgres
--

CREATE TABLE encounters."Encounters" (
    "Id" bigint NOT NULL,
    "UserId" integer NOT NULL,
    "Name" text NOT NULL,
    "Description" text NOT NULL,
    "Latitude" double precision NOT NULL,
    "Longitude" double precision NOT NULL,
    "Xp" integer NOT NULL,
    "Status" integer NOT NULL,
    "Type" integer NOT NULL,
    "Range" double precision NOT NULL,
    "Image" text,
    "PeopleCount" integer,
    "ApprovalStatus" integer DEFAULT 0 NOT NULL,
    "ImageLatitude" double precision,
    "ImageLongitude" double precision
);


ALTER TABLE encounters."Encounters" OWNER TO postgres;

--
-- Name: Encounters_Id_seq; Type: SEQUENCE; Schema: encounters; Owner: postgres
--

ALTER TABLE encounters."Encounters" ALTER COLUMN "Id" ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME encounters."Encounters_Id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: KeypointEncounters; Type: TABLE; Schema: encounters; Owner: postgres
--

CREATE TABLE encounters."KeypointEncounters" (
    "Id" bigint NOT NULL,
    "EncounterId" bigint NOT NULL,
    "KeyPointId" bigint NOT NULL,
    "IsRequired" boolean NOT NULL
);


ALTER TABLE encounters."KeypointEncounters" OWNER TO postgres;

--
-- Name: KeypointEncounters_Id_seq; Type: SEQUENCE; Schema: encounters; Owner: postgres
--

ALTER TABLE encounters."KeypointEncounters" ALTER COLUMN "Id" ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME encounters."KeypointEncounters_Id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: __EFMigrationsHistory; Type: TABLE; Schema: encounters; Owner: postgres
--

CREATE TABLE encounters."__EFMigrationsHistory" (
    "MigrationId" character varying(150) NOT NULL,
    "ProductVersion" character varying(32) NOT NULL
);


ALTER TABLE encounters."__EFMigrationsHistory" OWNER TO postgres;


--
-- Data for Name: EncounterCompletions; Type: TABLE DATA; Schema: encounters; Owner: postgres
--

COPY encounters."EncounterCompletions" ("Id", "UserId", "LastUpdatedAt", "EncounterId", "Xp", "Status") FROM stdin;
1	4	2024-01-18 00:00:00+00	100	28	2
2	7	2024-01-18 00:00:00+00	100	144	2
3	8	2024-01-18 00:00:00+00	100	103	2
4	9	2024-01-18 00:00:00+00	100	220	2
5	10	2024-01-18 00:00:00+00	100	12	2
6	11	2024-01-18 00:00:00+00	100	10	2
7	14	2024-01-18 00:00:00+00	100	55	2
8	19	2024-01-18 00:00:00+00	100	10	2
9	5	2024-01-18 00:00:00+00	100	328	2
10	6	2024-01-18 00:00:00+00	100	50	2
11	12	2024-01-18 00:00:00+00	100	150	2
12	13	2024-01-18 00:00:00+00	100	149	2
13	17	2024-01-18 00:00:00+00	100	33	2
14	18	2024-01-18 00:00:00+00	100	18	2
109	1	2024-01-18 14:11:23.118514+00	101	150	2
\.


--
-- Data for Name: Encounters; Type: TABLE DATA; Schema: encounters; Owner: postgres
--

COPY encounters."Encounters" ("Id", "UserId", "Name", "Description", "Latitude", "Longitude", "Xp", "Status", "Type", "Range", "Image", "PeopleCount", "ApprovalStatus", "ImageLatitude", "ImageLongitude") FROM stdin;
102	19	Find a cat	Find a cat and pet it	45.256125909748846	19.844571164364595	20	1	2	1	\N	0	1	0	0
103	19	Shopping Spree	Buy something in Panda	45.25958524334332	19.83216331284152	20	1	2	1	\N	0	1	0	0
104	19	Play in the park	Play in the park and have fun	45.259147646137286	19.833865482175163	10	1	2	1	\N	0	1	0	0
105	19	Friends Night Out	Go out with friends, requires 5+ people	45.25547471700954	19.843383009550617	150	1	0	1	\N	5	1	0	0
106	19	Bridge Fun	Go over the bridge and don't jump	45.234236388477065	19.849983595808702	10	1	2	1	\N	0	1	0	0
107	19	Ice Magic	Go ice skating	45.25541933237368	19.851234425399287	10	1	2	1	\N	0	1	0	0
108	19	Friend Meeting	Meet 10+ people here	45.2631098125524	19.83646162360244	100	1	0	1	\N	10	1	0	0
109	19	Travel?	Meet some friends here and go on a bus or train ride.	45.26539871706191	19.830793644526153	150	1	0	1	\N	10	1	0	0
110	19	School Time	Go to Elektro	45.25014657699679	19.834380594231572	200	1	0	1	\N	5	1	0	0
100	19	yes	yes	0	0	2	1	1	1	\N	\N	1	\N	\N
101	19	Train? Bus? Find it 	Find this image	45.26495528894582	19.829776374241465	150	1	1	1	data:image/jpeg;base64,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	0	1	45.26483069286267	19.829197272945347
\.


--
-- Data for Name: KeypointEncounters; Type: TABLE DATA; Schema: encounters; Owner: postgres
--

COPY encounters."KeypointEncounters" ("Id", "EncounterId", "KeyPointId", "IsRequired") FROM stdin;
\.


--
-- Data for Name: __EFMigrationsHistory; Type: TABLE DATA; Schema: encounters; Owner: postgres
--

COPY encounters."__EFMigrationsHistory" ("MigrationId", "ProductVersion") FROM stdin;
20231202131432_Initial	7.0.5
20231227193208_jk	7.0.5
20231227221108_awewerwerwerwerwerwer	7.0.5
20240116172818_mrzimsvojzivot	7.0.5
\.

-- Name: EncounterCompletions_Id_seq; Type: SEQUENCE SET; Schema: encounters; Owner: postgres
--

SELECT pg_catalog.setval('encounters."EncounterCompletions_Id_seq"', 109, true);


--
-- Name: Encounters_Id_seq; Type: SEQUENCE SET; Schema: encounters; Owner: postgres
--

SELECT pg_catalog.setval('encounters."Encounters_Id_seq"', 111, true);


--
-- Name: KeypointEncounters_Id_seq; Type: SEQUENCE SET; Schema: encounters; Owner: postgres
--

SELECT pg_catalog.setval('encounters."KeypointEncounters_Id_seq"', 100, true);



--
-- Name: EncounterCompletions PK_EncounterCompletions; Type: CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."EncounterCompletions"
    ADD CONSTRAINT "PK_EncounterCompletions" PRIMARY KEY ("Id");


--
-- Name: Encounters PK_Encounters; Type: CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."Encounters"
    ADD CONSTRAINT "PK_Encounters" PRIMARY KEY ("Id");


--
-- Name: KeypointEncounters PK_KeypointEncounters; Type: CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."KeypointEncounters"
    ADD CONSTRAINT "PK_KeypointEncounters" PRIMARY KEY ("Id");


--
-- Name: __EFMigrationsHistory PK___EFMigrationsHistory; Type: CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."__EFMigrationsHistory"
    ADD CONSTRAINT "PK___EFMigrationsHistory" PRIMARY KEY ("MigrationId");


--
-- Name: IX_EncounterCompletions_EncounterId; Type: INDEX; Schema: encounters; Owner: postgres
--

CREATE INDEX "IX_EncounterCompletions_EncounterId" ON encounters."EncounterCompletions" USING btree ("EncounterId");


--
-- Name: IX_KeypointEncounters_EncounterId; Type: INDEX; Schema: encounters; Owner: postgres
--

CREATE UNIQUE INDEX "IX_KeypointEncounters_EncounterId" ON encounters."KeypointEncounters" USING btree ("EncounterId");


--
-- Name: EncounterCompletions FK_EncounterCompletions_Encounters_EncounterId; Type: FK CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."EncounterCompletions"
    ADD CONSTRAINT "FK_EncounterCompletions_Encounters_EncounterId" FOREIGN KEY ("EncounterId") REFERENCES encounters."Encounters"("Id") ON DELETE CASCADE;


--
-- Name: KeypointEncounters FK_KeypointEncounters_Encounters_EncounterId; Type: FK CONSTRAINT; Schema: encounters; Owner: postgres
--

ALTER TABLE ONLY encounters."KeypointEncounters"
    ADD CONSTRAINT "FK_KeypointEncounters_Encounters_EncounterId" FOREIGN KEY ("EncounterId") REFERENCES encounters."Encounters"("Id") ON DELETE CASCADE;
