import { defineRelations } from "drizzle-orm";
import { authors } from "./author";
import { books } from "./book";

export const relations = defineRelations({ authors, books }, (r) => ({
    authors: {
        books: r.many.books(),
    },

    books: {
        author: r.one.authors({
            from: r.books.authorId,
            to: r.authors.id,
        }),
    },
}));
