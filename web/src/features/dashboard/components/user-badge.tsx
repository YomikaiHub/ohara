"use client";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useMe } from "@/hooks/use-auth";
import { ChevronsUpDown, Cog, LogOut } from "lucide-react";

function UserBadge() {
    const { data: user, isLoading, isError } = useMe();

    if (isLoading) return "Loading...";

    if (isError || !user) return "Error...";

    return (
        <DropdownMenu>
            <DropdownMenuTrigger>
                <div className="flex items-center gap-3 rounded-full bg-background px-2 py-1">
                    <Avatar className="h-8 w-8">
                        <AvatarImage src="/profile.jpg" alt={user.first_name} />
                        <AvatarFallback className="rounded-lg">
                            {user.first_name}
                        </AvatarFallback>
                    </Avatar>
                    <div className="flex flex-1 items-center gap-2 text-left text-sm leading-tight">
                        Hello,{" "}
                        <span className="truncate font-bold text-teal-300">
                            {user.first_name}
                        </span>
                    </div>
                    <ChevronsUpDown className="ml-auto size-4" />
                </div>
            </DropdownMenuTrigger>
            <DropdownMenuContent
                className="w-fit min-w-56 rounded-lg"
                side={"bottom"}
                align="end"
                sideOffset={8}
            >
                <DropdownMenuGroup>
                    <DropdownMenuLabel className="p-0 font-normal">
                        <div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
                            <Avatar className="h-8 w-8 rounded-lg">
                                <AvatarImage
                                    src="/profile.jpg"
                                    alt={user.first_name}
                                />
                                <AvatarFallback className="rounded-lg">
                                    {user.first_name}
                                </AvatarFallback>
                            </Avatar>
                            <div className="grid flex-1 text-left text-sm leading-tight">
                                <h4 className="flex items-center gap-1 truncate font-medium">
                                    {user.first_name} {user.last_name}
                                </h4>
                                <span className="truncate text-xs">
                                    {user.email}
                                </span>
                            </div>
                        </div>
                    </DropdownMenuLabel>
                </DropdownMenuGroup>
                <DropdownMenuSeparator />
                <DropdownMenuGroup>
                    <DropdownMenuItem>
                        <Cog /> Setting
                    </DropdownMenuItem>
                    <DropdownMenuItem>
                        <LogOut />
                        Log out
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

export default UserBadge;
