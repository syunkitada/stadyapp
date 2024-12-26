import { ChakraProvider } from "@chakra-ui/react";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// Import the generated route tree
import { routeTree } from "./routeTree.gen";

import "./index.css";
import App from "./App.tsx";

import { client as clientIAM } from "./clients/iam/sdk.gen";
import { client as clientServer } from "./clients/compute/sdk.gen";

clientIAM.setConfig({
  baseURL: import.meta.env.VITE_IAM_API_URL,
  headers: {
    "x-user": import.meta.env.VITE_TEST_USER,
    "x-email": import.meta.env.VITE_TEST_EMAIL,
  },
});

clientServer.setConfig({
  baseURL: import.meta.env.VITE_COMPUTE_API_URL,
  headers: {
    "x-user": import.meta.env.VITE_TEST_USER,
    "x-email": import.meta.env.VITE_TEST_EMAIL,
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
    <ChakraProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ChakraProvider>
  </StrictMode>,
);
