import { relations } from "drizzle-orm";
import { authors } from "./author";
import { books } from "./book";

export const authorsRelations = relations(authors, ({ many }) => ({
    books: many(books),
}));

export const bookRelations = relations(books, ({ one }) => ({
    author: one(authors, {
        fields: [books.authorId],
        references: [authors.id],
    }),
}));
