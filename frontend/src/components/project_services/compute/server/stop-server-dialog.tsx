"use client";

import { ActionServerDialog } from "./action-server-dialog";
import { actionNovaServer } from "@/clients/compute/sdk.gen";
import { ACTION_STATUS, useReloadServers } from "@/hooks/useCompute";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

const formSchema = z.object({});

export function StopServerDialog({
  open,
  setOpen,
  actionTargets,
  setActionTargets,
  setActionTargetStatus,
}: {
  open: any;
  setOpen: any;
  actionTargets: any[];
  setActionTargets: any;
  setActionTargetStatus: any;
}) {
  const {reloadServers} = useReloadServers();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {},
  });

  const mutation = useMutation({
    mutationFn: ({ index, id }: { index: number; id: string }) =>
      actionNovaServer({ path: { id }, body: { "os-stop": null } }),
    onSuccess: (data, variables, context) => {
      if (data.error) {
        setActionTargetStatus(
          variables.index,
          ACTION_STATUS.ERROR,
          data.message,
        );
      } else {
        setActionTargetStatus(variables.index, ACTION_STATUS.PROCESSED, "");
      }
    },
    onError: (error, variables, context) => {
      console.error(error, variables, context);
      setActionTargetStatus(variables.index, ACTION_STATUS.ERROR, "");
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    setActionTargets(actionTargets, ACTION_STATUS.PROCESSING);

    for (const [index, target] of actionTargets.entries()) {
      mutation.mutate(
        { index: index, id: target.id },
        {
          onSuccess: (data, variables, context) => {
            reloadServers();
            toast.success("Requested to stop server");
          },
          onError: (error, variables, context) => {
            console.error(error, variables, context);
            toast.error("Failed to stop server");
          },
        },
      );
    }
  }

  return (
    <ActionServerDialog
      title="Stop"
      description="Stop Server"
      submitName="Stop"
      open={open}
      setOpen={setOpen}
      actionTargets={actionTargets}
      onSubmit={onSubmit}
      form={form}
    />
  );
}
