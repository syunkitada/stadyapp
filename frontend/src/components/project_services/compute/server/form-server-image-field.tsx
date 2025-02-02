"use client";

import { getGlanceImages } from "@/clients/compute/sdk.gen";
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

export function FormServerImageField({ form }: { form: any }) {
  const [open, setOpen] = React.useState(false);

  const {
    isPending: imagesIsPending,
    isError: imagesIsError,
    data: imagesData,
    error: imagesError,
  } = useQuery({
    queryKey: ["images"],
    queryFn: getGlanceImages,
  });

  console.log(
    "images",
    imagesIsPending,
    imagesIsError,
    imagesData,
    imagesError,
  );

  if (imagesIsPending) {
    return <CenterLoader />;
  }

  if (imagesIsError) {
    console.error(imagesError);
    return <div>Error</div>;
  }

  if (imagesData.status != 200) {
    console.error(imagesData);
    return <div>Unexpected Status</div>;
  }

  if (!imagesData.data || !imagesData.data.images) {
    return <div>No images found</div>;
  }

  const images = imagesData.data.images;

  return (
    <FormField
      control={form.control}
      name="image"
      render={({ field }) => (
        <FormItem className="flex flex-col">
          <FormLabel>Image</FormLabel>
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
                    ? images.find((image) => image.id === field.value)?.name
                    : "Select image"}
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
                    {images.map((image) => (
                      <CommandItem
                        value={image.id}
                        key={image.id}
                        onSelect={() => {
                          form.setValue("image", image.id);
                          setOpen(false);
                        }}
                      >
                        {image.name}
                        <Check
                          className={cn(
                            "ml-auto",
                            image.id === field.value
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
            This is the image that will be used in the dashboard.
          </FormDescription>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
