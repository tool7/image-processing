<script lang="ts" setup>
import { onMounted, ref } from "vue";

const canvasRef = ref<HTMLCanvasElement>();

const props = defineProps({
  imageData: {
    type: Object,
    required: true,
  },
});

onMounted(() => {
  const ctx = canvasRef.value!.getContext("2d");
  ctx!.putImageData(props.imageData as ImageData, 0, 0);
});
</script>

<template>
  <div id="image-viewer">
    <canvas ref="canvasRef" :width="props.imageData.width" :height="props.imageData.height"></canvas>
  </div>
</template>

<style scoped>
#image-viewer {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

canvas {
  max-width: 100%;
  max-height: 100%;
}
</style>
