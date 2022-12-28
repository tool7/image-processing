<script lang="ts" setup>
import Navbar from "./components/Navbar.vue";
import ImageViewer from "./components/ImageViewer.vue";
import { useImageProcessing } from "./composables/image-processing";

const { openImageFileSelector, processedImage, resetAppState, isLoading } = useImageProcessing();

const onSelectImage = async () => {
  try {
    openImageFileSelector();
  } catch (err) {
    console.log(err);
  }
};

const onResetAll = async () => {
  try {
    resetAppState();
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <div class="h-100 w-100">
    <Navbar id="navbar" />

    <div v-if="!processedImage && isLoading" class="h-100 w-100 d-flex justify-center align-center">
      <v-progress-linear id="loading-indicator" class="mt-8" indeterminate color="yellow-darken-2" />
      <span>Please wait for image to load...</span>
    </div>
    <main v-else class="h-100 w-100">
      <v-btn
        v-if="!processedImage"
        id="select-image-btn"
        variant="tonal"
        size="x-large"
        prepend-icon="fas fa-image"
        @click="onSelectImage"
      >
        Select image to begin
      </v-btn>
      <div v-else>
        <ImageViewer :width="processedImage.width" :height="processedImage.height" :base64="processedImage.base64" />
        <v-btn color="warning" variant="tonal" size="small" @click="onResetAll">Reset All</v-btn>
      </div>
    </main>
  </div>
</template>

<style scoped>
#navbar {
  position: fixed;
  top: 0;
  left: 0;
}

#loading-indicator {
  position: fixed;
  top: 0;
  left: 0;
}

main {
  padding: 52px 20px 20px 20px;
}

#select-image-btn {
  top: 50%;
  transform: translateY(-50%);
}
</style>
