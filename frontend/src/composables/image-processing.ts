import { ref, readonly } from "vue";
import { nanoid } from "nanoid";

import { main } from "../../wailsjs/go/models";
import {
  OpenImageFileSelector,
  ProcessImage,
  GetOriginalImage,
  SetOriginalImage,
  ResetAppState,
  AppendImageOperation,
  RemoveImageOperationAtIndex,
  UpdateImageOperationAtIndex,
  ReplaceImageOperationAtIndex,
  MoveImageOperation,
  ToggleImageOperation,
  RotateImageBy90Deg,
  MirrorImageVertically,
  MirrorImageHorizontally,
} from "../../wailsjs/go/main/App";
import { ImageOperationDraggableItem } from "../types/image";

const isLoading = ref<boolean>(false);
const processedImage = ref<main.Base64Image | undefined>();
const operationDraggableItems = ref<Array<ImageOperationDraggableItem>>([]);

const setIsLoading = (value: boolean) => {
  isLoading.value = value;
};

const openImageFileSelector = async () => {
  setIsLoading(true);

  try {
    const isFileSelected = await OpenImageFileSelector();
    if (!isFileSelected) {
      setIsLoading(false);
      return;
    }

    const result = await ProcessImage(0);
    processedImage.value = result;
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

const getOriginalImage = async () => {
  return await GetOriginalImage();
};

const setOriginalImage = async (image: main.Base64Image) => {
  const rawBase64 = image.base64.split(",")[1];
  await SetOriginalImage(rawBase64);
};

const addImageOperation = async (operation: main.ImageOperation) => {
  await AppendImageOperation(operation);
  operationDraggableItems.value.push({ id: nanoid(), operation, isEnabled: operation.isEnabled });
};

const removeImageOperation = async (index: number) => {
  await RemoveImageOperationAtIndex(index);
  operationDraggableItems.value.splice(index, 1);
};

const updateImageOperation = async (index: number, level?: number, tint?: main.TintRGB, kernelSize?: number) => {
  const operation = operationDraggableItems.value[index].operation;
  operation.level = level;
  operation.tint = tint;
  operation.kernelSize = kernelSize;

  await UpdateImageOperationAtIndex(index, operation);
};

const replaceImageOperation = async (index: number, operation: main.ImageOperation) => {
  await ReplaceImageOperationAtIndex(index, operation);
  operationDraggableItems.value[index].operation.type = operation.type;
};

const moveImageOperation = async (oldIndex: number, newIndex: number) => {
  if (oldIndex === newIndex) {
    return;
  }
  await MoveImageOperation(oldIndex, newIndex);
};

const toggleImageOperation = async (index: number) => {
  await ToggleImageOperation(index);

  const { operation } = operationDraggableItems.value[index];
  operation.isEnabled = !operation.isEnabled;
};

const rotateImageBy90Deg = async () => {
  await RotateImageBy90Deg();
};

const mirrorImageVertically = async () => {
  await MirrorImageVertically();
};

const mirrorImageHorizontally = async () => {
  await MirrorImageHorizontally();
};

const processImage = async (indexToExecuteFrom: number = 0) => {
  if (indexToExecuteFrom < 0) {
    indexToExecuteFrom = 0;
  }

  setIsLoading(true);

  try {
    const result = await ProcessImage(indexToExecuteFrom);
    processedImage.value = result;
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

const resetAppState = async () => {
  await ResetAppState();

  setIsLoading(false);
  processedImage.value = undefined;
  operationDraggableItems.value = [];
};

export function useImageProcessing() {
  return {
    isLoading: readonly(isLoading),
    processedImage: readonly(processedImage),
    operationDraggableItems,
    getOriginalImage,
    setOriginalImage,
    openImageFileSelector,
    addImageOperation,
    removeImageOperation,
    updateImageOperation,
    replaceImageOperation,
    moveImageOperation,
    toggleImageOperation,
    rotateImageBy90Deg,
    mirrorImageVertically,
    mirrorImageHorizontally,
    processImage,
    resetAppState,
  };
}
