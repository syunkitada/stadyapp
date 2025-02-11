"use client";

import { ActionServerDialog } from "./action-server-dialog";
import { deleteNovaServerById } from "@/clients/compute/sdk.gen";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

const formSchema = z.object({});

export function DeleteServerDialog({
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
      toast.success("Requested to delete server");
      setOpen(false);
    },
    onError: (err: any) => {
      console.log("error", err);
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("delete onSubmit", values);
    mutation.mutate(targets[0].id);
  }

  return (
    <ActionServerDialog
      title="Delete"
      description="Delete Server"
      submitName="Delete"
      open={open}
      setOpen={setOpen}
      targets={targets}
      onSubmit={onSubmit}
      form={form}
    />
  );
}
