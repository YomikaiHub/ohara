import { z } from "zod";

export const registerSchema = z.object({
    first_name: z
        .string()
        .min(1, "First name is required")
        .max(50, "First name must be less than 50 characters"),

    last_name: z
        .string()
        .min(1, "Last name is required")
        .max(50, "Last name must be less than 50 characters"),

    username: z
        .string()
        .min(3, "Username must be at least 3 characters")
        .max(30, "Username must be less than 30 characters"),

    email: z.email("Please enter a valid email address"),

    password: z
        .string()
        .min(8, "Password must be at least 8 characters")
        .regex(
            /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z0-9]).{9,}$/,
            "Password must contain uppercase, lowercase, number, and special character"
        ),
});

export type RegisterInput = z.infer<typeof registerSchema>;
