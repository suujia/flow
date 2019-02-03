CREATE DATABASE spots;

CREATE TABLE "spot" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"location" text NOT NULL,
	"tags" text REFERENCES tags (id)
);

CREATE TABLE "tagcount" (
    "id" text PRIMARY KEY,
	"count" integer NOT NULL,
	"tag" text NOT NULL,
	"spot" text NOT NULL,
	FOREIGN KEY (tag) REFERENCES tags(id) ON UPDATE CASCADE,
	FOREIGN KEY (spot) REFERENCES spot(id) ON UPDATE CASCADE
);

CREATE TABLE "tags" (
    "id" text PRIMARY KEY,
	"colour" text NOT NULL,
	"tag" text NOT NULL,
);

CREATE TABLE "user" (
	"id" text PRIMARY KEY,
	"username" text NOT NULL,
	"city" text NOT NULL,
	"favourites" text NOT NULL,
);