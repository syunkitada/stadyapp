import { Button } from "@/components/ui/button";

import { AppSidebar } from "@/components/app-sidebar";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";

import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

import useAuth from "../hooks/useAuth";

export const Route = createFileRoute("/_layout")({
  component: Layout,
});

function Layout() {
  const { isPending } = useAuth();

  return (
    <div>
      {isPending ? (
        <>
          <Button>Spinner</Button>
        </>
      ) : (
        <>
          <Outlet />
        </>
      )}
    </div>
  );
}
