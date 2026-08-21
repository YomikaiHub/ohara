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

function RegisterForm() {
    const [viewPassword, setViewPassword] = useState(false);
    const form = useForm({
        defaultValues: {
            first_name: "",
            last_name: "",
            username: "",
            email: "",
            password: "",
        },
        validators: { onSubmit: registerSchema },
        onSubmit: async ({ value }) => {
            toast("You submitted the following values:", {
                description: (
                    <pre className="bg-code text-code-foreground mt-2 w-[320px] overflow-x-auto rounded-md p-4">
                        <code>{JSON.stringify(value, null, 2)}</code>
                    </pre>
                ),
                position: "bottom-right",
                classNames: {
                    content: "flex flex-col gap-2",
                },
                style: {
                    "--border-radius": "calc(var(--radius)  + 4px)",
                } as React.CSSProperties,
            });
        },
    });

    function handleSubmit(e: SubmitEvent<HTMLFormElement>) {
        e.preventDefault();
        form.handleSubmit();
    }

    return (
        <div>
            <div>
                <form id="ohara-register-form" onSubmit={handleSubmit}>
                    <FieldGroup>
                        <form.Field
                            name="first_name"
                            children={(field) => (
                                <FormField field={field} label="First Name" />
                            )}
                        />
                        <form.Field
                            name="last_name"
                            children={(field) => (
                                <FormField field={field} label="Last Name" />
                            )}
                        />
                        <form.Field
                            name="username"
                            children={(field) => (
                                <FormField field={field} label="Username" />
                            )}
                        />
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
                                    autoComplete="new-password"
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
                        <Button
                            size="lg"
                            type="submit"
                            form="ohara-register-form"
                        >
                            Register Now
                        </Button>
                    </Field>
                </div>
            </div>
        </div>
    );
}

export default RegisterForm;
