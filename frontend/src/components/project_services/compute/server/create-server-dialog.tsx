"use client";

import { FormServerFlavorField } from "./form-server-flavor-field";
import { FormServerImageField } from "./form-server-image-field";
import { FormServerNameField } from "./form-server-name-field";
import { FormServerNetworkField } from "./form-server-network-field";
import { createNovaServer } from "@/clients/compute/sdk.gen";
import { CreateServerRequest } from "@/clients/compute/types.gen";
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

export function CreateServerDialog() {
  const [open, setOpen] = React.useState(false);

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
    mutationFn: (data: CreateServerRequest) => createNovaServer({ body: data }),
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
    const req: CreateServerRequest = {
      server: {
        name: values.name,
        min_count: 1,
        max_count: 1,
        networks: [{ uuid: values.network }],
        flavorRef: values.flavor,
        imageRef: values.image,
        block_device_mapping_v2: [
          {
            uuid: values.image,
            boot_index: 0,
            source_type: "image",
            destination_type: "local",
            delete_on_termination: true,
          },
        ],
      },
    };
    mutation.mutate(req);
    console.log("dialog onSubmit", values);
    console.log(values.name);
    console.log(values.image);
    console.log(values.flavor);
    console.log(values.network);
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">Create Server</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[800px]">
        <DialogHeader>
          <DialogTitle>Create Server</DialogTitle>
          <DialogDescription>
            Make changes to your profile here. Click save when you're done.
          </DialogDescription>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <FormServerNameField form={form} />
            <FormServerFlavorField form={form} />
            <FormServerImageField form={form} />
            <FormServerNetworkField form={form} />

            <DialogFooter>
              <Button type="submit">Create</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
