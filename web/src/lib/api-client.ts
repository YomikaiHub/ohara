import axios from "axios";
import { getAccessToken, setAccessToken } from "./auth-store";

export const api = axios.create({
    baseURL: "http://localhost:8000/api/v1",
    withCredentials: true,
});

api.interceptors.request.use((config) => {
    const token = getAccessToken();
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

let isRefreshing = false;
let queue: {
    resolve: (t: string) => void;
    reject: (e: unknown) => void;
}[] = [];

const processQueue = (error: unknown, token: string | null = null) => {
    queue.forEach(({ resolve, reject }) => {
        if (error) reject(error);
        else resolve(token!);
    });
    queue = [];
};

api.interceptors.response.use(
    (res) => res,
    async (error) => {
        const originalRequest = error.config;

        if (error.response?.status !== 401 || originalRequest._retry) {
            return Promise.reject(error);
        }

        if (isRefreshing) {
            // queue this request until the in-flight refresh finishes
            return new Promise((resolve, reject) => {
                queue.push({ resolve, reject });
            }).then((token) => {
                originalRequest.headers.Authorization = `Bearer ${token}`;

                return api(originalRequest);
            });
        }

        originalRequest._retry = true;
        isRefreshing = true;

        try {
            const { data } = await axios.post(
                "http://localhost:8000/api/v1/auth/refresh",
                {},
                { withCredentials: true }
            );

            setAccessToken(data.accessToken);
            processQueue(null, data.accessToken);

            originalRequest.headers.Authorization = `Bearer ${data.accessToken}`;

            return api(originalRequest);
        } catch (refreshError) {
            processQueue(refreshError, null);
            setAccessToken(null);

            if (typeof window !== "undefined") {
                window.location.href = "/auth?mode=login";
            }

            return Promise.reject(refreshError);
        } finally {
            isRefreshing = false;
        }
    }
);
