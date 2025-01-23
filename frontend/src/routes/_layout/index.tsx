import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/")({
  beforeLoad: ({ context, location }) => {
    throw redirect({
      to: "/services/",
    });
  },
});
