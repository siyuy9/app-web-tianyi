export const AdminSidebarItems = [
  {
    label: "API",
    items: [
      {
        label: "Swagger",
        icon: "pi pi-key",
        to: { name: "admin_swagger" },
      },
    ],
  },
];

export const BrowseSidebarItems = [
  {
    label: "Projects",
    icon: "pi pi-folder",
    to: { name: "browse_projects" },
  },
  {
    label: "Groups",
    icon: "pi pi-share-alt",
    to: { name: "browse_groups" },
  },
];

export const BrowseSidebarItemsWrapper = [
  {
    label: "Browse",
    items: BrowseSidebarItems,
  },
];
