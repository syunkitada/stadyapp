"use client";

import { AlertDialog, AlertDialogContent } from "@/components/ui/alert-dialog";
import PuffLoader from "react-spinners/PuffLoader";

export function ButtonLoader() {
  return <PuffLoader color="#05bcf6" size={20} />;
}

export function Loader() {
  return <PuffLoader color="#05bcf6" size={80} />;
}

export function DialogLoader() {
  return (
    <AlertDialog defaultOpen={true}>
      <AlertDialogContent>
        <div className={"flex flex-col space-y-2 items-center"}>
          <Loader />
        </div>
      </AlertDialogContent>
    </AlertDialog>
  );
}

export function CenterLoader() {
  return (
    <div className={"flex flex-col space-y-2 items-center"}>
      <Loader />
    </div>
  );
}
