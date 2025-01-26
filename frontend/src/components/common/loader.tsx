"use client";

import reactLogo from "../../assets/react.svg";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";

import PuffLoader from "react-spinners/PuffLoader";

export function Loader() {
  return <PuffLoader color="#19fc74" size={80} />;
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
