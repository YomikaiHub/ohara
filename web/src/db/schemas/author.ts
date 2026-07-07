import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";

export const authors = pgTable("authors", {
    id: uuid("id").primaryKey().defaultRandom(),
    name: text("name").notNull(),
    bio: text("bio"),
    avatar: text("avatar"),
    createdAt: timestamp("created_at").notNull().defaultNow(),
    updatedAt: timestamp("updated_at").notNull().defaultNow(),
});
