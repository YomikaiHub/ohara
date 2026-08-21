import Image from "next/image";
import type { ReactNode } from "react";

interface BookShowcaseProps {
    src: string;
    alt: string;
    className?: string;
    imgWidth?: number;
    imgHeight?: number;
}

export function BookShowcase({
    src,
    alt = "",
    className = "",
    imgWidth = 300,
    imgHeight = 450,
}: BookShowcaseProps) {
    return (
        <div className={className}>
            <div className="relative">
                <div
                    className="absolute inset-[1%] translate-x-[4%] translate-y-[3%] rounded-sm bg-[#cfc6aa] shadow-xl"
                    aria-hidden="true"
                />

                <div
                    className="absolute inset-[2%] translate-x-[2%] translate-y-[2%] rounded-sm bg-[#e5dec8]"
                    aria-hidden="true"
                />

                <Image
                    src={src}
                    alt={alt}
                    width={imgWidth}
                    height={imgHeight}
                    className="relative h-auto w-full rounded-sm shadow-2xl"
                />
            </div>
        </div>
    );
}
