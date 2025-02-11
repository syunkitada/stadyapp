"use client";

import { ActionServerDialog } from "./action-server-dialog";
import { actionNovaServer } from "@/clients/compute/sdk.gen";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import * as React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

const formSchema = z.object({});

export function StartServerDialog({
  open,
  setOpen,
  targets,
  setTargets,
}: {
  open: any;
  setOpen: any;
  targets: any[];
  setTargets: any;
}) {
  console.log("DEBUG StartServerDialog", targets);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {},
  });

  const mutation = useMutation({
    mutationFn: ({ index, id }: { index: number; id: string }) =>
      actionNovaServer({ path: { id }, body: { "os-start": null } }),
    onSuccess: (data, variables, context) => {
      targets[variables.index].actionStatus = "Processed";
      setTargets(targets);
    },
    onError: (err: any) => {
      console.log("start onerror", err);
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    for (const [index, target] of targets.entries()) {
      targets[index].actionStatus = "Processing";
    }
    setTargets(targets);

    for (const [index, target] of targets.entries()) {
      console.log("DEBUG start index", index);
      mutation.mutate(
        { index: index, id: target.id },
        {
          onSuccess: (data, variables, context) => {
            queryClient.invalidateQueries({
              queryKey: ["getNovaServersDetail"],
            });
            toast.success("Requested to start server");
          },
          onError: (error, variables, context) => {
            console.log("DEBUG start onError", index, target.id);
          },
        },
      );
    }
    console.log("delete onSubmit2", values);
  }

  return (
    <ActionServerDialog
      title="Start"
      description="Start Server"
      submitName="Start"
      open={open}
      setOpen={setOpen}
      targets={targets}
      onSubmit={onSubmit}
      form={form}
    />
  );
}
