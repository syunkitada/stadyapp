"use client";

import {
  getNovaServersDetail,
  getNovaServerById,
} from "@/clients/compute/sdk.gen";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";

export const ACTION_STATUS = {
  NONE: "None",
  PROCESSING: "Processing",
  PROCESSED: "Processed",
  ERROR: "Error",
};

export function useActionTargets() {
  const [targets, setTargets] = useState([]);

  const setInitActionTargets = (selectedRows) => {
    const actionTargets = selectedRows.map((row) => {
      return {
        id: row.original.id,
        name: row.original.name,
        status: ACTION_STATUS.NONE,
        statusMessage: "",
      };
    });
    setTargets(actionTargets);
  };

  const setActionTargets = (targets, status) => {
    for (const [index, target] of targets.entries()) {
      targets[index].status = status;
    }
    setTargets(targets);
  };

  const setActionTargetStatus = (index, status, statusMessage) => {
    targets[index].status = status;
    targets[index].statusMessage = statusMessage;
    setTargets(targets);
  };

  return [
    targets,
    setInitActionTargets,
    setActionTargets,
    setActionTargetStatus,
  ];
}

export function useServers({ refreshInterval }: { refreshInterval: number }) {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["getNovaServersDetail"],
    queryFn: getNovaServersDetail,
    refetchInterval: refreshInterval,
  });

  return {
    isPending,
    isError,
    data,
    error,
  };
}

export function useReloadServers() {
  const queryClient = useQueryClient();

  const reloadServers = () => {
    queryClient.invalidateQueries({
      queryKey: ["getNovaServersDetail"],
    });
  };

  return {
    reloadServers,
  };
}

export function useServer({
  id,
  refreshInterval,
}: {
  id: string;
  refreshInterval: number;
}) {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["getNovaServerById", { id }],
    queryFn: () => getNovaServerById({ path: { id } }),
    refetchInterval: refreshInterval,
  });

  return {
    isPending,
    isError,
    data,
    error,
  };
}
