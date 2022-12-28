<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";

const canvasRef = ref<HTMLCanvasElement>();

const props = defineProps({
  width: {
    type: Number,
    required: true,
  },
  height: {
    type: Number,
    required: true,
  },
  base64: {
    type: String,
    required: true,
  },
});

const renderImage = () => {
  const img = new Image();
  img.src = props.base64;

  // TODO: Find out why this is required for successful render
  setTimeout(() => {
    const ctx = canvasRef.value!.getContext("2d");
    ctx!.drawImage(img, 0, 0);
  }, 0);
};

onMounted(() => {
  renderImage();
});

watch(props, renderImage);
</script>

<template>
  <div class="h-100 w-100 d-flex justify-center align-center">
    <canvas ref="canvasRef" :width="props.width" :height="props.height"></canvas>
  </div>
</template>

<style scoped>
canvas {
  max-width: 100%;
  max-height: 100%;
}
</style>
