<script setup lang="ts">
import { WindowMinimise } from '@runtime'
// import { QuitSteam } from '@backend';

const quit = () => {
  WindowMinimise()
  // QuitSteam()
  console.log('hello')
}
</script>

<template>
  <div style="background-color: red; width: 100px; height: 80px">
    Hello
    <el-button @click="quit">
      确认
    </el-button>
  </div>
</template>
