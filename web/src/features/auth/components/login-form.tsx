import { useForm } from "@tanstack/react-form";
import { registerSchema } from "../schemas/register-schema";
import { useState, type SubmitEvent } from "react";
import {
    Field,
    FieldError,
    FieldGroup,
    FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Eye, EyeClosed } from "lucide-react";
import { Button } from "@/components/ui/button";
import { FormField } from "./form-field";
import { toast } from "sonner";
import { loginSchema } from "../schemas/login-schema";
import { useLogin } from "@/hooks/use-auth";
import { useRouter } from "next/navigation";

function LoginForm() {
    const [viewPassword, setViewPassword] = useState(false);
    const login = useLogin();
    const router = useRouter();

    const form = useForm({
        defaultValues: {
            email: "",
            password: "",
        },
        validators: { onSubmit: loginSchema },
        onSubmit: async ({ value }) => {
            try {
                await login.mutateAsync(value);

                toast.success("Welcome back!");
                router.replace("/dashboard");
            } catch (error) {
                toast.error("Unable to log in. Please check your credentials.");
            }
        },
    });

    function handleSubmit(e: SubmitEvent<HTMLFormElement>) {
        e.preventDefault();
        form.handleSubmit();
    }

    return (
        <div>
            <div>
                <form id="ohara-login-form" onSubmit={handleSubmit}>
                    <FieldGroup>
                        <form.Field
                            name="email"
                            children={(field) => (
                                <FormField
                                    field={field}
                                    type="email"
                                    label="Email"
                                    autoComplete="email"
                                />
                            )}
                        />
                        <form.Field
                            name="password"
                            children={(field) => (
                                <FormField
                                    field={field}
                                    label="Password"
                                    type={viewPassword ? "text" : "password"}
                                    autoComplete="password"
                                    endAdornment={
                                        <button
                                            type="button"
                                            onClick={() =>
                                                setViewPassword((prev) => !prev)
                                            }
                                            aria-label={
                                                viewPassword
                                                    ? "Hide password"
                                                    : "Show password"
                                            }
                                        >
                                            {viewPassword ? (
                                                <EyeClosed className="size-4" />
                                            ) : (
                                                <Eye className="size-4" />
                                            )}
                                        </button>
                                    }
                                />
                            )}
                        />
                    </FieldGroup>
                </form>
                <div className="my-8">
                    <Field orientation="horizontal" className="justify-between">
                        <Button
                            size="lg"
                            type="button"
                            variant="outline"
                            onClick={() => form.reset()}
                        >
                            Reset
                        </Button>
                        <Button size="lg" type="submit" form="ohara-login-form">
                            Register Now
                        </Button>
                    </Field>
                </div>
            </div>
        </div>
    );
}

export default LoginForm;
