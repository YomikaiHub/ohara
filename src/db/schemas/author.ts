import { randomUUID } from "node:crypto";
import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const authors = sqliteTable("authors", {
    id: text("id")
        .primaryKey()
        .$defaultFn(() => randomUUID()),
    name: text("name").notNull(),
    bio: text("bio"),
    avatar: text("avatar"),
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
