import { getKeystoneDomains } from "@/clients/iam/sdk.gen";
import { useQuery } from "@tanstack/react-query";

import { createFileRoute } from "@tanstack/react-router";

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

  return <div>hoge</div>;

  // return (
  //   <div>
  //     <h1>IAM</h1>
  //     <ul>
  //       {data.data.data.map((item) => (
  //         <li key={item.id}>
  //           {item.title}: {item.description}
  //         </li>
  //       ))}
  //     </ul>
  //   </div>
  // );
}
