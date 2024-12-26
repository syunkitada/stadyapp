import { createFileRoute } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";

import { readItemsV1ItemsGet } from "../../../clients/iam/sdk.gen";

export const Route = createFileRoute("/_layout/iam/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["items"],
    queryFn: readItemsV1ItemsGet,
  });

  console.log("DEBUG", isPending, isError, data, error);

  if (isPending) {
    return <div>Pending</div>;
  }

  return (
    <div>
      <h1>IAM</h1>
      <ul>
        {data.data.data.map((item) => (
          <li key={item.id}>
            {item.title}: {item.description}
          </li>
        ))}
      </ul>
    </div>
  );
}
