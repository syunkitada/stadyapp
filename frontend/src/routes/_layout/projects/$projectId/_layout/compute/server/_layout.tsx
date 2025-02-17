import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute(
  "/_layout/projects/$projectId/_layout/compute/server/_layout",
)({
  component: RouteComponent,
  loader: () => ({
    crumb: "Server",
  }),
});

function RouteComponent() {
  return <Outlet />;
}
