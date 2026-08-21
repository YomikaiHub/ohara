import type { ReactNode } from "react";
import { Field, FieldError, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";

interface FormFieldProps {
    field: any;
    label: string;
    type?: React.HTMLInputTypeAttribute;
    autoComplete?: string;
    endAdornment?: ReactNode;
}

export function FormField({
    field,
    label,
    type = "text",
    autoComplete = "off",
    endAdornment,
}: FormFieldProps) {
    const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

    return (
        <Field data-invalid={isInvalid}>
            <div className="flex w-full items-center justify-between">
                <FieldLabel htmlFor={field.name}>{label}</FieldLabel>

                {endAdornment}
            </div>

            <Input
                id={field.name}
                name={field.name}
                type={type}
                value={field.state.value}
                onBlur={field.handleBlur}
                onChange={(e) => field.handleChange(e.target.value)}
                aria-invalid={isInvalid}
                autoComplete={autoComplete}
            />

            {isInvalid && <FieldError errors={field.state.meta.errors} />}
        </Field>
    );
}
