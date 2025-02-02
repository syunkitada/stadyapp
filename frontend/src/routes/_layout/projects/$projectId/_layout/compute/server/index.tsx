import { CenterLoader } from "@/components/common/loader";

import { getNovaServersDetail } from "@/clients/compute/sdk.gen";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { DataTable } from "@/components/project_services/compute/server/data-table";

export const Route = createFileRoute(
  "/_layout/projects/$projectId/_layout/compute/server/",
)({
  component: RouteComponent,
});

function RouteComponent() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["servers"],
    queryFn: getNovaServersDetail,
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
