import {defineConfig} from "umi";

export default defineConfig({
    npmClient: 'yarn',
    routes: [
        {path: '/', redirect: '/sign-in'},
        {path: '/sign-in', component: 'signIn', layout: false},
        {path: '/sign-up', component: 'signUp', layout: false},

        // {path: '/total', component: 'account'},
        {path: '/bill', component: 'bill'},
        {path: '/account', component: 'account'},
        // {path: '/report', component: 'account'},
        // {path: '/settings', component: 'settings'}

    ],
    proxy: {
        '/api': {
            target: 'http://127.0.0.1:8989',
            changeOrigin: true,
            pathRewrite: {'^/': ''},
        },
    },
});
