<script lang="ts" setup>
import { ref, watch } from "vue";

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

watch([selectedOperationType, selectedLevel, selectedTint], () => {
  emit("change", selectedOperationType.value, selectedLevel.value, selectedTint.value);
});

const onRemove = () => {
  emit("remove");
};
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

      <div v-if="selectedOperationType === ImageOperationType.Brightness">
        <div class="text-caption">Level</div>
        <v-slider v-model="selectedLevel" :min="0" :max="2" :step="0.2" color="grey" thumb-label thumb-color="white" />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Contrast">
        <div class="text-caption">Factor</div>
        <v-slider v-model="selectedLevel" :min="0" :max="9" :step="0.5" color="grey" thumb-label thumb-color="white" />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Saturation">
        <div class="text-caption">Level</div>
        <v-slider v-model="selectedLevel" :min="0" :max="3" :step="0.2" color="grey" thumb-label thumb-color="white" />
      </div>

      <div v-if="selectedOperationType === ImageOperationType.Tint">
        <div class="text-caption">Intensity</div>
        <v-slider v-model="selectedLevel" :min="0" :max="3" :step="0.2" color="grey" thumb-label thumb-color="white" />
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

.remove-btn {
  justify-self: flex-end;
  align-self: flex-end;
  border-top-left-radius: 6px !important;
}

.remove-btn:hover {
  background-color: red;
}
</style>