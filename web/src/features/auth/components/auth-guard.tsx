"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useMe } from "@/hooks/use-auth";

interface AuthGuardProps {
    children: React.ReactNode;
}

function AuthGuard({ children }: AuthGuardProps) {
    const router = useRouter();
    const { data: user, isLoading, isError } = useMe();

    useEffect(() => {
        if (isError) {
            router.replace("/auth?mode=login");
        }
    }, [isError, router]);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (isError || !user) {
        return null;
    }

    return <>{children}</>;
}

export default AuthGuard;
