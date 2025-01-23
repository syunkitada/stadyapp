import { getWebUser } from "../clients/iam/sdk.gen";
import { useQuery } from "@tanstack/react-query";

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
