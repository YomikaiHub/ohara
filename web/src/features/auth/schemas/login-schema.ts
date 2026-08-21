import { z } from "zod";

export const loginSchema = z.object({
    email: z.email("Please enter a valid email address"),

    password: z
        .string()
        .min(8, "Password must be at least 8 characters")
        .regex(
            /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z0-9]).{9,}$/,
            "Password must contain uppercase, lowercase, number, and special character"
        ),
});

export type LoginInput = z.infer<typeof loginSchema>;
