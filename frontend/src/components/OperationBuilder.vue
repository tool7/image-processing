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
  (e: "change", type: ImageOperationType, level?: number, tint?: main.TintRGB): void;
  (e: "remove"): void;
  (e: "toggle"): void;
}>();

const { isLoading } = useImageProcessing();
const selectedOperationType = ref<ImageOperationType>(props.initialOperationType);
const selectedLevel = ref<number>(1);
const selectedTint = ref<main.TintRGB>({ r: 0, g: 0, b: 255 });
const selectedColorPickerValue = ref<main.TintRGB>({ r: 0, g: 0, b: 255 });
const isColorPickerOpen = ref<boolean>(false);

const onRemove = () => emit("remove");
const onToggle = () => emit("toggle");

const onColorSelect = () => {
  selectedTint.value = selectedColorPickerValue.value;
  isColorPickerOpen.value = false;
};

watch([selectedOperationType, selectedLevel, selectedTint], (newValues, oldValues) => {
  const oldOperationType = oldValues[0];
  const newOperationType = newValues[0];

  if (oldOperationType !== newOperationType) {
    selectedLevel.value = 1;
    selectedTint.value = { r: 0, g: 0, b: 255 };
  }

  emit("change", selectedOperationType.value, selectedLevel.value, selectedTint.value);
});
</script>

<template>
  <v-card :disabled="isLoading" height="100%" width="240" max-width="240" variant="tonal" :rounded="1">
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
        <v-btn
          variant="tonal"
          size="x-small"
          :icon="isEnabled ? 'fas fa-eye' : 'fas fa-eye-slash'"
          :rounded="0"
          class="toggle-btn"
          @click="onToggle"
        />
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
          :max="1"
          :step="0.01"
          :format="(v: number) => v"
          show-tooltip="drag"
        />

        <v-dialog v-model="isColorPickerOpen" :max-width="340">
          <template v-slot:activator="{ props }">
            <v-btn
              id="tint-picker-btn"
              v-bind="props"
              :color="rgbToHex(selectedTint.r, selectedTint.g, selectedTint.b, 140)"
              prepend-icon="fas fa-palette"
              variant="flat"
              size="small"
              class="mb-2"
            >
              Color
            </v-btn>
          </template>
          <v-card>
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
              <v-btn variant="tonal" class="mb-3 px-4" @click="onColorSelect">Confirm</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
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

.controls > .slider-horizontal {
  margin: 8px 12px;
}

#tint-picker-btn {
  margin-top: 20px;
}

.v-color-picker :deep(.v-color-picker-preview) {
  margin-bottom: 0px !important;
}
</style>
