import { randomUUID } from "node:crypto";
import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const books = sqliteTable("books", {
    id: text("id")
        .primaryKey()
        .$defaultFn(() => randomUUID()),
    title: text("title").notNull(),
    authorId: text("author_id").notNull(),
    description: text("description"),
    coverUrl: text("cover_url"),
    publisher: text("publisher"),
    publishedDate: text("published_date"),
    isbn: text("isbn"),
    pageCount: integer("page_count"),
    language: text("language"),
    source: text("source").notNull().default("kindle"),
    importedAt: integer("imported_at", {
        mode: "timestamp",
    })
        .notNull()
        .$defaultFn(() => new Date()),
    createdAt: integer("created_at", {
        mode: "timestamp",
    })
        .notNull()
        .$defaultFn(() => new Date()),
    updatedAt: integer("updated_at", {
        mode: "timestamp",
    })
        .notNull()
        .$defaultFn(() => new Date()),
});
