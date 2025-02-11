"use client";

import { DialogDataTable } from "./dialog-data-table";
import { ButtonLoader } from "@/components/common/loader";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogClose,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Form } from "@/components/ui/form";
import { ACTION_STATUS } from "@/hooks/useCompute";

export function ActionServerDialog({
  title,
  description,
  submitName,
  open,
  setOpen,
  actionTargets,
  onSubmit,
  form,
}: {
  title: string;
  description: string;
  submitName: string;
  open: any;
  setOpen: any;
  actionTargets: any[];
  form: any;
  onSubmit: any;
  mutation: any;
}) {
  let isProcessing = false;
  let isProcessed = false;
  let processed = 0;

  for (const [_, target] of actionTargets.entries()) {
    if (target.status == ACTION_STATUS.PROCESSING) {
      isProcessing = true;
      break;
    } else if (target.status == ACTION_STATUS.PROCESSED) {
      processed += 1;
    } else if (target.status == ACTION_STATUS.ERROR) {
      processed += 1;
    }
  }

  if (processed == actionTargets.length) {
    isProcessed = true;
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[800px]">
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>

        <DialogDataTable data={actionTargets} />

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <DialogFooter className="sm:justify-between">
              <DialogClose asChild>
                <Button type="button" variant="secondary">
                  Close
                </Button>
              </DialogClose>

              {isProcessing ? (
                <>
                  <Button type="submit" disabled>
                    {ACTION_STATUS.PROCESSING}
                    <span>
                      <ButtonLoader />
                    </span>
                  </Button>
                </>
              ) : (
                <>
                  {isProcessed ? (
                    <Button type="submit" disabled>
                      {ACTION_STATUS.PROCESSED}
                    </Button>
                  ) : (
                    <Button type="submit">{submitName}</Button>
                  )}
                </>
              )}
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
