import {defineConfig} from "umi";

export default defineConfig({
    routes: [
        {path: '/', redirect: '/sign-in'},
        {path: '/sign-in', component: 'signIn'},
        {path: '/sign-up', component: 'signUp'},
        {path: '/', component: 'layout',
            routes: [
                // {path: '/total', component: 'account'},
                {path: '/bill', component: 'bill'},
                {path: '/account', component: 'account'},
                // {path: '/report', component: 'account'},
                // {path: '/settings', component: 'settings'}
            ]
        },
    ],
    npmClient: 'yarn',
    proxy: {
        '/api': {
            target: 'http://127.0.0.1:8989',
            changeOrigin: true,
            pathRewrite: {'^/': ''},
        },
    },
});
