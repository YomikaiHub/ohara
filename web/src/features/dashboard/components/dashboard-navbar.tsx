"use client";

import { LogoAvatar } from "@/components/ui/avatar";
import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import { BookSearch, FileDown, House, Notebook } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import UserBadge from "./user-badge";

const navbarLinks = [
    {
        id: "home",
        title: "Home",
        icon: House,
        href: "/dashboard",
    },
    {
        id: "library",
        title: "Library",
        icon: Notebook,
        href: "/dashboard/library",
    },
    {
        id: "search",
        title: "Search Book",
        icon: BookSearch,
        href: "/dashboard/search",
    },
    {
        id: "import",
        title: "Import Book",
        icon: FileDown,
        href: "/dashboard/import",
    },
];

function DashboardNavbar() {
    const pathname = usePathname();

    return (
        <header className="h-18 w-full px-2 py-3">
            <nav className="flex h-full items-center justify-between rounded-full bg-secondary px-2">
                <div className="flex items-center gap-1 rounded-full bg-background px-2 py-1">
                    <LogoAvatar
                        imageUrl="/logo-tr.svg"
                        imageClassName="w-7 h-7"
                    />
                    <h1 className="pr-1 text-lg">Ohara</h1>
                </div>
                <div className="flex items-center gap-3">
                    {navbarLinks.map((nl) => {
                        const Icon = nl.icon;

                        return (
                            <NavigationMenu key={nl.id}>
                                <NavigationMenuList>
                                    <NavigationMenuItem>
                                        <NavigationMenuLink
                                            render={<Link href={nl.href} />}
                                            className={cn(
                                                navigationMenuTriggerStyle(),
                                                pathname === nl.href
                                                    ? "bg-primary text-primary-foreground hover:bg-primary"
                                                    : "hover:bg-background"
                                            )}
                                        >
                                            <Icon className="size-4" />
                                            {nl.title}
                                        </NavigationMenuLink>
                                    </NavigationMenuItem>
                                </NavigationMenuList>
                            </NavigationMenu>
                        );
                    })}
                </div>
                <div className="flex items-center justify-center">
                    <UserBadge />
                </div>
            </nav>
        </header>
    );
}

export default DashboardNavbar;
