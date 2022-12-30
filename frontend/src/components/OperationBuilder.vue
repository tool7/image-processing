<script lang="ts" setup>
import { ref, watch } from "vue";
import Slider from "@vueform/slider";

import { useImageProcessing } from "../composables/image-processing";
import { Color, ImageOperationType, imageOperationSelectItems } from "../types/image";

const props = defineProps({
  initialOperationType: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits<{
  (e: "change", type: ImageOperationType, level?: number, tint?: Color): void;
  (e: "remove"): void;
}>();

const { isLoading } = useImageProcessing();
const selectedOperationType = ref<ImageOperationType>(props.initialOperationType);
const selectedLevel = ref<number>(1);
const selectedTint = ref<Color>({ r: 0, g: 0, b: 0, a: 255 });

const onRemove = () => emit("remove");

watch([selectedOperationType, selectedLevel, selectedTint], (newValues, oldValues) => {
  const oldOperationType = oldValues[0];
  const newOperationType = newValues[0];

  if (oldOperationType !== newOperationType) {
    selectedLevel.value = 1;
    selectedTint.value = { r: 0, g: 0, b: 0, a: 255 };
  }

  emit("change", selectedOperationType.value, selectedLevel.value, selectedTint.value);
});
</script>

<template>
  <v-card :disabled="isLoading" height="100%" width="240" max-width="240" variant="tonal" :rounded="1">
    <div class="d-flex justify-end">
      <v-btn variant="plain" size="small" icon="fas fa-grip-lines" :rounded="0" class="reorder-handle" />
    </div>

    <v-card-item>
      <v-select
        v-model="selectedOperationType"
        :items="imageOperationSelectItems"
        item-title="label"
        item-value="type"
        label="Choose operation..."
        single-line
        density="compact"
        variant="solo"
      />

      <div v-if="selectedOperationType === ImageOperationType.Brightness" class="controls">
        <div class="text-caption">Level</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="2"
          :step="0.2"
          :format="(v: number) => v"
          show-tooltip="drag"
        />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Contrast" class="controls">
        <div class="text-caption">Factor</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="4"
          :step="0.1"
          :format="(v: number) => v"
          show-tooltip="drag"
        />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Saturation" class="controls">
        <div class="text-caption">Level</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="3"
          :step="0.1"
          :format="(v: number) => v"
          show-tooltip="drag"
        />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Tint" class="controls">
        <div class="text-caption">Intensity</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="3"
          :step="0.1"
          :format="(v: number) => v"
          show-tooltip="drag"
        />
        <v-btn prepend-icon="fas fa-palette" variant="outlined" size="small" class="mb-2">{{ selectedTint }}</v-btn>
      </div>
    </v-card-item>

    <div class="d-flex justify-end">
      <v-btn variant="tonal" size="x-small" icon="fas fa-trash-can" :rounded="0" class="remove-btn" @click="onRemove" />
    </div>
  </v-card>
</template>

<style scoped>
.v-card--disabled {
  color: var(--color-dark-grey);
}

.reorder-handle {
  cursor: grab;
  border-bottom-left-radius: 6px !important;
}

.controls > .slider-horizontal {
  margin: 8px 12px;
}

.remove-btn {
  justify-self: flex-end;
  align-self: flex-end;
  border-top-left-radius: 6px !important;
}

.remove-btn:hover {
  background-color: red;
}
</style>
