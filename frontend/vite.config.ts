import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import { PrimeVueResolver } from '@primevue/auto-import-resolver'

export default defineConfig({
  plugins: [
    vue(),
    Components({ resolvers: [PrimeVueResolver()] }),
  ],
  test: {
    environment: 'happy-dom',
    globals: true,
  },
})
