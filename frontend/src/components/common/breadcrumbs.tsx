import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { useMatches, isMatch, Link } from "@tanstack/react-router";

export const Breadcrumbs = () => {
  const matches = useMatches();

  if (matches.some((match) => match.status === "pending")) return null;

  console.log("DEBUG matches", matches);

  const matchesWithCrumbs = matches.filter((match) =>
    isMatch(match, "loaderData.crumb"),
  );

  console.log("DEBUG matchesWithCrumbs", matchesWithCrumbs);

  return (
    <Breadcrumb>
      <BreadcrumbList>
        {matchesWithCrumbs.map((match, i) => (
          <>
            <BreadcrumbItem key={i}>
              <BreadcrumbLink href="#">
                <Link from={match.fullPath}>{match.loaderData?.crumb}</Link>
              </BreadcrumbLink>
            </BreadcrumbItem>
            {i + 1 < matchesWithCrumbs.length ? <BreadcrumbSeparator /> : null}
          </>
        ))}
      </BreadcrumbList>
    </Breadcrumb>
  );
};
