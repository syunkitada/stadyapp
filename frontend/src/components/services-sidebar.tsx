"use client";

import * as React from "react";
import {
  AudioWaveform,
  BookOpen,
  Bot,
  Command,
  Frame,
  GalleryVerticalEnd,
  Map,
  PieChart,
  Settings2,
  SquareTerminal,
} from "lucide-react";

import { NavMain } from "@/components/nav-main";
import { NavProjects } from "@/components/nav-projects";
import { NavUser } from "@/components/nav-user";
import { ProjectSwitcher } from "@/components/project-switcher";
import { Logo } from "@/components/logo";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";

import useAuth from "../hooks/useAuth";

// This is sample data.
const data = {
  user: {
    name: "shadcn",
    email: "m@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  navMain: [
    {
      title: "IAM",
      url: "#",
      icon: SquareTerminal,
      isActive: true,
      items: [
        {
          title: "Domains",
          url: "/services/iam/domains",
        },
        {
          title: "Users",
          url: "/services/iam/users",
        },
        {
          title: "Teams",
          url: "/services/iam/teams",
        },
        {
          title: "Organizations",
          url: "/services/iam/organizations",
        },
        {
          title: "Projects",
          url: "/services/iam/projects",
        },
      ],
    },
  ],
};

export function ServicesSidebar({
  ...props
}: React.ComponentProps<typeof Sidebar>) {
  const { user, logout } = useAuth();
  console.log("user", user.data.user.name);
  console.log("user", user.data.projects);
  const projects = user.data.projects;

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <Logo />
      </SidebarHeader>
      <SidebarHeader>
        <ProjectSwitcher selectedProject={null} projects={projects} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
