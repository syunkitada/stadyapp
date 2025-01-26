import { useParams } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";

import { getWebUser } from "@/clients/iam/sdk.gen";

const useAuth = () => {
  const { projectId } = useParams({ strict: false });

  const { isPending, isError, data, error } = useQuery({
    queryKey: ["getWebUser", { projectId }],
    queryFn: () => {
      return getWebUser({ query: { project_id: projectId } });
    },
  });

  const logout = () => {
    console.log("logout");
  };

  return {
    user: data,
    isPending,
    isError,
    error,
    logout,
  };
};

export default useAuth;
