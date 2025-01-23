import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import { Tabs } from "@chakra-ui/react";

export const Route = createFileRoute("/_layout/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <>
      <div>
        <p>hoge</p>
        <Button />
      </div>
    </>
  );
}
