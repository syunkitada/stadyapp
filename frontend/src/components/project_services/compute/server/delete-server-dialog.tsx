"use client";

import { deleteNovaServerById } from "@/clients/compute/sdk.gen";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Form } from "@/components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import * as React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  name: z.string().min(2, {
    message: "Server name must be at least 2 characters.",
  }),
  flavor: z.string({
    required_error: "Please select a flavor.",
  }),
  image: z.string({
    required_error: "Please select a image.",
  }),
  network: z.string({
    required_error: "Please select a network.",
  }),
});

export function DeleteServerDialog({
  open,
  setOpen,
}: {
  open: any;
  setOpen: any;
}) {
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      image: "",
      flavor: "",
      network: "",
    },
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

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("dialog onSubmit", values);
    console.log(values.name);
    console.log(values.image);
    console.log(values.flavor);
    console.log(values.network);
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[800px]">
        <DialogHeader>
          <DialogTitle>Delete Server</DialogTitle>
          <DialogDescription>
            Make changes to your profile here. Click save when you're done.
          </DialogDescription>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <DialogFooter>
              <Button type="submit">Delete</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
