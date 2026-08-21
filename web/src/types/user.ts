export interface User {
    id: string;
    first_name: string;
    last_name: string;
    username: string;
    email: string;
    image?: string | null;
    email_verified: boolean;
    created_at: string;
    updated_at: string;
}
