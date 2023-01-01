<script lang="ts" setup>
import { useImageProcessing } from "../composables/image-processing";

const { rotateImageBy90Deg, mirrorImageVertically, mirrorImageHorizontally, processImage, isLoading } =
  useImageProcessing();

const rotateImageBy90Degrees = async () => {
  try {
    await rotateImageBy90Deg();
    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const applyVerticalMirrorToImage = async () => {
  try {
    await mirrorImageVertically();
    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const applyHorizontalMirrorToImage = async () => {
  try {
    await mirrorImageHorizontally();
    await processImage();
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <div class="d-flex">
    <v-tooltip text="Rotate by 90°" location="bottom">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          variant="tonal"
          icon="fas fa-arrow-rotate-right"
          size="x-small"
          :disabled="isLoading"
          :rounded="0"
          @click="rotateImageBy90Degrees"
        />
      </template>
    </v-tooltip>

    <v-tooltip text="Vertical mirror" location="bottom">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          variant="tonal"
          icon="fas fa-arrow-down-up-across-line"
          size="x-small"
          :disabled="isLoading"
          :rounded="0"
          @click="applyVerticalMirrorToImage"
          class="rotate-by-90-deg"
        />
      </template>
    </v-tooltip>

    <v-tooltip text="Horizontal mirror" location="bottom">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          variant="tonal"
          icon="fas fa-arrow-down-up-across-line"
          size="x-small"
          :disabled="isLoading"
          :rounded="0"
          @click="applyHorizontalMirrorToImage"
        />
      </template>
    </v-tooltip>
  </div>
</template>
