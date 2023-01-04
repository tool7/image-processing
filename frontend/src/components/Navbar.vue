<script lang="ts" setup>
import { computed } from "vue";
import { WindowMinimise, WindowToggleMaximise, Quit } from "../../wailsjs/runtime/runtime";
import { useFileManager } from "../composables/file-manager";
import { useImageProcessing } from "../composables/image-processing";
import { NavbarMenuItem } from "../types/navbar";

const { processedImage, resetAppState, isLoading } = useImageProcessing();
const { loadProject, saveProject, exportPng } = useFileManager();

const onMinimise = () => WindowMinimise();
const onToggleMaximise = () => WindowToggleMaximise();
const onQuit = () => Quit();

const menuItems = computed<Array<NavbarMenuItem>>(() => {
  return [
    {
      title: "Load Project",
      icon: "fas fa-file-import",
      isEnabled: !isLoading.value,
      onClick: () => loadProject(),
    },
    {
      title: "Save Project",
      icon: "fas fa-floppy-disk",
      isEnabled: Boolean(processedImage.value) && !isLoading.value,
      onClick: () => saveProject(),
    },
    {
      title: "Export PNG",
      icon: "fas fa-file-image",
      isEnabled: Boolean(processedImage.value) && !isLoading.value,
      onClick: () => exportPng(processedImage.value!.base64),
    },
    {
      title: "Reset All",
      icon: "fas fa-trash-can",
      isEnabled: !isLoading.value,
      onClick: () => resetAppState(),
    },
  ];
});
</script>

<template>
  <div id="navbar" style="--wails-draggable: drag">
    <v-menu transition="fade-transition">
      <template v-slot:activator="{ props }">
        <v-btn
          icon="fas fa-ellipsis-vertical"
          v-bind="props"
          size="x-small"
          variant="text"
          :rounded="0"
          style="--wails-draggable: none"
        ></v-btn>
      </template>

      <v-list>
        <v-list-item
          v-for="(menuItem, i) in menuItems"
          :key="i"
          class="menu-item"
          @click="() => menuItem.onClick()"
          :disabled="!menuItem.isEnabled"
        >
          <template v-slot:prepend>
            <v-icon :icon="menuItem.icon" size="small"></v-icon>
          </template>

          <v-list-item-title class="text-left mr-4">{{ menuItem.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

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
  background-color: var(--app-bg-color);
}

.menu-item:hover {
  cursor: pointer;
}

#btn-quit:hover {
  background-color: red;
}

.btn {
  font-size: 10px !important;
}
</style>
