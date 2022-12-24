<script lang="ts" setup>
import { WindowMinimise, WindowToggleMaximise, Quit, WindowIsMaximised } from "../../wailsjs/runtime/runtime";

const onMinimise = () => WindowMinimise();
const onToggleMaximise = () => WindowToggleMaximise();
const onQuit = () => Quit();
</script>

<template>
  <div id="navbar" style="--wails-draggable: drag">
    <span id="title">Go Image Processing</span>

    <div>
      <v-btn
        class="btn"
        :width="32"
        :height="32"
        variant="text"
        icon="fas fa-minus"
        :rounded="0"
        style="--wails-draggable: none"
        @click="onMinimise"
      />
      <v-btn
        class="btn"
        :width="32"
        :height="32"
        variant="text"
        icon="fas fa-stop"
        :rounded="0"
        style="--wails-draggable: none"
        @click="onToggleMaximise"
      />
      <v-btn
        id="btn-quit"
        class="btn"
        :width="32"
        :height="32"
        variant="text"
        icon="fas fa-sharp fa-xmark"
        :rounded="0"
        style="--wails-draggable: none"
        @click="onQuit"
      />
    </div>
  </div>
</template>

<style scoped>
#navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 32px;
  border-bottom: 1px solid rgba(40, 40, 40, 1);
  font-size: 12px;
  margin: auto;
}

#title {
  padding: 10px;
  text-align: center;
  color: var(--color-white);
  cursor: default;
  user-select: none;
}

#btn-quit:hover {
  background-color: red;
}

.btn {
  font-size: 10px !important;
}
</style>
