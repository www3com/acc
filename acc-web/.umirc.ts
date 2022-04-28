import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  // routes: [
  //   { path: '/', component: '@/pages/index' },
  //   { path: '/account', component: '@/pages/account/index' },
  // ],
  fastRefresh: {},
});
