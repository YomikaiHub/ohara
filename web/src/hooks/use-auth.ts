import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { loginEmail, logout, getMe } from "@/lib/auth-api";
import { setAccessToken } from "@/lib/auth-store";

export function useLogin() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: loginEmail,
        onSuccess: (data) => {
            setAccessToken(data.accessToken);
            qc.invalidateQueries({ queryKey: ["me"] });
        },
    });
}

export function useLogout() {
    const qc = useQueryClient();
    return useMutation({
        mutationFn: logout,
        onSuccess: () => {
            setAccessToken(null);
            qc.clear();
        },
    });
}

export function useMe() {
    return useQuery({
        queryKey: ["me"],
        queryFn: getMe,
        retry: false,
    });
}
