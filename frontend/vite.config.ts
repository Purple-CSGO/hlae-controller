import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import { createHtmlPlugin } from 'vite-plugin-html'
import pkg from './package.json'

// 自动路由
// import Pages from 'vite-plugin-pages'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(), // Vue SFC
    // Pages(), // Auto Routing '/pages'
    AutoImport({
      dts: true, // TS support
      imports: ['vue'],
      resolvers: [ElementPlusResolver()],
      dirs: [
        'src/store', // Pinia Store Modules
        'src/composables'
      ],
      eslintrc: {
        enabled: false
      }
    }),
    Components({
      dts: true, // TS support
      resolvers: [ElementPlusResolver()]
    }),
    createHtmlPlugin({
      minify: false,
      // entry: "src/main.js",
      template: 'index.html',
      inject: {
        data: {
          title: `${pkg.name}`
        }
      }
    })
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@runtime': path.resolve(__dirname, 'wailsjs', 'runtime'),
      '@backend': path.resolve(__dirname, 'wailsjs', 'go', 'backend', 'App'),
      '@models': path.resolve(__dirname, 'wailsjs', 'go', 'models'),
    }
  }
})
