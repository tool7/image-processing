<script lang="ts" setup>
import { ref } from "vue";
import draggable from "vuedraggable";

import { main } from "../../wailsjs/go/models";
import OperationBuilder from "./OperationBuilder.vue";
import { useImageProcessing } from "../composables/image-processing";
import { Color, imageOperationSelectItems, ImageOperationType } from "../types/image";

const { addImageOperation, removeImageOperation, processImage } = useImageProcessing();

const operations = ref<Array<{ id: number; operation: main.ImageOperation }>>([]);

const dragOptions = {
  animation: 200,
  group: "description",
  disabled: false,
  ghostClass: "ghost",
};

// TODO: Use nanoid ?
let id = 0;

const onAddOperation = async (operationType: ImageOperationType) => {
  const operation = new main.ImageOperation({
    type: operationType,
    level: 1,
    tint: { r: 255, g: 0, b: 255, a: 255 },
  });
  operations.value.push({ id, operation });

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

const onOperationChange = async (index: number, type: ImageOperationType, level?: number, tint?: Color) => {
  console.log(index, type, level, tint);
};

const onDragEnd = () => {
  console.log(operations.value);
};
</script>

<template>
  <v-menu transition="slide-y-transition">
    <template v-slot:activator="{ props }">
      <v-btn v-bind="props" variant="tonal" size="small" prepend-icon="fas fa-plus" :rounded="0" class="mb-5">
        Add New Operation
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
        class="mr-4"
        @change="(type, level, tint) => onOperationChange(index, type, level, tint)"
        @remove="() => onRemoveOperation(index)"
      />
    </template>
  </draggable>
</template>
