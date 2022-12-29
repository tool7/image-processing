<script lang="ts" setup>
import Navbar from "./components/Navbar.vue";
import ImageViewer from "./components/ImageViewer.vue";
import { useImageProcessing } from "./composables/image-processing";
import OperationGroupManager from "./components/OperationGroupManager.vue";

const { openImageFileSelector, processedImage, isLoading } = useImageProcessing();

const onSelectImage = async () => {
  try {
    await openImageFileSelector();
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <div class="h-100 w-100">
    <Navbar id="navbar" />

    <v-progress-linear v-if="isLoading" id="loading-indicator" class="mt-8" indeterminate color="yellow-darken-2" />

    <div v-if="!processedImage && isLoading" class="h-100 w-100 d-flex justify-center align-center">
      <h3>Please wait for image to load...</h3>
    </div>

    <v-btn
      v-if="!processedImage && !isLoading"
      id="select-image-btn"
      variant="tonal"
      size="x-large"
      prepend-icon="fas fa-image"
      @click="onSelectImage"
    >
      Select image to begin
    </v-btn>

    <main v-if="processedImage && !isLoading" class="h-100 w-100 d-flex flex-column">
      <div id="image-viewer">
        <ImageViewer :width="processedImage.width" :height="processedImage.height" :base64="processedImage.base64" />
      </div>
      <div id="operation-group">
        <OperationGroupManager />
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

#image-viewer {
  max-height: 60%;
}

#operation-group {
  height: 200px;
  margin-top: 20px;
}

#select-image-btn {
  top: 50%;
  transform: translateY(-50%);
}
</style>
