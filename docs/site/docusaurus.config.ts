import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";
import remarkMath from 'remark-math';
import rehypeKatex from 'rehype-katex';

const config: Config = {
  title: "Omni Docs",
  tagline: "Documentation for the Omni Network",
  favicon: "img/favicon.svg",

  // Set the production url of your site here
  url: "https://your-docusaurus-site.example.com",
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "omni-network", // Usually your GitHub org/user name.
  projectName: "docs", // Usually your repo name.

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl: "https://github.com/omni-network/omni/docs/",
          remarkPlugins: [remarkMath],
          rehypePlugins: [rehypeKatex],
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      } satisfies Preset.Options,
    ],
  ],
  stylesheets: [
    {
      href: 'https://cdn.jsdelivr.net/npm/katex@0.13.24/dist/katex.min.css',
      type: 'text/css',
      integrity: 'sha384-odtC+0UGzzFL/6PNoE8rX/SPcQDXBJ+uRepguP4QkPCm2LBxH3FA3y+fKSiJ+AmM', // pragma: allowlist secret
      crossorigin: 'anonymous',
    },
  ],

  themeConfig: {
    // Replace with your project's social card
    image: "img/docusaurus-social-card.jpg",
    navbar: {
      title: "Omni Docs",
      logo: {
        alt: "Omni Logo",
        src: "img/logo.svg",
      },
      items: [
        {
          position: "left",
          label: "Learn",
          type: "doc",
          docId: "learn/introduction",
        },
        {
          position: "left",
          label: "Protocol",
          type: "doc",
          docId: "protocol/overview",
        },
        {
          position: "left",
          label: "Develop",
          type: "doc",
          docId: "develop/contracts",
        },
        {
          position: "left",
          label: "Operate",
          type: "doc",
          docId: "operate/run",
        },
        {
          href: "https://github.com/omni-network/omni",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      // style: "dark",
      links: [
        {
          label: "Main Site",
          href: "https://omni.network",
        },
        {
          label: "Discord",
          href: "https://discord.gg/bKNXmaX9VD",
        },
        {
          label: "Twitter",
          href: "https://twitter.com/OmniFDN",
        },
        {
          label: "Telegram",
          href: "https://t.me/omnifdn",
        },
        {
          label: "GitHub",
          href: "https://github.com/omni-network/omni/",
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} The Omni Network`,
    },
    prism: {
      theme: prismThemes.vsLight,
      darkTheme: prismThemes.vsDark,
      additionalLanguages: ["solidity"],
    },
    algolia: {
      appId: "<NEW_APP_ID>", // pragma: allowlist secret
      apiKey: "<NEW_API_KEY>", // pragma: allowlist secret
      indexName: "index-name",
      contextualSearch: true,
      searchParameters: {
        clickAnalytics: true,
        analytics: true,
        enableReRanking: true,
      },
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
