CREATE TABLE `authors` (
	`id` text PRIMARY KEY DEFAULT lower(hex(randomblob(16))) NOT NULL,
	`name` text NOT NULL,
	`bio` text,
	`avatar` text,
	`created_at` integer NOT NULL,
	`updated_at` integer NOT NULL
);
--> statement-breakpoint
CREATE TABLE `books` (
	`id` text PRIMARY KEY DEFAULT lower(hex(randomblob(16))) NOT NULL,
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
