import { ReactNode } from "react";

interface StatBadgeProps {
    icon: ReactNode;
    label: string;
    value: string;
    iconColor?: string;
}

function StatsBadge({
    icon,
    label,
    value,
    iconColor = "text-teal-400",
}: StatBadgeProps) {
    return (
        <div className="flex items-center gap-3 rounded-xl border border-white/10 bg-white/[0.04] px-4 py-2.5 shadow-lg shadow-black/30 backdrop-blur-sm">
            <span className={`shrink-0 ${iconColor}`}>{icon}</span>
            <div className="leading-tight">
                <p className="text-[11px] text-slate-300/80">{label}</p>
                <p className="text-lg font-semibold text-white">{value}</p>
            </div>
        </div>
    );
}

export default StatsBadge;
