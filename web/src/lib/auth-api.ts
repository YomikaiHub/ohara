import axios from "axios";
import { api } from "./api-client";
import { User } from "@/types/user";

export const healthCheck = () => api.get("/health").then((r) => r.data);

export const registerEmail = (payload: { email: string; password: string }) =>
    api.post("/auth/register/email", payload).then((r) => r.data);

export const loginEmail = (payload: { email: string; password: string }) =>
    api.post("/auth/login/email", payload).then((r) => r.data); // { accessToken, user }

export const logout = () => api.post("/auth/logout").then((r) => r.data);

export const getMe = () => api.get<User>("/users/me").then((r) => r.data);

export const refreshToken = () =>
    axios
        .post(
            "http://localhost:8000/api/v1/auth/refresh",
            {},
            { withCredentials: true }
        )
        .then((r) => r.data);
