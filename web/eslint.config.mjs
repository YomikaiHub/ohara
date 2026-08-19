import tanstackQuery from "@tanstack/eslint-plugin-query";

export default [
    {
        plugins: {
            "@tanstack/query": tanstackQuery,
        },
        rules: {
            ...tanstackQuery.configs.recommended.rules,
        },
    },
];
