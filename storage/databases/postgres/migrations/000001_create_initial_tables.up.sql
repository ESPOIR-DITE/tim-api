
CREATE TABLE IF NOT EXISTS "channel" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar,
  "channel_type_id" varchar,
  "account_id" varchar,
  "region" varchar,
  "date" varchar,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "channel_type" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar,
  "description" varchar,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "channel_subscription" (
  "channel_id" varchar  NOT NULL,
  "user_id" varchar  NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY (channel_id, user_id)
);

CREATE TABLE IF NOT EXISTS "channel_video" (
  "channel_id" varchar  NOT NULL,
  "video_id" varchar  NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY (channel_id, video_id)
);





CREATE TABLE IF NOT EXISTS "role" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
  "email" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar,
  "role_id" varchar,
  "birth_date" timestamptz,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_account" (
  "id" varchar PRIMARY KEY NOT NULL,
  "account_id" varchar NOT NULL,
  "user_detail_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "account" (
  "id" varchar PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL unique ,
  "password" varchar NOT NULL,
  "status" boolean ,
  "token" varchar ,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_detail" (
  "id" varchar PRIMARY KEY NOT NULL,
  "bank_id" varchar NOT NULL,
  "company_registered_number" varchar NOT NULL,
  "tax_number" varchar NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "user_bank" (
  "id" varchar PRIMARY KEY NOT NULL,
  "bank_type" varchar NOT NULL,
  "bank_name" varchar NOT NULL,
  "branch_code" varchar NOT NULL,
  "bank_number" varchar NOT NULL,
  "cvc_code" varchar NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "user_subscription" (
  "id" varchar PRIMARY KEY NOT NULL,
  "account_id" varchar NOT NULL,
  "stat" varchar NOT NULL,
  "subscription_id" varchar NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "user_video" (
  "id" varchar PRIMARY KEY NOT NULL,
  "account_id" varchar NOT NULL,
  "video_id" varchar NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);


-- Video class
CREATE TABLE IF NOT EXISTS "category" (
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "video_category" (
  "id" varchar PRIMARY KEY NOT NULL,
  "video_id" varchar NOT NULL,
  "category_id" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "video_comment" (
  "id" varchar PRIMARY KEY NOT NULL,
  "video_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "comment" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "video_data" (
  "id" varchar PRIMARY KEY NOT NULL,
  "file_type" varchar NOT NULL,
  "file_size" varchar NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "video_reaction" (
  "video_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "like" integer NOT NULL,
  "unlike" integer NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  PRIMARY KEY (video_id, user_id)
);

CREATE TABLE IF NOT EXISTS "video" (
  "id" varchar PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "url" varchar NOT NULL,
  "date" timestamptz NOT NULL,
  "price" float,
  "is_private" boolean,
  "date_uploaded" timestamptz NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

ALTER TABLE "video_reaction" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");
ALTER TABLE "video_reaction" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("email");

ALTER TABLE "video_comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("email");
ALTER TABLE "video_comment" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");


ALTER TABLE "video_category" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");
ALTER TABLE "video_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");




ALTER TABLE "channel" ADD FOREIGN KEY ("channel_type_id") REFERENCES "channel_type" ("id");
ALTER TABLE "channel" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "channel_subscription" ADD FOREIGN KEY ("user_id") REFERENCES "account" ("id");
ALTER TABLE "channel_subscription" ADD FOREIGN KEY ("channel_id") REFERENCES "channel" ("id");


ALTER TABLE "channel_video" ADD FOREIGN KEY ("channel_id") REFERENCES "channel" ("id");
ALTER TABLE "channel_video" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");


ALTER TABLE "user_subscription" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "user_video" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");
ALTER TABLE "user_video" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "user_account" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "user_account" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("email");

-- ALTER TABLE "user_v" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("email");

ALTER TABLE "user_account" ADD FOREIGN KEY ("user_detail_id") REFERENCES "user_detail" ("id");
