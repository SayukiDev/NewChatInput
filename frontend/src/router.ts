import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  { path: '/', redirect: '/chat' },
  { path: '/chat', component: () => import('./views/ChatView.vue') },
  { path: '/options', component: () => import('./views/OptionsView.vue') },
  { path: '/tts', component: () => import('./views/TTSView.vue') },
]

export default createRouter({ history: createWebHashHistory(), routes })
