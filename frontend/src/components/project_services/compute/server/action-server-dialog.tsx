"use client";

import { DialogDataTable } from "./dialog-data-table";
import { ButtonLoader } from "@/components/common/loader";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Form } from "@/components/ui/form";

export function ActionServerDialog({
  title,
  description,
  submitName,
  open,
  setOpen,
  targets,
  onSubmit,
  form,
  mutation,
}: {
  title: string;
  description: string;
  submitName: string;
  open: any;
  setOpen: any;
  targets: any[];
  form: any;
  onSubmit: any;
  mutation: any;
}) {
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[800px]">
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>

        <DialogDataTable data={targets} />

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <DialogFooter>
              {mutation.isPending ? (
                <>
                  <Button type="submit" disabled>
                    <ButtonLoader />
                    Processing
                  </Button>
                </>
              ) : (
                <Button type="submit">{submitName}</Button>
              )}
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
