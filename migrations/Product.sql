/*
 Navicat Premium Data Transfer

 Source Server         : Con Voi Database
 Source Server Type    : PostgreSQL
 Source Server Version : 160002 (160002)
 Source Host           : localhost:5432
 Source Catalog        : backend-tit
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160002 (160002)
 File Encoding         : 65001

 Date: 02/05/2024 10:25:42
*/


-- ----------------------------
-- Table structure for Product
-- ----------------------------
DROP TABLE IF EXISTS "public"."Product";
CREATE TABLE "public"."Product" (
  "product_id" int4 NOT NULL,
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "slug" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "brand" varchar(255) COLLATE "pg_catalog"."default",
  "price" numeric(10,2) NOT NULL,
  "sale_price" numeric(10,2),
  "category_id" int4,
  "quantity" int4 DEFAULT 0,
  "sold" int4 DEFAULT 0,
  "images" text COLLATE "pg_catalog"."default",
  "color" varchar(40) COLLATE "pg_catalog"."default",
  "rating" numeric(10,2),
  "created_at" timestamp(6),
  "updated_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."Product"."sold" IS 'solved (đã bán)';

-- ----------------------------
-- Primary Key structure for table Product
-- ----------------------------
ALTER TABLE "public"."Product" ADD CONSTRAINT "Product_pkey" PRIMARY KEY ("product_id");
