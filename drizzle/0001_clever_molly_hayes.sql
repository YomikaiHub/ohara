PRAGMA foreign_keys=OFF;--> statement-breakpoint
CREATE TABLE `__new_authors` (
	`id` text PRIMARY KEY NOT NULL,
	`name` text NOT NULL,
	`bio` text,
	`avatar` text,
	`created_at` integer NOT NULL,
	`updated_at` integer NOT NULL
);
--> statement-breakpoint
INSERT INTO `__new_authors`("id", "name", "bio", "avatar", "created_at", "updated_at") SELECT "id", "name", "bio", "avatar", "created_at", "updated_at" FROM `authors`;--> statement-breakpoint
DROP TABLE `authors`;--> statement-breakpoint
ALTER TABLE `__new_authors` RENAME TO `authors`;--> statement-breakpoint
PRAGMA foreign_keys=ON;--> statement-breakpoint
CREATE TABLE `__new_books` (
	`id` text PRIMARY KEY NOT NULL,
	`title` text NOT NULL,
	`author_id` text NOT NULL,
	`description` text,
	`cover_url` text,
	`publisher` text,
	`published_date` text,
	`isbn` text,
	`page_count` integer,
	`language` text,
	`source` text DEFAULT 'kindle' NOT NULL,
	`imported_at` integer NOT NULL,
	`created_at` integer NOT NULL,
	`updated_at` integer NOT NULL
);
--> statement-breakpoint
INSERT INTO `__new_books`("id", "title", "author_id", "description", "cover_url", "publisher", "published_date", "isbn", "page_count", "language", "source", "imported_at", "created_at", "updated_at") SELECT "id", "title", "author_id", "description", "cover_url", "publisher", "published_date", "isbn", "page_count", "language", "source", "imported_at", "created_at", "updated_at" FROM `books`;--> statement-breakpoint
DROP TABLE `books`;--> statement-breakpoint
ALTER TABLE `__new_books` RENAME TO `books`;