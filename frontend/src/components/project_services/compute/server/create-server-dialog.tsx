"use client";

import { FormServerFlavorField } from "./form-server-flavor-field";
import { FormServerImageField } from "./form-server-image-field";
import { FormServerNameField } from "./form-server-name-field";
import { FormServerNetworkField } from "./form-server-network-field";
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
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      image: "",
      flavor: "",
      network: "",
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("dialog onSubmit", values);
  }

  return (
    <Dialog>
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
