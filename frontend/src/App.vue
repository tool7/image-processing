<script lang="ts" setup>
import { reactive } from "vue";
import Navbar from "./components/Navbar.vue";
import ImageViewer from "./components/ImageViewer.vue";
import { OpenImageFileSelector, ProccessImage } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";

interface AppState {
  processedImage?: main.ProcessedImage;
}

const state = reactive<AppState>({
  processedImage: undefined,
});

const onSelectImage = async () => {
  await OpenImageFileSelector();

  const result = await ProccessImage();
  state.processedImage = result;
};

const onClearImage = () => {
  state.processedImage = undefined;
};
</script>

<template>
  <div id="app-container">
    <Navbar />

    <main>
      <v-btn v-if="!state.processedImage" color="success" variant="flat" @click="onSelectImage"> Select Image </v-btn>
      <ImageViewer
        v-else
        :width="state.processedImage.width"
        :height="state.processedImage.height"
        :base64="state.processedImage.base64"
      />

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
