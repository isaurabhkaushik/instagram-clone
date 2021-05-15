CREATE TYPE user_language AS ENUM('HINDI', 'ENGLISH', 'GARHWALI');
CREATE TYPE gender AS ENUM('MALE', 'FEMALE', 'PREFER_NOT_TO_SAY');
CREATE TYPE post_type AS ENUM('VIDEO', 'TEXT', 'IMAGE', 'HYBRID');
CREATE TYPE parent_type AS ENUM('POST', 'COMMENT');
CREATE TYPE category AS ENUM('SPORTS', 'NEWS', 'LIFESTYLE', 'INFOTAINMENT', 'ENTERTAINMENT', 'TECHNOLOGY', 'SPIRITUAL');

CREATE TABLE user_profile (
      id varchar(50) primary key,
      is_active bool default false,
      is_completed bool default false,
      create_time timestamp with time zone,
      update_time timestamp with time zone,
      category_id varchar(50),
      first_name varchar(50),
      last_name varchar(50),
      language user_language,
      gender gender,
      address text,
      mobile varchar(50),
      email varchar(50),
      username varchar(50),
      password varchar(50),
      labels jsonb
);

CREATE TABLE posts (
       id varchar(50) primary key,
       is_active bool default false,
       create_time timestamp with time zone,
       update_time timestamp with time zone,
       type post_type,
       category category,
       creator_id varchar(50),
       heading varchar(50),
       body text,
       resource_urls jsonb,
       labels jsonb
);

CREATE TABLE comments (
      id varchar(50) primary key,
      is_active bool default false,
      create_time timestamp with time zone,
      update_time timestamp with time zone,
      parent_id varchar(50),
      creator_id varchar(50),
      body text
);

CREATE TABLE likes (
       id varchar(50) primary key,
       is_active bool default false,
       create_time timestamp with time zone,
       update_time timestamp with time zone,
       parent_id varchar(50),
       creator_id varchar(50),
);

-- CREATE TABLE activity (
--       id varchar(50) primary key,
--       is_active bool default false,
--       create_time timestamp with time zone,
--       update_time timestamp with time zone,
--       parent_id varchar(50),
--       parent_type parent_type,
--       total_likes bigint,
--       total_comments bigint,
--       labels jsonb
-- )

-- CREATE TABLE user_category (
--        id varchar(50) primary key,
--        user_id varchar(50),
--        categories jsonb
-- )