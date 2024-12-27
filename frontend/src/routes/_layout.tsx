import { Flex, Spinner } from "@chakra-ui/react";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

import Sidebar from "../components/Common/Sidebar";
import UserMenu from "../components/Common/UserMenu";
import HeaderTabs from "../components/Common/HeaderTabs";
import useAuth from "../hooks/useAuth";

export const Route = createFileRoute("/_layout")({
  component: Layout,
});

function Layout() {
  const { isPending } = useAuth();

  return (
    <Flex maxW="large" h="auto" position="relative">
      {isPending ? (
        <>
          <Flex justify="center" align="center" height="100vh" width="full">
            <Spinner size="xl" color="ui.main" />
          </Flex>
        </>
      ) : (
        <>
          <Sidebar />
          <Outlet />
          <UserMenu />
        </>
      )}
    </Flex>
  );
}
