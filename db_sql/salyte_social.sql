--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

-- Started on 2025-03-04 14:32:02

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 5010 (class 1262 OID 24720)
-- Name: salyte_social; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE salyte_social IF EXISTS WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';


ALTER DATABASE salyte_social OWNER TO postgres;

\connect salyte_social

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 24727)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 5011 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 227 (class 1259 OID 24963)
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    comment_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    post_id uuid NOT NULL,
    content text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 24899)
-- Name: communities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.communities (
    community_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    profile_picture_url character varying(255),
    visibility character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.communities OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 24909)
-- Name: communitymembers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.communitymembers (
    member_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    community_id uuid NOT NULL,
    profile_id uuid NOT NULL,
    role character varying(50) NOT NULL,
    joined_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.communitymembers OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 24839)
-- Name: followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.followers (
    follower_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    follower_profile_id uuid NOT NULL,
    followed_profile_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.followers OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 24856)
-- Name: interchange; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.interchange (
    interchange_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    friend_profile_id uuid NOT NULL,
    status character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.interchange OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 24983)
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    like_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    post_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 24926)
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    post_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    community_id uuid,
    content text NOT NULL,
    media_url character varying(255),
    visibility character varying(50) DEFAULT 'visible'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 24829)
-- Name: profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profile (
    profile_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    username character varying(255) NOT NULL,
    bio text,
    profile_picture_url character varying(255),
    visibility character varying(50) DEFAULT 'visible'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.profile OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 24887)
-- Name: restrictions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.restrictions (
    restrictions_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    type character varying(50) NOT NULL,
    source_id uuid NOT NULL,
    is_read boolean NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.restrictions OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 24874)
-- Name: settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.settings (
    setting_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    dark_mode_enabled boolean NOT NULL,
    language character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.settings OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 24946)
-- Name: share; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.share (
    share_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    post_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.share OWNER TO postgres;

--
-- TOC entry 4841 (class 2606 OID 24972)
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (comment_id);


--
-- TOC entry 4833 (class 2606 OID 24908)
-- Name: communities communities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT communities_pkey PRIMARY KEY (community_id);


--
-- TOC entry 4835 (class 2606 OID 24915)
-- Name: communitymembers communitymembers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.communitymembers
    ADD CONSTRAINT communitymembers_pkey PRIMARY KEY (member_id);


--
-- TOC entry 4825 (class 2606 OID 24845)
-- Name: followers followers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.followers
    ADD CONSTRAINT followers_pkey PRIMARY KEY (follower_id);


--
-- TOC entry 4827 (class 2606 OID 24863)
-- Name: interchange interchange_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interchange
    ADD CONSTRAINT interchange_pkey PRIMARY KEY (interchange_id);


--
-- TOC entry 4843 (class 2606 OID 24989)
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (like_id);


--
-- TOC entry 4837 (class 2606 OID 24935)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (post_id);


--
-- TOC entry 4823 (class 2606 OID 24838)
-- Name: profile profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (profile_id);


--
-- TOC entry 4831 (class 2606 OID 24893)
-- Name: restrictions restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restrictions
    ADD CONSTRAINT restrictions_pkey PRIMARY KEY (restrictions_id);


--
-- TOC entry 4829 (class 2606 OID 24881)
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (setting_id);


--
-- TOC entry 4839 (class 2606 OID 24952)
-- Name: share share_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.share
    ADD CONSTRAINT share_pkey PRIMARY KEY (share_id);


--
-- TOC entry 4856 (class 2606 OID 24978)
-- Name: comments comments_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON DELETE CASCADE;


--
-- TOC entry 4857 (class 2606 OID 24973)
-- Name: comments comments_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4850 (class 2606 OID 24916)
-- Name: communitymembers communitymembers_community_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.communitymembers
    ADD CONSTRAINT communitymembers_community_id_fkey FOREIGN KEY (community_id) REFERENCES public.communities(community_id) ON DELETE CASCADE;


--
-- TOC entry 4851 (class 2606 OID 24921)
-- Name: communitymembers communitymembers_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.communitymembers
    ADD CONSTRAINT communitymembers_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4844 (class 2606 OID 24851)
-- Name: followers followers_followed_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.followers
    ADD CONSTRAINT followers_followed_profile_id_fkey FOREIGN KEY (followed_profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4845 (class 2606 OID 24846)
-- Name: followers followers_follower_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.followers
    ADD CONSTRAINT followers_follower_profile_id_fkey FOREIGN KEY (follower_profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4846 (class 2606 OID 24869)
-- Name: interchange interchange_friend_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interchange
    ADD CONSTRAINT interchange_friend_profile_id_fkey FOREIGN KEY (friend_profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4847 (class 2606 OID 24864)
-- Name: interchange interchange_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interchange
    ADD CONSTRAINT interchange_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4858 (class 2606 OID 24995)
-- Name: likes likes_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON DELETE CASCADE;


--
-- TOC entry 4859 (class 2606 OID 24990)
-- Name: likes likes_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4852 (class 2606 OID 24941)
-- Name: posts posts_community_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_community_id_fkey FOREIGN KEY (community_id) REFERENCES public.communities(community_id) ON DELETE SET NULL;


--
-- TOC entry 4853 (class 2606 OID 24936)
-- Name: posts posts_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4849 (class 2606 OID 24894)
-- Name: restrictions restrictions_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restrictions
    ADD CONSTRAINT restrictions_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4848 (class 2606 OID 24882)
-- Name: settings settings_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


--
-- TOC entry 4854 (class 2606 OID 24958)
-- Name: share share_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.share
    ADD CONSTRAINT share_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON DELETE CASCADE;


--
-- TOC entry 4855 (class 2606 OID 24953)
-- Name: share share_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.share
    ADD CONSTRAINT share_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profile(profile_id) ON DELETE CASCADE;


-- Completed on 2025-03-04 14:32:02

--
-- PostgreSQL database dump complete
--

