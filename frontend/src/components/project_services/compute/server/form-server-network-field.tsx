"use client";

import { getNeutronNetworks } from "@/clients/compute/sdk.gen";
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

export function FormServerNetworkField({ form }: { form: any }) {
  const [open, setOpen] = React.useState(false);

  const {
    isPending: networksIsPending,
    isError: networksIsError,
    data: networksData,
    error: networksError,
  } = useQuery({
    queryKey: ["networks"],
    queryFn: getNeutronNetworks,
  });

  console.log(
    "networks",
    networksIsPending,
    networksIsError,
    networksData,
    networksError,
  );

  if (networksIsPending) {
    return <CenterLoader />;
  }

  if (networksIsError) {
    console.error(networksError);
    return <div>Error</div>;
  }

  if (networksData.status != 200) {
    console.error(networksData);
    return <div>Unexpected Status</div>;
  }

  if (!networksData.data || !networksData.data.networks) {
    return <div>No networks found</div>;
  }

  const networks = networksData.data.networks;

  return (
    <FormField
      control={form.control}
      name="network"
      render={({ field }) => (
        <FormItem className="flex flex-col">
          <FormLabel>Network</FormLabel>
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
                    ? networks.find((network) => network.id === field.value)
                        ?.name
                    : "Select network"}
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
                    {networks.map((network) => (
                      <CommandItem
                        value={network.id}
                        key={network.id}
                        onSelect={() => {
                          form.setValue("network", network.id);
                          setOpen(false);
                        }}
                      >
                        {network.name}
                        <Check
                          className={cn(
                            "ml-auto",
                            network.id === field.value
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
            This is the network that will be used in the dashboard.
          </FormDescription>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
