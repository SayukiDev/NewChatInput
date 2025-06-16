import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { setupLayouts } from 'virtual:generated-layouts'
import { createPinia } from 'pinia'
import App from './App.vue'

// Vuetify (import before global styles)
import vuetify from './plugins/vuetify'

// Global styles (import after Vuetify)
import './style.css'

// Use auto-generated routes from unplugin-vue-router
import { routes } from 'vue-router/auto-routes'

// Create router instance with auto-generated routes and layouts
const router = createRouter({
  history: createWebHistory(),
  routes: setupLayouts(routes),
})

// Create Pinia instance
const pinia = createPinia()

const app = createApp(App)

app.use(pinia)
app.use(vuetify)
app.use(router)

app.mount('#app')
