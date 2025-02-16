import { CenterLoader } from "@/components/common/loader";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import { useServer } from "@/hooks/useCompute";

export function ServerDetail({ id }: { id: string }) {
  const { isPending, isError, data, error } = useServer({
    id: id,
    refreshInterval: 10000,
  });

  if (isPending) {
    return <CenterLoader />;
  }

  if (isError) {
    console.error(error);
    return <div>Error</div>;
  }

  if (!data) {
    console.error(error);
    return <div>Error</div>;
  }

  if (data.status != 200) {
    console.error(data);
    return <div>Unexpected Status</div>;
  }

  if (!data.data.server) {
    return <div>No Server found</div>;
  }

  const server = data.data.server;

  return (
    <Table>
      <TableBody>
        <TableRow>
          <TableCell className="font-medium">ID</TableCell>
          <TableCell className="font-medium">{server.id}</TableCell>
        </TableRow>
        <TableRow>
          <TableCell className="font-medium">Name</TableCell>
          <TableCell className="font-medium">{server.name}</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  );
}
