import { ref, readonly } from "vue";

import { main } from "../../wailsjs/go/models";
import {
  OpenImageFileSelector,
  ProcessImage,
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
const originalImage = ref<main.ProcessedImage | undefined>();
const processedImage = ref<main.ProcessedImage | undefined>();
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

    originalImage.value = result;
    processedImage.value = result;
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

const setOriginalImageBase64 = async (image: main.ProcessedImage) => {
  // TODO: Improve
  const base64 = image.base64.substring(image.base64.indexOf("base64,") + 7);
  await SetOriginalImage(base64);
  originalImage.value = image;
};

const addImageOperation = async (operation: main.ImageOperation) => {
  await AppendImageOperation(operation);
};

const removeImageOperation = async (index: number) => {
  await RemoveImageOperationAtIndex(index);
};

const updateImageOperation = async (index: number, operation: main.ImageOperation) => {
  await UpdateImageOperationAtIndex(index, operation);
};

const replaceImageOperation = async (index: number, operation: main.ImageOperation) => {
  await ReplaceImageOperationAtIndex(index, operation);
};

const moveImageOperation = async (oldIndex: number, newIndex: number) => {
  if (oldIndex === newIndex) {
    return;
  }
  await MoveImageOperation(oldIndex, newIndex);
};

const toggleImageOperation = async (index: number, enable: boolean) => {
  await ToggleImageOperation(index, enable);
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
    originalImage: readonly(originalImage),
    processedImage: readonly(processedImage),
    setOriginalImageBase64,
    operationDraggableItems,
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
