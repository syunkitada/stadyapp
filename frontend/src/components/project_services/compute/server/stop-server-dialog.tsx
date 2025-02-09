"use client";

import { ActionServerDialog } from "./action-server-dialog";
import { deleteNovaServerById } from "@/clients/compute/sdk.gen";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({});

export function StopServerDialog({
  open,
  setOpen,
  targets,
}: {
  open: any;
  setOpen: any;
  targets: any[];
}) {
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {},
  });

  const mutation = useMutation({
    mutationFn: (id: string) => deleteNovaServerById({ path: { id } }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["getNovaServersDetail"],
      });
      setOpen(false);
    },
    onError: (err: any) => {
      console.log("error", err);
    },
  });

  console.log("StopServerDialog", targets);

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("delete onSubmit", values);
    mutation.mutate(targets[0].id);
  }

  return (
    <ActionServerDialog
      title="Stop"
      description="Stop Server"
      submitName="Stop"
      open={open}
      setOpen={setOpen}
      targets={targets}
      onSubmit={onSubmit}
      form={form}
    />
  );
}
