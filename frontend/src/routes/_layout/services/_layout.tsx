import { Breadcrumbs } from "@/components/common/breadcrumbs";
import { CenterLoader } from "@/components/common/loader";
import { ServicesSidebar } from "@/components/services-sidebar";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import useAuth from "@/hooks/useAuth";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/services/_layout")({
  component: RouteComponent,
});

function RouteComponent() {
  const { isPending } = useAuth();

  if (isPending) {
    return (
      <>
        <CenterLoader />;
      </>
    );
  }

  return (
    <SidebarProvider>
      <ServicesSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
          <div className="flex items-center gap-2 px-4">
            <SidebarTrigger className="-ml-1" />
            <Separator orientation="vertical" className="mr-2 h-4" />
            <Breadcrumbs />
          </div>
        </header>
        <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
          <Outlet />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
