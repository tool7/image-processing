<script lang="ts" setup>
import { ref } from "vue";
import draggable from "vuedraggable";

import { main } from "../../wailsjs/go/models";
import OperationBuilder from "./OperationBuilder.vue";
import { useImageProcessing } from "../composables/image-processing";
import { imageOperationSelectItems, ImageOperationType } from "../types/image";

const {
  addImageOperation,
  removeImageOperation,
  updateImageOperation,
  replaceImageOperation,
  moveImageOperation,
  toggleImageOperation,
  processImage,
  isLoading,
} = useImageProcessing();

const operations = ref<Array<{ id: number; operation: main.ImageOperation; isEnabled: boolean }>>([]);

const dragOptions = {
  animation: 200,
  group: "description",
  disabled: false,
  ghostClass: "ghost",
};

// TODO: Use nanoid ?
let id = 0;

const onAddOperation = async (type: ImageOperationType) => {
  const operation = new main.ImageOperation({ type, level: 1, tint: { r: 0, g: 0, b: 255 } });
  operations.value.push({ id, operation, isEnabled: true });

  id++;

  try {
    await addImageOperation(operation);
    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const onRemoveOperation = async (index: number) => {
  try {
    await removeImageOperation(index);
    operations.value.splice(index, 1);

    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const onToggleOperation = async (index: number, isEnabled: boolean) => {
  try {
    await toggleImageOperation(index, !isEnabled);
    operations.value[index].isEnabled = !isEnabled;

    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const onOperationChange = async (index: number, type: ImageOperationType, level?: number, tint?: main.TintRGB) => {
  const changedOperation = operations.value[index].operation;

  try {
    if (changedOperation.type === type) {
      changedOperation.level = level;
      changedOperation.tint = tint;

      await updateImageOperation(index, changedOperation);
    } else {
      const newOperation = new main.ImageOperation({ type, level, tint });
      await replaceImageOperation(index, newOperation);

      operations.value[index].operation.type = type;
    }

    await processImage();
  } catch (err) {
    console.log(err);
  }
};

const onDragEnd = async ({ oldIndex, newIndex }: { oldIndex: number; newIndex: number }) => {
  try {
    await moveImageOperation(oldIndex, newIndex);
    await processImage();
  } catch (err) {
    console.log(err);
  }
};
</script>

<template>
  <v-menu transition="slide-y-transition">
    <template v-slot:activator="{ props }">
      <v-btn
        v-bind="props"
        variant="tonal"
        size="small"
        prepend-icon="fas fa-plus"
        :rounded="0"
        :disabled="isLoading"
        class="mb-5"
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

  <draggable
    v-model="operations"
    v-bind="dragOptions"
    item-key="id"
    handle=".reorder-handle"
    class="d-flex"
    @end="onDragEnd"
  >
    <template #item="{ element, index }">
      <OperationBuilder
        :initial-operation-type="element.operation.type"
        :is-enabled="element.isEnabled"
        class="mr-4"
        @change="(type, level, tint) => onOperationChange(index, type, level, tint)"
        @remove="() => onRemoveOperation(index)"
        @toggle="() => onToggleOperation(index, element.isEnabled)"
      />
    </template>
  </draggable>
</template>
