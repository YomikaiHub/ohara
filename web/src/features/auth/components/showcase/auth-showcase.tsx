import { Bookmark, BookOpen, Flame, Heart, Plus, Star } from "lucide-react";
import StatsBadge from "./stats-badge";
import Feature from "./feature";
import { BookShowcase } from "./book-showcase";

function AuthShowcase() {
    return (
        <div className="relative flex h-full min-h-screen w-full flex-col overflow-hidden bg-background px-10 py-10 md:px-16 md:py-14">
            <div
                className="pointer-events-none absolute inset-0 opacity-60"
                style={{
                    backgroundImage:
                        "radial-gradient(1px 1px at 20% 30%, white, transparent), radial-gradient(1px 1px at 70% 15%, white, transparent), radial-gradient(1.5px 1.5px at 40% 70%, white, transparent), radial-gradient(1px 1px at 85% 60%, white, transparent), radial-gradient(1px 1px at 10% 80%, white, transparent), radial-gradient(1.5px 1.5px at 55% 40%, white, transparent), radial-gradient(1px 1px at 90% 90%, white, transparent), radial-gradient(1px 1px at 30% 10%, white, transparent)",
                    backgroundSize: "100% 100%",
                }}
            />
            <div className="pointer-events-none absolute top-[52%] left-1/2 h-105 w-155 -translate-x-1/2 -translate-y-1/2 rounded-full bg-teal-500/20 blur-[100px]" />

            <div className="relative z-10 text-center">
                <h1 className="text-4xl leading-tight font-bold text-white md:text-[42px]">
                    Your next chapter
                </h1>
                <h1 className="text-4xl leading-tight font-bold text-teal-400 md:text-[42px]">
                    starts here.
                </h1>
                <p className="mx-auto mt-4 max-w-sm text-sm text-slate-300/80">
                    Discover stories, track your journey and build a library
                    you&rsquo;ll love.
                </p>
            </div>

            <div className="relative z-10 mt-8 flex items-start justify-between">
                <StatsBadge
                    icon={<BookOpen size={18} />}
                    label="Currently Reading"
                    value="12"
                />
                <StatsBadge
                    icon={
                        <Star
                            size={18}
                            className="fill-amber-400 text-amber-400"
                        />
                    }
                    label="Books in Library"
                    value="127"
                    iconColor="text-amber-400"
                />
            </div>

            <div className="relative z-10 mx-auto flex flex-1 items-center justify-center">
                <BookShowcase
                    src="/book1.jpg"
                    alt="1984"
                    imgWidth={190}
                    imgHeight={300}
                    className="relative z-10 -mr-8 translate-y-5 rotate-[-9deg]"
                />
                <BookShowcase
                    src="/book2.jpg"
                    alt="Dune"
                    imgWidth={240}
                    imgHeight={360}
                    className="relative z-20 -mr-8 -translate-y-2 -rotate-2"
                />
                <BookShowcase
                    src="/book3.jpg"
                    alt="Project Hail Mary"
                    imgWidth={190}
                    imgHeight={360}
                    className="relative z-10 -mr-8 translate-y-4 rotate-5"
                />
                <BookShowcase
                    src="/book4.jpg"
                    alt="Book"
                    imgWidth={180}
                    imgHeight={270}
                    className="relative z-0 translate-y-8 rotate-10"
                />
            </div>
            <div className="relative z-10 mt-6 flex items-end justify-between">
                <StatsBadge
                    icon={
                        <Flame
                            size={18}
                            className="fill-orange-400 text-orange-400"
                        />
                    }
                    label="Reading Streak"
                    value="12 Days"
                    iconColor="text-orange-400"
                />
                <StatsBadge
                    icon={<Plus size={18} />}
                    label="Added to Library"
                    value="8 Today"
                />
            </div>

            <div className="relative z-10 mt-10 grid grid-cols-3 gap-6 text-center md:text-left">
                <Feature
                    icon={<BookOpen size={20} />}
                    title="Discover"
                    desc="Find your next favorite book"
                />
                <Feature
                    icon={<Bookmark size={20} />}
                    title="Track"
                    desc="Keep track of your reading journey"
                />
                <Feature
                    icon={<Heart size={20} />}
                    title="Explore"
                    desc="Explore books and grow your library"
                />
            </div>
        </div>
    );
}

export default AuthShowcase;
