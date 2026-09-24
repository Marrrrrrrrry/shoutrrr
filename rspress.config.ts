import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { defineConfig } from '@rspress/core';

export default defineConfig({
  root: 'docs',
  lang: 'en',
  // 多版本：目录结构 docs/<version>/<lang>/，默认版本路由不带版本前缀。
  // 文档版本随 shoutrrr 项目（当前 0.10.0；0.9 未随文档发布，跳过）；未来新版本在 docs/ 下新增版本目录即可
  multiVersion: {
    default: 'v0.10.0',
    versions: ['v0.8', 'v0.10.0'],
  },
  // 首次访问者按浏览器语言自动跳转（原生行为，显式声明防止默认值变化）
  route: {
    localeRedirect: 'auto',
  },
  // 部署子路径：CI 里设 RSPRESS_BASE=/shoutrrr-docs/（GitHub Pages 项目站点），
  // 本地开发与根路径部署不设置该环境变量即为 '/'
  base: process.env.RSPRESS_BASE ?? '/',
  locales: [
    {
      lang: 'en',
      label: 'English',
      title: 'Shoutrrr',
      description: 'Notification library for gophers and their furry friends',
    },
    {
      lang: 'zh',
      label: '简体中文',
      title: 'Shoutrrr',
      description: 'gopher 和它们毛茸茸朋友的通知库',
    },
  ],
  globalStyles: path.join(path.dirname(fileURLToPath(import.meta.url)), 'styles/index.css'),
  logo: '/shoutrrr-180px.png',
  logoText: 'Shoutrrr',
  icon: '/favicon.ico',
  themeConfig: {
    footer: {
      message: 'Released under the MIT License.',
    },
    // TODO: 新文档仓库的远程地址确定后启用
    // editLink: {
    //   docRepoBaseUrl: 'https://github.com/<owner>/shoutrrr-docs/tree/main/docs',
    // },
    socialLinks: [
      { icon: 'github', mode: 'link', content: 'https://github.com/marrrrrrrrry/shoutrrr' },
    ],
  },
});
