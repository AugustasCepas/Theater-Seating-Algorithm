--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-01-08 01:51:40

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: admin
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 24611)
-- Name: ranks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ranks (
    id integer NOT NULL,
    price numeric NOT NULL
);


ALTER TABLE public.ranks OWNER TO admin;

--
-- TOC entry 217 (class 1259 OID 24614)
-- Name: Rank_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public."Rank_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Rank_id_seq" OWNER TO admin;

--
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 217
-- Name: Rank_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public."Rank_id_seq" OWNED BY public.ranks.id;


--
-- TOC entry 209 (class 1259 OID 16386)
-- Name: layouts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.layouts (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.layouts OWNER TO admin;

--
-- TOC entry 210 (class 1259 OID 16401)
-- Name: layout_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.layout_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.layout_id_seq OWNER TO admin;

--
-- TOC entry 3366 (class 0 OID 0)
-- Dependencies: 210
-- Name: layout_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.layout_id_seq OWNED BY public.layouts.id;


--
-- TOC entry 215 (class 1259 OID 24598)
-- Name: rows; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.rows (
    id integer NOT NULL,
    section_id integer NOT NULL,
    number integer NOT NULL
);


ALTER TABLE public.rows OWNER TO admin;

--
-- TOC entry 214 (class 1259 OID 24597)
-- Name: row_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.row_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.row_id_seq OWNER TO admin;

--
-- TOC entry 3367 (class 0 OID 0)
-- Dependencies: 214
-- Name: row_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.row_id_seq OWNED BY public.rows.id;


--
-- TOC entry 218 (class 1259 OID 24624)
-- Name: seats; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    seat_index integer NOT NULL,
    seat_number integer NOT NULL,
    rank_id integer NOT NULL,
    row_id integer NOT NULL,
    reservation_id integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.seats OWNER TO admin;

--
-- TOC entry 219 (class 1259 OID 24627)
-- Name: seat_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.seat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seat_id_seq OWNER TO admin;

--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 219
-- Name: seat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.seat_id_seq OWNED BY public.seats.id;


--
-- TOC entry 212 (class 1259 OID 24578)
-- Name: sections; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sections (
    id integer NOT NULL,
    layout_id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.sections OWNER TO admin;

--
-- TOC entry 211 (class 1259 OID 24577)
-- Name: section_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.section_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.section_id_seq OWNER TO admin;

--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 211
-- Name: section_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.section_id_seq OWNED BY public.sections.id;


--
-- TOC entry 213 (class 1259 OID 24584)
-- Name: section_layoutid_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.section_layoutid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.section_layoutid_seq OWNER TO admin;

--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 213
-- Name: section_layoutid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.section_layoutid_seq OWNED BY public.sections.layout_id;


--
-- TOC entry 3188 (class 2604 OID 16402)
-- Name: layouts id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.layouts ALTER COLUMN id SET DEFAULT nextval('public.layout_id_seq'::regclass);


--
-- TOC entry 3192 (class 2604 OID 24615)
-- Name: ranks id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ranks ALTER COLUMN id SET DEFAULT nextval('public."Rank_id_seq"'::regclass);


--
-- TOC entry 3191 (class 2604 OID 24601)
-- Name: rows id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rows ALTER COLUMN id SET DEFAULT nextval('public.row_id_seq'::regclass);


--
-- TOC entry 3193 (class 2604 OID 24628)
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seat_id_seq'::regclass);


--
-- TOC entry 3189 (class 2604 OID 24581)
-- Name: sections id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sections ALTER COLUMN id SET DEFAULT nextval('public.section_id_seq'::regclass);


--
-- TOC entry 3190 (class 2604 OID 24585)
-- Name: sections layout_id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sections ALTER COLUMN layout_id SET DEFAULT nextval('public.section_layoutid_seq'::regclass);


--
-- TOC entry 3348 (class 0 OID 16386)
-- Dependencies: 209
-- Data for Name: layouts; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.layouts VALUES (1, 'Gas concert');
INSERT INTO public.layouts VALUES (2, 'Bass concert');
INSERT INTO public.layouts VALUES (3, 'Drake concert');


--
-- TOC entry 3355 (class 0 OID 24611)
-- Dependencies: 216
-- Data for Name: ranks; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.ranks VALUES (2, 15.99);
INSERT INTO public.ranks VALUES (1, 10.99);


--
-- TOC entry 3354 (class 0 OID 24598)
-- Dependencies: 215
-- Data for Name: rows; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.rows VALUES (1, 1, 1);
INSERT INTO public.rows VALUES (2, 1, 2);
INSERT INTO public.rows VALUES (3, 1, 3);


--
-- TOC entry 3357 (class 0 OID 24624)
-- Dependencies: 218
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.seats VALUES (10, 2, 3, 1, 2, 0);
INSERT INTO public.seats VALUES (11, 3, 5, 1, 2, 0);
INSERT INTO public.seats VALUES (12, 4, 6, 1, 2, 0);
INSERT INTO public.seats VALUES (13, 5, 2, 1, 2, 0);
INSERT INTO public.seats VALUES (5, 5, 2, 1, 1, 4);
INSERT INTO public.seats VALUES (6, 6, 8, 1, 1, 4);
INSERT INTO public.seats VALUES (1, 1, 1, 1, 1, 1);
INSERT INTO public.seats VALUES (2, 2, 3, 1, 1, 2);
INSERT INTO public.seats VALUES (3, 3, 5, 1, 1, 2);
INSERT INTO public.seats VALUES (4, 4, 6, 1, 1, 3);
INSERT INTO public.seats VALUES (16, 8, 7, 1, 2, 6);
INSERT INTO public.seats VALUES (15, 7, 4, 1, 2, 6);
INSERT INTO public.seats VALUES (14, 6, 8, 1, 2, 6);
INSERT INTO public.seats VALUES (17, 1, 1, 2, 3, 0);
INSERT INTO public.seats VALUES (18, 2, 3, 2, 3, 0);
INSERT INTO public.seats VALUES (19, 3, 5, 2, 3, 0);
INSERT INTO public.seats VALUES (20, 4, 6, 2, 3, 0);
INSERT INTO public.seats VALUES (21, 5, 2, 2, 3, 0);
INSERT INTO public.seats VALUES (22, 6, 8, 2, 3, 0);
INSERT INTO public.seats VALUES (23, 7, 4, 2, 3, 0);
INSERT INTO public.seats VALUES (24, 8, 7, 2, 3, 0);
INSERT INTO public.seats VALUES (9, 1, 1, 1, 2, 0);
INSERT INTO public.seats VALUES (7, 7, 4, 1, 1, 5);
INSERT INTO public.seats VALUES (8, 8, 7, 1, 1, 5);


--
-- TOC entry 3351 (class 0 OID 24578)
-- Dependencies: 212
-- Data for Name: sections; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.sections VALUES (1, 1, '101');
INSERT INTO public.sections VALUES (2, 1, '102');


--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 217
-- Name: Rank_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public."Rank_id_seq"', 2, true);


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 210
-- Name: layout_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.layout_id_seq', 7, true);


--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 214
-- Name: row_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.row_id_seq', 3, true);


--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 219
-- Name: seat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.seat_id_seq', 25, true);


--
-- TOC entry 3375 (class 0 OID 0)
-- Dependencies: 211
-- Name: section_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.section_id_seq', 2, true);


--
-- TOC entry 3376 (class 0 OID 0)
-- Dependencies: 213
-- Name: section_layoutid_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.section_layoutid_seq', 1, false);


--
-- TOC entry 3196 (class 2606 OID 16407)
-- Name: layouts layout_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.layouts
    ADD CONSTRAINT layout_pk PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 24633)
-- Name: ranks rank_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ranks
    ADD CONSTRAINT rank_pk PRIMARY KEY (id);


--
-- TOC entry 3200 (class 2606 OID 24603)
-- Name: rows row_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rows
    ADD CONSTRAINT row_pk PRIMARY KEY (id);


--
-- TOC entry 3204 (class 2606 OID 24635)
-- Name: seats seat_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seat_pk PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 24583)
-- Name: sections section_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT section_pk PRIMARY KEY (id);


--
-- TOC entry 3206 (class 2606 OID 24604)
-- Name: rows row_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rows
    ADD CONSTRAINT row_fk FOREIGN KEY (section_id) REFERENCES public.rows(id);


--
-- TOC entry 3207 (class 2606 OID 24636)
-- Name: seats seat_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seat_fk FOREIGN KEY (rank_id) REFERENCES public.ranks(id);


--
-- TOC entry 3208 (class 2606 OID 24641)
-- Name: seats seat_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seat_fk_1 FOREIGN KEY (row_id) REFERENCES public.rows(id);


--
-- TOC entry 3205 (class 2606 OID 24590)
-- Name: sections section_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT section_fk FOREIGN KEY (layout_id) REFERENCES public.layouts(id);


-- Completed on 2022-01-08 01:51:40

--
-- PostgreSQL database dump complete
--

