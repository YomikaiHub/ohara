import { integer, pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";

export const books = pgTable("books", {
    id: uuid("id").primaryKey().defaultRandom(),
    title: text("title").notNull(),
    authorId: text("author_id").notNull(),
    description: text("description"),
    coverUrl: text("cover_url"),
    publisher: text("publisher"),
    publishedDate: text("published_date"),
    isbn: text("isbn"),
    pageCount: integer("page_count"),
    language: text("language"),
    source: text("source"),
    importedAt: timestamp("imported_at").notNull().defaultNow(),
    createdAt: timestamp("created_at").notNull().defaultNow(),
    updatedAt: timestamp("updated_at").notNull().defaultNow(),
});
