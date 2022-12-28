import { ref, readonly } from "vue";
import { main } from "../../wailsjs/go/models";
import { OpenImageFileSelector, ProccessImage } from "../../wailsjs/go/main/App";

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

    const result = await ProccessImage();
    setProcessedImage(result);
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

const resetAppState = () => {
  setIsLoading(false);
  setProcessedImage(undefined);
};

const processImage = async () => {
  setIsLoading(true);

  try {
    const result = await ProccessImage();
    setProcessedImage(result);
  } catch (err) {
    throw err;
  } finally {
    setIsLoading(false);
  }
};

export function useImageProcessing() {
  return {
    isLoading: readonly(isLoading),
    processedImage: readonly(processedImage),
    openImageFileSelector,
    processImage,
    resetAppState,
  };
}
