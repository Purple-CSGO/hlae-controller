<script setup lang="ts">
import * as backend from '@backend'
import * as runtime from '@runtime'
import { EventsOn } from "@runtime";

import useClipboard from 'vue-clipboard3'
const { toClipboard } = useClipboard()

// App
const port = ref('54321')
const cmd = ref('echo hello')
const logfile = ref('cfg\\a_hlae_control.log')
const csgoPath = ref('D:\\SteamLibrary\\steamapps\\common\\Counter-Strike Global Offensive\\csgo')
// const csgoPath = ref('C:\\Program Files (x86)\\Steam\\steamapps\\common\\Counter-Strike Global Offensive\\csgo')
const res = ref([])
const started = ref(false)
const folded = ref(false)

// 启动服务器
function startServer() {
  backend.StartServer(port.value).then((res)=>{
    started.value = true
    alert("启动服务成功")
    // success("启动服务成功", '')
  }).catch((err:any) => {
    alert("启动服务失败" + err)
  })
}

function startLog() {
  backend.StartLog(csgoPath.value, logfile.value).then((res)=>{
    alert("监听日志成功" + res)
  }).catch((err:any) => {
    alert("监听日志失败" + err)
  })
}

function endLog() {
  backend.EndLog().then((res)=>{
    started.value = true
    console.log(res)
    alert("关闭日志成功" + res)
    // success("启动服务成功", '')
  }).catch((err:any) => {
    alert("关闭日志失败" + err)
  })
}

// 发送指令
function send() {
  if (started.value === false) {
    alert("发送指令失败 服务尚未开启")
    return
  }
  backend.Send(cmd.value)
  res.value.unshift('] ' + cmd.value)
  cmd.value = ''
}

function clearLog() {
  backend.ClearLog(csgoPath.value, logfile.value)
}

EventsOn('res', result => {
  res.value.unshift(result)
})
</script>

<template>
  <div id="window" ref="root" class="w-full h-full flex flex-col gap-2 p-2 select-none overflow-clip ">
    <section class="flex flex-col gap-1 text-left" v-show="folded === false">
      <h2 class="font-bold">服务端口</h2>

      <div class="flex gap-6">
        <div class="flex gap-2">
          <input v-model="port" type="text" class="p-2 rounded outline-none" />
          <button class="px-3 py-2 bg-white rounded active:scale-95" @click="startServer()">启动服务</button>
        </div>

        <div class="flex gap-2">
          <input v-model="logfile" type="text" class="p-2 rounded outline-none" />
          <button class="px-3 py-2 bg-white rounded active:scale-95" @click="startLog()">监听日志</button>
          <button class="px-3 py-2 bg-white rounded active:scale-95" @click="clearLog()">清除日志</button>
        </div>
      </div>
    </section>

    <section class="flex flex-col gap-1 text-left" v-show="folded === false">
      <h2 class="font-bold">CSGO路径 <span class="text-sm font-normal">不含exe</span></h2>
      <input v-model="csgoPath" type="text" class="p-2 rounded outline-none" />
    </section>

    <section class="flex flex-col gap-1 text-left" v-show="folded === false">
      <h2 class="font-bold">连接指令 <span class="text-sm font-normal">点击复制，分开一次输入一块</span></h2>

      <div class="bg-white rounded select-text flex-grow">
        <p class="text-sm p-2  hover:bg-gray-100 cursor-pointer active:scale-95" @click="toClipboard(`mirv_pgl url &quot;ws://localhost:${port}/mirv&quot;`)">mirv_pgl url "ws://localhost:{{port}}/mirv";</p>
        <p class="text-sm p-2  hover:bg-gray-100 cursor-pointer active:scale-95" @click="toClipboard('mirv_pgl start; mirv_pgl datastart;con_logfile ' + logfile + ';')">mirv_pgl start; mirv_pgl datastart; con_logfile {{logfile}};</p>
      </div>
    </section>

    <section class="flex flex-col gap-1 text-left" v-show="folded === false">
      <h2 class="font-bold">关闭连接 <span class="text-sm font-normal">否则关闭工具导致游戏无响应</span></h2>
      <p class="bg-white rounded p-2 text-sm cursor-pointer active:scale-95 hover:bg-gray-100" @click="toClipboard('mirv_pgl dataStop; mirv_pgl stop;')">
        mirv_pgl dataStop; mirv_pgl stop;
      </p>
    </section>

    <section class="flex flex-col flex-grow gap-1 text-left overflow-hidden">
      <div class="font-bold flex gap-2">
        <p class="flex-grow">输出</p>
        <button @click="res = []" class="text-sm font-normal bg-white px-1.5 py-0.5 rounded outline-none border-b active:scale-95">清空</button>
        <button @click="folded = folded === false" class="text-sm font-normal bg-white px-1.5 py-0.5 rounded outline-none border-b active:scale-95">折叠</button>
      </div>
      <div class="bg-white flex-grow flex flex-col-reverse rounded p-1 overflow-auto scroll">
        <p v-for="i in res" class="select-text cursor-pointer hover:bg-gray-100 p-1" @click="toClipboard(i)">{{i}}</p>
      </div>
    </section>

    <section class="flex gap-2" >
      <input v-model="cmd" type="text" class="flex-grow outline-none px-2 py-3 rounded" @keydown.enter="send()" />
      <button class="px-4 py-1 bg-white rounded active:scale-95" @click="send()">输入</button>
    </section>
  </div>
</template>

<style lang="scss">
* {
  @apply transition-all duration-200;
}
</style>
