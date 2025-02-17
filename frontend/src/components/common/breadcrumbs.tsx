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
  const breadcrumbs = [];
  for (const [index, match] of matchesWithCrumbs.entries()) {
    breadcrumbs.push(
      <BreadcrumbItem key={index}>
        <BreadcrumbLink asChild>
          <Link from={match.fullPath}>{match.loaderData?.crumb}</Link>
        </BreadcrumbLink>
      </BreadcrumbItem>,
    );
    if (index + 1 < matchesWithCrumbs.length) {
      breadcrumbs.push(<BreadcrumbSeparator key={`${index}_`} />);
    }
  }

  return (
    <Breadcrumb>
      <BreadcrumbList>{breadcrumbs}</BreadcrumbList>
    </Breadcrumb>
  );
};
