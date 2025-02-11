import { Tabs, TabsContent, TabsList, Tab } from "@/components/common/tabs";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Link } from "@tanstack/react-router";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute(
  "/_layout/projects/$projectId/_layout/compute/server/$serverId/_layout",
)({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <>
      <Tab>
        <Link to={"/projects/$projectId/compute/server/$serverId/detail"}>
          Detail
        </Link>
      </Tab>
      <Tab>
        <Link to={"/projects/$projectId/compute/server/$serverId/eventlog"}>
          Event Log
        </Link>
      </Tab>
      <Outlet />
    </>
  );
}
