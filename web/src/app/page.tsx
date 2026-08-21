"use client";

import { healthCheck } from "@/lib/auth-api";
import { useQuery } from "@tanstack/react-query";

export default function Home() {
    const { data, isLoading, isError, error } = useQuery({
        queryKey: ["health"],
        queryFn: healthCheck,
    });

    if (isLoading) {
        return <p>Checking API...</p>;
    }

    if (isError) {
        return <p>Error: {error.message}</p>;
    }

    return (
        <div>
            <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
    );
}
