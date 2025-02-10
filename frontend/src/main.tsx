import { client as clientServer } from "./clients/compute/client.gen";
import { client as clientIAM } from "./clients/iam/client.gen";
import "./index.css";
// Import the generated route tree
import { routeTree } from "./routeTree.gen";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

clientIAM.setConfig({
  baseURL: import.meta.env.VITE_IAM_API_URL,
  headers: {
    "x-user-id": import.meta.env.VITE_TEST_USER,
  },
});

clientServer.setConfig({
  baseURL: import.meta.env.VITE_COMPUTE_API_URL,
  headers: {
    "x-user-id": import.meta.env.VITE_TEST_USER,
  },
});

const queryClient = new QueryClient();

// Create a new router instance
const router = createRouter({ routeTree });

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </StrictMode>,
);
