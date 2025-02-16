import { Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout")({
  component: Layout,
  loader: () => ({
    crumb: "Home",
  }),
});

function Layout() {
  return <Outlet />;
}
