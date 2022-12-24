<script lang="ts" setup>
import { onMounted, onUnmounted } from "vue";
import { getImageData, openImageFileSelector } from "../helpers";

const emit = defineEmits<{
  (e: "image-select", imageData: ImageData): void;
}>();

const onDrop = async (e: DragEvent) => {
  const file = e.dataTransfer?.files[0];
  if (!file) {
    return;
  }

  const imageData = await getImageData(file);
  emit("image-select", imageData);
};

const onFileSelect = async (file: File) => {
  const imageData = await getImageData(file);
  emit("image-select", imageData);
};

const onDropZoneClick = () => {
  openImageFileSelector(onFileSelect);
};

const preventDefaults = (e: Event) => {
  e.preventDefault();
};

const events = ["dragenter", "dragover", "dragleave", "drop"];

onMounted(() => {
  events.forEach((eventName) => {
    document.body.addEventListener(eventName, preventDefaults);
  });
});

onUnmounted(() => {
  events.forEach((eventName) => {
    document.body.removeEventListener(eventName, preventDefaults);
  });
});
</script>

<template>
  <div id="drop-zone" @drop.prevent="onDrop" @click="onDropZoneClick">
    <v-icon id="icon" icon="fas fa-image" size="x-large"></v-icon>
    <span id="title">Drop image here to begin</span>
  </div>
</template>

<style scoped>
#drop-zone {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  color: var(--color-dark-grey);
  background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='10' ry='10' stroke='%237B7B7BFF' stroke-width='4' stroke-dasharray='6%2c 14' stroke-dashoffset='10' stroke-linecap='square'/%3e%3c/svg%3e");
}

#drop-zone:hover {
  cursor: pointer;
  color: var(--color-white);
  background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='10' ry='10' stroke='%23F0F0F0FF' stroke-width='4' stroke-dasharray='6%2c 14' stroke-dashoffset='10' stroke-linecap='square'/%3e%3c/svg%3e");
}

#icon {
  margin-bottom: 10px;
}

#title {
  user-select: none;
}
</style>
