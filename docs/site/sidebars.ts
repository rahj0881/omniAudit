import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */
const sidebars: SidebarsConfig = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  // tutorialSidebar: [{type: 'autogenerated', dirName: '.', exclude: ['omni', 'home']}],
  oldSidebar: [
    "home",
    // {
    //   type: "html",
    //   value: "<b><small> OMNI </small></b>"
    // },
    // {
    //   type: 'autogenerated',
    //   dirName: 'omni',
    // },
    {
      type: "category",
      label: "Omni",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "omni",
        }
      ]
    },
    // {
    //   type: "html",
    //   value: "<b><small> PROTOCOL </small></b>"
    // },
    // {
    //   type: 'autogenerated',
    //   dirName: 'protocol',
    // },
    {
      type: "category",
      label: "Protocol",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol",
        }
      ]
    },
    {
      type: "category",
      label: "Background",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "background",
        }
      ]
    },
    {
      type: "category",
      label: "Developers",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "developers",
        }
      ]
    },
  ]

  // But you can create a sidebar manually
  /*
  tutorialSidebar: [
    'intro',
    'hello',
    {
      type: 'category',
      label: 'Tutorial',
      items: ['tutorial-basics/create-a-document'],
    },
  ],
   */
};

export default sidebars;
