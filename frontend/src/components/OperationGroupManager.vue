<script lang="ts" setup>
import draggable from "vuedraggable";

import { main } from "../../wailsjs/go/models";
import { useImageProcessing } from "../composables/image-processing";
import { useProjectManager } from "../composables/project-manager";
import TransformationActions from "./TransformationActions.vue";
import OperationBuilder from "./OperationBuilder.vue";
import { imageOperationSelectItems, ImageOperationType } from "../types/image";

const {
  operationDraggableItems,
  addImageOperation,
  removeImageOperation,
  updateImageOperation,
  replaceImageOperation,
  moveImageOperation,
  toggleImageOperation,
  processImage,
  isLoading: isProcessingImage,
} = useImageProcessing();
const { isSaving: isSavingProject } = useProjectManager();

const dragOptions = { animation: 200, group: "description", disabled: false, ghostClass: "ghost" };

const onAddOperation = async (type: ImageOperationType) => {
  const lastOperationIndex = operationDraggableItems.value.length - 1;
  const operation = new main.ImageOperation({
    type,
    level: 1,
    tint: { r: 0, g: 0, b: 255 },
    kernelSize: 3,
    isEnabled: true,
  });

  try {
    await addImageOperation(operation);
    await processImage(lastOperationIndex);
  } catch (err) {
    console.log(err);
  }
};

const onRemoveOperation = async (index: number) => {
  try {
    await removeImageOperation(index);

    if (operationDraggableItems.value.length === index) {
      index -= 1;
    }
    await processImage(index);
  } catch (err) {
    console.log(err);
  }
};

const onToggleOperation = async (index: number) => {
  try {
    await toggleImageOperation(index);
    operationDraggableItems.value[index].isEnabled = !operationDraggableItems.value[index].isEnabled;

    await processImage(index);
  } catch (err) {
    console.log(err);
  }
};

const onOperationChange = async (
  index: number,
  type: ImageOperationType,
  level?: number,
  tint?: main.TintRGB,
  kernelSize?: number
) => {
  const { isEnabled, type: previousType } = operationDraggableItems.value[index].operation;
  const isTypeChanged = previousType !== type;

  try {
    if (isTypeChanged) {
      const newOperation = new main.ImageOperation({ type, level, tint, kernelSize, isEnabled });
      await replaceImageOperation(index, newOperation);
    } else {
      await updateImageOperation(index, level, tint, kernelSize);
    }

    await processImage(index);
  } catch (err) {
    console.log(err);
  }
};

const onDragEnd = async ({ oldIndex, newIndex }: { oldIndex: number; newIndex: number }) => {
  const indexToProcessImageFrom = Math.min(oldIndex, newIndex);

  try {
    await moveImageOperation(oldIndex, newIndex);
    await processImage(indexToProcessImageFrom);
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <div class="d-flex justify-center align-center mb-5">
    <v-menu location="center" transition="fade-transition">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          variant="tonal"
          size="small"
          prepend-icon="fas fa-plus"
          :rounded="0"
          :disabled="isProcessingImage || isSavingProject"
          class="mr-4"
        >
          New Operation
        </v-btn>
      </template>
      <v-list>
        <v-list-item v-for="(op, i) in imageOperationSelectItems" :key="i" @click="() => onAddOperation(op.type)">
          <v-list-item-title>{{ op.label }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <TransformationActions :disabled="isProcessingImage || isSavingProject" />
  </div>

  <draggable
    v-model="operationDraggableItems"
    v-bind="dragOptions"
    item-key="id"
    handle=".reorder-handle"
    class="d-flex overflow-x-auto pb-4"
    @end="onDragEnd"
  >
    <template #item="{ element, index }">
      <OperationBuilder
        :initial-operation="element.operation"
        :is-enabled="element.isEnabled"
        class="mr-4"
        @change="(type, level, tint, kernelSize) => onOperationChange(index, type, level, tint, kernelSize)"
        @remove="() => onRemoveOperation(index)"
        @toggle="() => onToggleOperation(index)"
      />
    </template>
  </draggable>
</template>
