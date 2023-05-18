BEGIN;
SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = PUBLIC,
    EXTENSIONS;
SET default_tablespace = '';
SET default_with_oids = FALSE;
-- Extensions --
CREATE EXTENSION IF NOT EXISTS pgcrypto;
-- Tables -- 
-- Table: public.category 
CREATE TABLE IF NOT EXISTS public.category (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    parent_category_id INT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (parent_category_id) REFERENCES public.category (id)
);
-- Table public.image --
CREATE TABLE IF NOT EXISTS public.image (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    path TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Table: public.order_status --
CREATE TABLE IF NOT EXISTS public.order_status (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);
-- Table: public.ticket_status --
CREATE TABLE IF NOT EXISTS public.ticket_status (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);
-- Table: public.user --
CREATE TABLE IF NOT EXISTS public.user (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Table: public.currency --
CREATE TABLE IF NOT EXISTS public.currency (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL
);
-- Table: public.language --
CREATE TABLE IF NOT EXISTS public.language (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL
);
-- Table: public.country --
CREATE TABLE IF NOT EXISTS public.country (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    language_id INT NOT NULL,
    currency_id INT NOT NULL,
    FOREIGN KEY (currency_id) REFERENCES public.currency (id),
    FOREIGN KEY (language_id) REFERENCES public.language (id)
);
-- Table: public.address --
CREATE TABLE IF NOT EXISTS public.address (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    country TEXT NOT NULL,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    house_number TEXT NOT NULL,
    apartment_number TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.user (id)
);
-- Table: public.order --
CREATE TABLE IF NOT EXISTS public.order (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    address_id UUID NOT NULL,
    status_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.user (id),
    FOREIGN KEY (address_id) REFERENCES public.address (id),
    FOREIGN KEY (status_id) REFERENCES public.order_status (id)
);
-- Table: public.product -- 
CREATE TABLE IF NOT EXISTS public.product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    sku INT NOT NULL,
    specification JSONB,
    category_id INT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES public.category (id)
);
-- Table public.order_item -- 
CREATE TABLE IF NOT EXISTS public.order_item (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES public.order (id),
    FOREIGN KEY (product_id) REFERENCES public.product (id)
);
-- Table: public.review --
CREATE TABLE IF NOT EXISTS public.review (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    rating INT NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES public.product (id),
    FOREIGN KEY (order_id) REFERENCES public.order (id),
    FOREIGN KEY (user_id) REFERENCES public.user (id)
);
-- Table: public.product_image --
CREATE TABLE IF NOT EXISTS public.product_image (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    image_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES public.product (id),
    FOREIGN KEY (image_id) REFERENCES public.image (id)
);
-- Table: public.price --
CREATE TABLE IF NOT EXISTS public.price (
    id INT PRIMARY KEY,
    product_id UUID NOT NULL,
    currency_id INT NOT NULL,
    value NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES public.product (id),
    FOREIGN KEY (currency_id) REFERENCES public.currency (id)
);
--Table: public.ticket --
CREATE TABLE IF NOT EXISTS public.ticket (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    order_id UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.user (id),
    FOREIGN KEY (status_id) REFERENCES public.ticket_status (id),
    FOREIGN KEY (order_id) REFERENCES public.order (id)
);
END;