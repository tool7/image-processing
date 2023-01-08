<script lang="ts" setup>
import { computed } from "vue";

import { useImageProcessing } from "./composables/image-processing";
import { useProjectManager } from "./composables/project-manager";
import Navbar from "./components/Navbar.vue";
import ImageViewer from "./components/ImageViewer.vue";
import OperationGroupManager from "./components/OperationGroupManager.vue";

const { openImageFileSelector, processedImage, isLoading: isProcessingImage } = useImageProcessing();
const { isLoading: isLoadingProject, isSaving: isSavingProject } = useProjectManager();

const isLoadingDialogOpen = computed<boolean>(() => {
  return isLoadingProject.value || isSavingProject.value;
});

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

    <v-progress-linear
      v-if="processedImage && isProcessingImage"
      id="loading-indicator"
      :height="1"
      color="blue-lighten-3"
      indeterminate
      class="mt-8"
    />

    <div v-if="!processedImage && isProcessingImage" class="h-100 w-100 d-flex justify-center align-center">
      <v-progress-circular indeterminate color="blue-lighten-3" class="mr-3" />
      <h3>Please wait for image to load...</h3>
    </div>

    <v-btn
      v-if="!processedImage && !isProcessingImage"
      id="select-image-btn"
      variant="tonal"
      size="x-large"
      prepend-icon="fas fa-image"
      @click="onSelectImage"
    >
      Select image to begin
    </v-btn>

    <main v-if="processedImage" class="h-100 w-100 d-flex flex-column">
      <div id="image-viewer">
        <ImageViewer :width="processedImage.width" :height="processedImage.height" :base64="processedImage.base64" />
      </div>
      <div id="operation-group">
        <OperationGroupManager />
      </div>
    </main>

    <v-dialog v-model="isLoadingDialogOpen" :scrim="false" persistent width="50%">
      <v-card color="gray" class="d-flex">
        <v-card-text>
          <div class="mb-3">
            <h3 v-if="isLoadingProject">Loading project...</h3>
            <h3 v-if="isSavingProject">Saving project...</h3>
          </div>
          <v-progress-linear indeterminate color="blue-lighten-3" class="mb-2"></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
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
