import { CenterLoader } from "@/components/common/loader";
import { DataTable } from "@/components/project_services/compute/server/data-table";
import { useServers } from "@/hooks/useCompute";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute(
  "/_layout/projects/$projectId/_layout/compute/server/",
)({
  component: RouteComponent,
  loader: () => ({
    crumb: "Server",
  }),
});

function RouteComponent() {
  const { isPending, isError, data, error } = useServers({
    refreshInterval: 10000,
  });

  console.log("DEBUG servers", isPending, isError, data, error);

  if (isPending) {
    return <CenterLoader />;
  }

  if (isError) {
    return <div>Error: {error}</div>;
  }

  if (data.error) {
    return <div>Error</div>;
  }

  return <DataTable data={data.data.servers} />;
}
