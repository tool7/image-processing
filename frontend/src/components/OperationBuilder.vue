<script lang="ts" setup>
import { ref, watch } from "vue";
import Slider from "@vueform/slider";

import { main } from "../../wailsjs/go/models";
import { useImageProcessing } from "../composables/image-processing";
import { ImageOperationType, imageOperationSelectItems, rgbToHex } from "../types/image";

const props = defineProps({
  initialOperationType: {
    type: Number,
    required: true,
  },
  isEnabled: {
    type: Boolean,
    required: true,
  },
});

const emit = defineEmits<{
  (e: "change", type: ImageOperationType, level?: number, tint?: main.TintRGB, kernelSize?: number): void;
  (e: "remove"): void;
  (e: "toggle"): void;
}>();

const { isLoading } = useImageProcessing();
const selectedOperationType = ref<ImageOperationType>(props.initialOperationType);
const selectedLevel = ref<number>(1);
const selectedTint = ref<main.TintRGB>({ r: 0, g: 0, b: 255 });
const selectedKernelSize = ref<number>(3);
const selectedColorPickerValue = ref<main.TintRGB>({ r: 0, g: 0, b: 255 });
const isColorPickerOpen = ref<boolean>(false);

const onRemove = () => emit("remove");
const onToggle = () => emit("toggle");

const onColorSelect = () => {
  selectedTint.value = selectedColorPickerValue.value;
  isColorPickerOpen.value = false;
};

const kernelSizeSliderFormat = (value: number) => {
  switch (value) {
    case 3:
      return 1;
    case 5:
      return 2;
    case 7:
      return 3;
    case 9:
      return 4;
  }
};

watch([selectedOperationType, selectedLevel, selectedTint, selectedKernelSize], (newValues, oldValues) => {
  const oldOperationType = oldValues[0];
  const newOperationType = newValues[0];

  if (oldOperationType !== newOperationType) {
    selectedLevel.value = 1;
    selectedTint.value = { r: 0, g: 0, b: 255 };
    selectedKernelSize.value = 3;
  }

  emit("change", selectedOperationType.value, selectedLevel.value, selectedTint.value, selectedKernelSize.value);
});
</script>

<template>
  <v-card :disabled="isLoading" height="100%" width="240" min-width="240" variant="tonal" :rounded="1">
    <div class="d-flex justify-space-between">
      <div>
        <v-btn
          variant="tonal"
          size="x-small"
          icon="fas fa-trash-can"
          :rounded="0"
          class="remove-btn"
          @click="onRemove"
        />
        <v-tooltip :text="isEnabled ? 'Disable' : 'Enable'" location="top">
          <template v-slot:activator="{ props }">
            <v-btn
              v-bind="props"
              variant="tonal"
              size="x-small"
              :icon="isEnabled ? 'fas fa-eye' : 'fas fa-eye-slash'"
              :rounded="0"
              class="toggle-btn"
              @click="onToggle"
            />
          </template>
        </v-tooltip>
      </div>
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
        class="operation-type-select mb-2"
      />

      <div v-if="selectedOperationType === ImageOperationType.Brightness" class="mt-4">
        <div class="text-caption">Level</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="2"
          :step="0.2"
          :format="(v: number) => v"
          show-tooltip="drag"
          class="mx-3 my-3"
        />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Contrast" class="mt-4">
        <div class="text-caption">Factor</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="4"
          :step="0.1"
          :format="(v: number) => v"
          show-tooltip="drag"
          class="mx-3 my-3"
        />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Saturation" class="mt-4">
        <div class="text-caption">Level</div>
        <Slider
          v-model="selectedLevel"
          v-bind="null"
          :min="0"
          :max="3"
          :step="0.1"
          :format="(v: number) => v"
          show-tooltip="drag"
          class="mx-3 my-3"
        />
      </div>

      <div
        v-if="selectedOperationType === ImageOperationType.Tint"
        class="mt-4 d-flex justify-space-between align-center"
      >
        <v-dialog v-model="isColorPickerOpen" :max-width="340">
          <template v-slot:activator="{ props }">
            <v-btn
              v-bind="props"
              :color="rgbToHex(selectedTint.r, selectedTint.g, selectedTint.b, 140)"
              icon="fas fa-palette"
              variant="elevated"
              size="x-small"
              class="mt-2"
            />
          </template>
          <v-card>
            <v-card-title>Choose color</v-card-title>
            <v-card-text>
              <v-color-picker
                v-model="selectedColorPickerValue"
                elevation="0"
                :modes="['rgb']"
                hide-canvas
                hide-inputs
              />
            </v-card-text>
            <v-card-actions class="d-flex justify-center">
              <v-btn variant="tonal" size="small" class="mb-3 px-4" @click="onColorSelect">Confirm</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>

        <div class="w-100">
          <div class="text-caption">Intensity</div>
          <Slider
            v-model="selectedLevel"
            v-bind="null"
            :min="0"
            :max="1"
            :step="0.01"
            :format="(v: number) => v"
            show-tooltip="drag"
            class="mx-3 my-3"
          />
        </div>
      </div>

      <div
        v-if="
          [
            ImageOperationType.BoxBlur,
            ImageOperationType.MotionBlur,
            ImageOperationType.Sharpen,
            ImageOperationType.Emboss,
          ].includes(selectedOperationType)
        "
        class="mt-4"
      >
        <div class="text-caption">Strength</div>
        <Slider
          v-model="selectedKernelSize"
          v-bind="null"
          :min="3"
          :max="9"
          :step="2"
          :format="kernelSizeSliderFormat"
          show-tooltip="drag"
          class="mx-3 my-3"
        />
      </div>
    </v-card-item>
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

.remove-btn:hover {
  background-color: red;
}

.toggle-btn {
  border-bottom-right-radius: 6px !important;
}

.operation-type-select :deep(.v-input__details) {
  display: none !important;
}

.v-color-picker :deep(.v-color-picker-preview) {
  margin-bottom: 0px !important;
}
</style>
