import { readUserMeV1UsersMeGet } from "../clients/iam/sdk.gen";
import { useQuery } from "@tanstack/react-query";

const useAuth = () => {
  const {
    isPending,
    isError,
    data: user,
    error,
  } = useQuery({
    queryKey: ["currentUser"],
    queryFn: readUserMeV1UsersMeGet,
  });

  const logout = () => {
    console.log("logout");
  };

  return {
    user,
    isPending,
    isError,
    error,
    logout,
  };
};

export default useAuth;
