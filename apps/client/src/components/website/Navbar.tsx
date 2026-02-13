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
import Link from 'next/link';
import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { useRouter } from 'next/navigation';
import { Menu, User, LogOut, LayoutDashboard } from 'lucide-react';
import { SignedIn, SignedOut, UserAvatar } from '@daveyplate/better-auth-ui';

import NavLink from './NavLink';
import { LogoIcon } from '../Icons';
import { ModeToggle } from '../mode-toggle';
import { Button, buttonVariants } from '../ui/button';
import { authClient, isAdmin, isEditor } from '@/lib/auth-client';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import LanguageSwitcher from '../language-switcher';

interface NavbarProps {
  appTitle?: string;
  appLogo?: string;
  navItems?: components['schemas']['HeaderNavItemDto'][];
}

export const Navbar = ({
  appTitle = 'NienSuViet',
  appLogo,
  navItems,
}: NavbarProps) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const router = useRouter();
  const { data: session } = authClient.useSession();
  const t = useTranslations('HomePage');

  const handleSignOut = async () => {
    await authClient.signOut();
    router.push('/');
  };

  return (
    <header className="sticky border-b top-0 z-40 w-full bg-white dark:border-b-slate-700 dark:bg-background">
      <NavigationMenu className="mx-auto">
        <NavigationMenuList className="container h-14 px-4 w-screen flex justify-start">
          <NavigationMenuItem className="font-bold flex">
            <Link
              rel="noreferrer noopener"
              href="/"
              className="ml-2 font-bold text-xl flex items-center py-1"
            >
              {appLogo ? (
                <Image
                  src={appLogo}
                  alt={appTitle}
                  width={100}
                  height={50}
                  className="mr-2 object-contain"
                />
              ) : (
                <LogoIcon />
              )}
            </Link>
          </NavigationMenuItem>

          {/* mobile */}
          <span className="flex md:hidden ml-auto">
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
                  </SheetTitle>
                </SheetHeader>
                <nav className="flex flex-col justify-center items-center gap-2">
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
          <nav className="hidden md:flex gap-2 ml-20">
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

          <div className="hidden md:flex gap-2 ml-auto">
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
                      {t('Auth.account')}
                    </Link>
                  </DropdownMenuItem>

                  {isEditor(session?.user.role) && (
                    <DropdownMenuItem asChild>
                      <Link href="/cmsdesk" className="cursor-pointer">
                        <LayoutDashboard className="mr-2 h-4 w-4" />
                        {t('Auth.cms-desk')}
                      </Link>
                    </DropdownMenuItem>
                  )}

                  {isAdmin(session?.user.role) && (
                    <DropdownMenuItem asChild>
                      <Link href="/admin" className="cursor-pointer">
                        <LayoutDashboard className="mr-2 h-4 w-4" />
                        {t('Auth.admin-dashboard')}
                      </Link>
                    </DropdownMenuItem>
                  )}
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    onClick={handleSignOut}
                    className="cursor-pointer text-destructive focus:text-destructive"
                  >
                    <LogOut className="mr-2 h-4 w-4" />
                    {t('Auth.sign-out')}
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </SignedIn>

            <SignedOut>
              <Button asChild>
                <Link href="/auth/sign-in">{t('Auth.sign-in')}</Link>
              </Button>
            </SignedOut>

            <LanguageSwitcher />

            <ModeToggle />
          </div>
        </NavigationMenuList>
      </NavigationMenu>
    </header>
  );
};
