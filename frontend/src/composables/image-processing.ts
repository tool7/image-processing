import { ref, readonly } from "vue";

import { main } from "../../wailsjs/go/models";
import {
  OpenImageFileSelector,
  ProcessImage,
  ResetAppState,
  AppendImageOperation,
  RemoveImageOperationAtIndex,
  UpdateImageOperationAtIndex,
  ReplaceImageOperationAtIndex,
} from "../../wailsjs/go/main/App";

const isLoading = ref<boolean>(false);
const processedImage = ref<main.ProcessedImage | undefined>();

const setIsLoading = (value: boolean) => {
  isLoading.value = value;
};

const setProcessedImage = (image?: main.ProcessedImage) => {
  processedImage.value = image;
};

const openImageFileSelector = async () => {
  setIsLoading(true);

  try {
    const isFileSelected = await OpenImageFileSelector();
    if (!isFileSelected) {
      setIsLoading(false);
      return;
    }

    const result = await ProcessImage();
    setProcessedImage(result);
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
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

const processImage = async () => {
  setIsLoading(true);

  try {
    const result = await ProcessImage();
    setProcessedImage(result);
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

const resetAppState = async () => {
  await ResetAppState();

  setIsLoading(false);
  setProcessedImage(undefined);
};

export function useImageProcessing() {
  return {
    isLoading: readonly(isLoading),
    processedImage: readonly(processedImage),
    openImageFileSelector,
    addImageOperation,
    removeImageOperation,
    updateImageOperation,
    replaceImageOperation,
    processImage,
    resetAppState,
  };
}
