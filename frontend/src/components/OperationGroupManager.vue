<script lang="ts" setup>
import { ref } from "vue";
import draggable from "vuedraggable";
import OperationPicker from "./OperationPicker.vue";
import { useImageProcessing } from "../composables/image-processing";

const isDragging = ref<boolean>(false);
const operationPickers = ref([
  { id: 1, order: 1 },
  { id: 2, order: 2 },
  { id: 3, order: 3 },
]);

const dragOptions = {
  animation: 200,
  group: "description",
  disabled: false,
  ghostClass: "ghost",
};

const onDragEnd = () => {
  console.log(operationPickers.value);
};
</script>

<template>
  <draggable
    v-model="operationPickers"
    v-bind="dragOptions"
    item-key="id"
    handle=".reorder-handle"
    class="d-flex"
    @end="onDragEnd"
  >
    <template #item="{ element }">
      <OperationPicker class="mr-4" />
    </template>
  </draggable>
</template>
