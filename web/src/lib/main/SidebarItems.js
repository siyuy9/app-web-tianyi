const SidebarItems = {
  admin: {
    asRoot: false,
    items: [],
  },
  user: {
    asRoot: false,
    items: [
      {
        label: "Profile",
        to: { name: "user_profile" },
      },
      {
        label: "Preferences",
        to: { name: "user_appearance" },
      },
    ],
  },
  help: {
    asRoot: false,
    items: [
      {
        label: "API",
        icon: "pi pi-key",
        to: { name: "help_swagger" },
      },
      {
        label: "Source code",
        icon: "pi pi-external-link",
        url: "https://gitlab.com/kongrentian-group/tianyi",
      },
    ],
  },
  project: {
    asRoot: true,
    items: [
      {
        label: "Workload",
        items: [
          {
            label: "Pipelines",
            to: { name: "project_pipelines" },
          },
        ],
      },
      {
        label: "Settings",
        items: [
          {
            label: "Repository",
            //to: { name: "project_pipelines" },
          },
        ],
      },
    ],
  },
};
export default SidebarItems;
