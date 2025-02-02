"use client";

import { getNovaFlavorsDetail } from "@/clients/compute/sdk.gen";
import { CenterLoader } from "@/components/common/loader";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import { Check, ChevronsUpDown } from "lucide-react";
import * as React from "react";

export function FormServerFlavorField({ form }: { form: any }) {
  const [open, setOpen] = React.useState(false);

  const {
    isPending: flavorsIsPending,
    isError: flavorsIsError,
    data: flavorsData,
    error: flavorsError,
  } = useQuery({
    queryKey: ["flavors"],
    queryFn: getNovaFlavorsDetail,
  });

  console.log(
    "flavors",
    flavorsIsPending,
    flavorsIsError,
    flavorsData,
    flavorsError,
  );

  if (flavorsIsPending) {
    return <CenterLoader />;
  }

  if (flavorsIsError) {
    console.error(flavorsError);
    return <div>Error</div>;
  }

  if (flavorsData.status != 200) {
    console.error(flavorsData);
    return <div>Unexpected Status</div>;
  }

  if (!flavorsData.data || !flavorsData.data.flavors) {
    return <div>No flavors found</div>;
  }

  const flavors = flavorsData.data.flavors;

  return (
    <FormField
      control={form.control}
      name="flavor"
      render={({ field }) => (
        <FormItem className="flex flex-col">
          <FormLabel>Flavor</FormLabel>
          <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
              <FormControl>
                <Button
                  variant="outline"
                  role="combobox"
                  className={cn(
                    "w-[200px] justify-between",
                    !field.value && "text-muted-foreground",
                  )}
                >
                  {field.value
                    ? flavors.find((flavor) => flavor.id === field.value)?.name
                    : "Select flavor"}
                  <ChevronsUpDown className="opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent className="w-[200px] p-0">
              <Command>
                <CommandInput
                  placeholder="Search framework..."
                  className="h-9"
                />
                <CommandList>
                  <CommandEmpty>No framework found.</CommandEmpty>
                  <CommandGroup>
                    {flavors.map((flavor) => (
                      <CommandItem
                        value={flavor.id}
                        key={flavor.id}
                        onSelect={() => {
                          form.setValue("flavor", flavor.id);
                          setOpen(false);
                        }}
                      >
                        {flavor.name}
                        <Check
                          className={cn(
                            "ml-auto",
                            flavor.id === field.value
                              ? "opacity-100"
                              : "opacity-0",
                          )}
                        />
                      </CommandItem>
                    ))}
                  </CommandGroup>
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
          <FormDescription>
            This is the flavor that will be used in the dashboard.
          </FormDescription>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
