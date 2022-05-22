import {defineConfig} from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    {path: '/sign-in', component: '@/pages/signIn/index'},
    {path: '/sign-up', component: '@/pages/signUp/index'},
    {
      path: '/', component: '@/layouts/index',
      wrappers: [
        '@/wrappers/auth'
      ],
      routes: [
        {path: '/account', component: '@/pages/account/index'}
      ]
    },
  ],
  fastRefresh: {},
  "theme": {
    "primary-color": "#FA541C",
  },
  proxy: {
    '/api/**/*': {
      target: 'http://127.0.0.1:8989',
      changeOrigin: true,
      pathRewrite: {'^/': '/'},
    },
  }
});
