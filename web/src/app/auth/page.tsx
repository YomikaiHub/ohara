"use client";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { LogoAvatar } from "@/components/ui/avatar";
import { useQueryState } from "nuqs";
import RegisterForm from "@/features/auth/components/register-form";
import AuthShowcase from "@/features/auth/components/showcase/auth-showcase";
import LoginForm from "@/features/auth/components/login-form";

function AuthPage() {
    const [mode, setMode] = useQueryState("mode", {
        defaultValue: "register",
        clearOnDefault: false,
    });

    return (
        <main className="flex min-h-screen">
            <div className="flex flex-1 flex-col items-center justify-center gap-10">
                <div className="flex flex-col items-center">
                    <LogoAvatar className="mb-4 size-14" />
                    <h1 className="text-2xl">Join. Explore.</h1>
                    <p className="text-lg">
                        Start your next chapter with{" "}
                        <b className="text-teal-300">Ohara</b>.
                    </p>
                </div>
                <div className="w-full">
                    <Tabs
                        value={mode}
                        onValueChange={(value) => setMode(value)}
                        className="mx-auto w-120"
                    >
                        <TabsList className="mx-auto flex w-2/3 flex-row rounded-full *:rounded-full">
                            <TabsTrigger value="register">Register</TabsTrigger>
                            <TabsTrigger value="login">Login</TabsTrigger>
                        </TabsList>
                        <TabsContent value="register" className="my-8 px-4">
                            <RegisterForm />
                        </TabsContent>
                        <TabsContent value="login" className="my-8 px-4">
                            <LoginForm />
                        </TabsContent>
                    </Tabs>
                </div>
            </div>
            <div className="flex-1">
                <AuthShowcase />
            </div>
        </main>
    );
}

export default AuthPage;
