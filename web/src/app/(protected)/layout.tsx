import DashboardNavbar from "@/features/dashboard/components/dashboard-navbar";

export default function ProtectedLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <>
            <DashboardNavbar />
            {children}
        </>
    );
}
