import { getWebUser } from "../clients/iam/sdk.gen";
import { useQuery } from "@tanstack/react-query";

import { client as clientIAM } from "./clients/iam/sdk.gen";
import { client as clientServer } from "./clients/compute/sdk.gen";

const useAuth = () => {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["getWebUser"],
    queryFn: getWebUser,
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
