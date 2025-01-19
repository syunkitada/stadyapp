import { createFileRoute } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";

import { readItemsV1ItemsGet } from "../../../clients/compute/sdk.gen";

export const Route = createFileRoute("/_layout/compute/")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Compute</div>;

  // const { isPending, isError, data, error } = useQuery({
  //   queryKey: ["servers"],
  //   queryFn: readItemsV1ItemsGet,
  // });

  // console.log("DEBUG", isPending, isError, data, error);

  // if (isPending) {
  //   return <div>Pending</div>;
  // }

  // return (
  //   <div>
  //     <h1>Compute</h1>
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
