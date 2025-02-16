import {
  Tabs,
  TabsTrigger,
  TabsContent,
  TabsList,
} from "@/components/common/tabs";
import { ServerDetail } from "@/components/project_services/compute/server/server-detail";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { z } from "zod";

const searchSchema = z.object({
  tab: z.string().default("detail"),
});

export const Route = createFileRoute(
  "/_layout/projects/$projectId/_layout/compute/server/$serverId/",
)({
  component: RouteComponent,
  validateSearch: searchSchema,
  loader: () => ({
    crumb: "Detail",
  }),
});

function RouteComponent() {
  const { serverId } = Route.useParams();
  const { tab } = Route.useSearch();

  const navigate = useNavigate({ from: Route.fullPath });

  const onValueChange = (value: string) => {
    navigate({ search: { tab: value } });
  };

  return (
    <Tabs defaultValue={tab} onValueChange={onValueChange}>
      <TabsList className="grid w-full grid-cols-3">
        <TabsTrigger value="detail">Detail</TabsTrigger>
        <TabsTrigger value="eventlog">EventLog</TabsTrigger>
        <TabsTrigger value="consolelog">ConsoleLog</TabsTrigger>
      </TabsList>
      <TabsContent value="detail">
        <ServerDetail id={serverId} />
      </TabsContent>
      <TabsContent value="eventlog">eventlog</TabsContent>
      <TabsContent value="consolelog">consolelog</TabsContent>
    </Tabs>
  );
}
