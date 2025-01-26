"use client";

import { useParams } from "@tanstack/react-router";

import { Logo } from "@/components/logo";

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
      title: "Project",
      url: "#",
      icon: SquareTerminal,
      isActive: true,
      items: [
        {
          title: "Detail",
          url: "/projects/$projectId/project/detail",
        },
      ],
    },
    {
      title: "Compute",
      url: "#",
      icon: SquareTerminal,
      isActive: true,
      items: [
        {
          title: "Server",
          url: "/projects/$projectId/compute/server",
        },
        {
          title: "Image",
          url: "/projects/$projectId/compute/image",
        },
        {
          title: "Network",
          url: "/projects/$projectId/compute/network",
        },
      ],
    },
    {
      title: "Monitoring",
      url: "#",
      icon: Bot,
      items: [
        {
          title: "Dashboard",
          url: "#",
        },
      ],
    },
  ],
};

export function ProjectServicesSidebar({
  ...props
}: React.ComponentProps<typeof Sidebar>) {
  const { user, logout } = useAuth();
  console.log("user", user.data.user.name);
  console.log("user", user.data.projects);
  const projects = user.data.projects;

  const { projectId } = useParams({ strict: false });

  const selectedIndex = projects.findIndex(
    (project) => project.id === projectId,
  );

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <Logo />
      </SidebarHeader>
      <SidebarHeader>
        <ProjectSwitcher selectedIndex={selectedIndex} projects={projects} />
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
