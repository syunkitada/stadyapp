import { Flex, Spinner } from "@chakra-ui/react";
import { createRootRoute, Link, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import useAuth from "../hooks/useAuth";

export const Route = createFileRoute("/_layout")({
  component: Layout,
});

function Layout() {
  const { user, isPending } = useAuth();

  console.log("isPending", isPending, user);
  return (
    <>
      {isPending ? <p>Loading...</p> : <p>Hello: {user.data.name}</p>}
      <div className="p-2 flex gap-2">Hello World</div>
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  );
}
