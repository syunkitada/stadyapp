import { getKeystoneDomains } from "@/clients/iam/sdk.gen";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { DataTable } from "@/components/services/iam/domains/data-table";

export const Route = createFileRoute("/_layout/services/_layout/iam/domains/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["domains"],
    queryFn: getKeystoneDomains,
  });

  console.log("DEBUG dmains", isPending, isError, data, error);

  if (isPending) {
    return <div>Pending</div>;
  }

  console.log("DEBUG domains", data.data.domains);

  return <DataTable data={data.data.domains} />;
}
