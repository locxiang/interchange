import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  name: 'Index',
  redirect: '/userLogin',
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue'),
},
{
  path: '/userLogin',
  name: 'userLogin',
  component: import('@/view/userLogin/index.vue'),
},
{
  path: '/platform',
  name: 'platform',
  component: import('@/view/platform/index.vue'),
},
{
  path: '/admin',
  name: 'Login',
  component: () => import('@/view/login/index.vue'),
},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
