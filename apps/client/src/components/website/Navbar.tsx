'use client';

import { useState } from 'react';
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from '@/components/ui/navigation-menu';
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import { Button, buttonVariants } from '../ui/button';
import { Menu, User, LogOut, LayoutDashboard } from 'lucide-react';
import { ModeToggle } from '../mode-toggle';
import { LogoIcon } from '../Icons';
import Link from 'next/link';
import { SignedIn, SignedOut, UserAvatar } from '@daveyplate/better-auth-ui';
import { useRouter } from 'next/navigation';
import { authClient, isAdmin, isEditor } from '@/lib/auth-client';
import Image from 'next/image';
import { components } from '@nsv-interfaces/cms-service';
import NavLink from './NavLink';

interface RouteProps {
  href: string;
  label: string;
}

interface NavbarProps {
  appTitle?: string;
  appLogo?: string;
  navItems?: components['schemas']['HeaderNavItemData'][];
}

export const Navbar = ({
  appTitle = 'NienSuViet',
  appLogo,
  navItems,
}: NavbarProps) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const router = useRouter();
  const { data: session } = authClient.useSession();

  const handleSignOut = async () => {
    await authClient.signOut();
    router.push('/');
  };

  return (
    <header className="sticky border-b top-0 z-40 w-full bg-white dark:border-b-slate-700 dark:bg-background">
      <NavigationMenu className="mx-auto">
        <NavigationMenuList className="container h-14 px-4 w-screen flex justify-between">
          <NavigationMenuItem className="font-bold flex">
            <Link
              rel="noreferrer noopener"
              href="/"
              className="ml-2 font-bold text-xl flex items-center"
            >
              {appLogo ? (
                <Image
                  src={appLogo}
                  alt={appTitle}
                  width={32}
                  height={32}
                  className="mr-2 h-8 w-8 object-contain"
                />
              ) : (
                <LogoIcon />
              )}
              {appTitle}
            </Link>
          </NavigationMenuItem>

          {/* mobile */}
          <span className="flex md:hidden">
            <ModeToggle />

            <Sheet open={isOpen} onOpenChange={setIsOpen}>
              <SheetTrigger className="px-2">
                <Menu
                  className="flex md:hidden h-5 w-5"
                  onClick={() => setIsOpen(true)}
                >
                  <span className="sr-only">Menu Icon</span>
                </Menu>
              </SheetTrigger>

              <SheetContent side={'left'}>
                <SheetHeader>
                  <SheetTitle className="font-bold text-xl flex items-center">
                    {appLogo ? (
                      <Image
                        src={appLogo}
                        alt={appTitle}
                        width={24}
                        height={24}
                        className="mr-2 h-6 w-6 object-contain"
                      />
                    ) : (
                      <LogoIcon />
                    )}
                    {appTitle}
                  </SheetTitle>
                </SheetHeader>
                <nav className="flex flex-col justify-center items-center gap-2 mt-4">
                  {navItems
                    ?.sort((a, b) => a.order - b.order)
                    .map((item, i) => (
                      <NavLink
                        key={i}
                        onClick={() => setIsOpen(false)}
                        className={buttonVariants({ variant: 'ghost' })}
                        navItem={item}
                      />
                    ))}
                </nav>
              </SheetContent>
            </Sheet>
          </span>

          {/* desktop */}
          <nav className="hidden md:flex gap-2">
            {navItems
              ?.sort((a, b) => a.order - b.order)
              .map((item, i) => (
                <NavLink
                  key={i}
                  className={`text-[17px] ${buttonVariants({
                    variant: 'ghost',
                  })}`}
                  navItem={item}
                />
              ))}
          </nav>

          <div className="hidden md:flex gap-2">
            <SignedIn>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button
                    variant="ghost"
                    className="relative h-9 w-9 rounded-full p-0"
                  >
                    <UserAvatar />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end" className="w-56">
                  <DropdownMenuLabel>
                    <div className="flex flex-col space-y-1">
                      <p className="text-sm font-medium leading-none">
                        {session?.user?.name || 'User'}
                      </p>
                      <p className="text-xs leading-none text-muted-foreground">
                        {session?.user?.email || ''}
                      </p>
                    </div>
                  </DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem asChild>
                    <Link href="/account/settings" className="cursor-pointer">
                      <User className="mr-2 h-4 w-4" />
                      Account
                    </Link>
                  </DropdownMenuItem>

                  {isEditor(session?.user.role) && (
                    <DropdownMenuItem asChild>
                      <Link href="/cmsdesk" className="cursor-pointer">
                        <LayoutDashboard className="mr-2 h-4 w-4" />
                        CMS Desk
                      </Link>
                    </DropdownMenuItem>
                  )}

                  {isAdmin(session?.user.role) && (
                    <DropdownMenuItem asChild>
                      <Link href="/admin" className="cursor-pointer">
                        <LayoutDashboard className="mr-2 h-4 w-4" />
                        Admin Dashboard
                      </Link>
                    </DropdownMenuItem>
                  )}
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    onClick={handleSignOut}
                    className="cursor-pointer text-destructive focus:text-destructive"
                  >
                    <LogOut className="mr-2 h-4 w-4" />
                    Sign Out
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </SignedIn>

            <SignedOut>
              <Button asChild>
                <Link href="/auth/sign-in">Login</Link>
              </Button>
            </SignedOut>
            <ModeToggle />
          </div>
        </NavigationMenuList>
      </NavigationMenu>
    </header>
  );
};
