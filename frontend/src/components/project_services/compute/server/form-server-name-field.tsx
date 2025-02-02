"use client";

import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

export function FormServerNameField({ form }: { form: any }) {
  return (
    <FormField
      control={form.control}
      name="name"
      render={({ field }) => (
        <FormItem>
          <FormLabel>Name</FormLabel>
          <FormControl>
            <Input placeholder="testvm" {...field} />
          </FormControl>
          <FormDescription>Server name.</FormDescription>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
