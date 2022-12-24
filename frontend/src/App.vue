<script lang="ts" setup>
import { reactive } from "vue";
import Navbar from "./components/Navbar.vue";
import ImageFileSelector from "./components/ImageFileSelector.vue";
import ImageViewer from "./components/ImageViewer.vue";

interface AppState {
  selectedImageData?: ImageData;
}

const state = reactive<AppState>({
  selectedImageData: undefined,
});

const onImageSelect = (image: ImageData) => {
  state.selectedImageData = image;
};

const onClearImage = () => {
  state.selectedImageData = undefined;
};
</script>

<template>
  <div id="app-container">
    <Navbar />

    <main>
      <ImageFileSelector v-if="!state.selectedImageData" @image-select="onImageSelect" />
      <ImageViewer v-else :image-data="state.selectedImageData" />

      <v-btn color="warning" variant="tonal" size="small" @click="onClearImage">Clear Image</v-btn>
    </main>
  </div>
</template>

<style scoped>
#app-container {
  width: 100%;
  height: 100%;
}

main {
  width: 100%;
  height: 70%;
  padding: 20px;
}
</style>
