import { ReactNode } from "react";

interface FeatureProps {
    icon: ReactNode;
    title: string;
    desc: string;
}

function Feature({ icon, title, desc }: FeatureProps) {
    return (
        <div className="flex flex-col items-center gap-2 md:flex-row md:items-start">
            <span className="text-teal-400">{icon}</span>
            <div>
                <p className="text-sm font-semibold text-white">{title}</p>
                <p className="text-xs text-slate-400">{desc}</p>
            </div>
        </div>
    );
}

export default Feature;
