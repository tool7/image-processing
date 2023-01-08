import { readonly, ref } from "vue";

import { GetUserSelectedProjectFileContent } from "../../wailsjs/go/main/App";
import { ProjectState } from "../types/project";
import { useImageProcessing } from "./image-processing";

const isLoading = ref<boolean>(false);
const isSaving = ref<boolean>(false);

const {
  processedImage,
  operationDraggableItems,
  addImageOperation,
  resetAppState,
  getOriginalImage,
  setOriginalImage,
  processImage,
} = useImageProcessing();

const convertUrlToFile = async (url: string, filename: string) => {
  const res = await fetch(url);
  const buf = await res.arrayBuffer();
  return new File([buf], filename);
};

const downloadFile = (file: File) => {
  const link = document.createElement("a");

  link.style.display = "none";
  link.href = URL.createObjectURL(file);
  link.download = file.name;

  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(link.href);
};

const loadProject = async () => {
  const selectedProjectFileContent = await GetUserSelectedProjectFileContent();
  if (!selectedProjectFileContent) {
    return;
  }

  const projectState: ProjectState = JSON.parse(selectedProjectFileContent);
  isLoading.value = true;

  try {
    await resetAppState();
    await setOriginalImage(projectState.originalImage);

    for (const operation of projectState.operations) {
      await addImageOperation(operation);
    }

    await processImage();
  } catch (err) {
    console.log(err);
  } finally {
    isLoading.value = false;
  }
};

const saveProject = async () => {
  isSaving.value = true;

  try {
    const originalImage = await getOriginalImage();
    const operations = operationDraggableItems.value.map(({ operation }) => operation);
    const projectState: ProjectState = { originalImage, operations };
    const projectStateString = JSON.stringify(projectState);

    const file = new File([projectStateString], "image-processing-project.goimp");
    downloadFile(file);
  } catch (err) {
    console.log(err);
  } finally {
    isSaving.value = false;
  }
};

const exportPng = async () => {
  const imageBase64 = processedImage.value?.base64;
  if (!imageBase64) {
    return;
  }

  const file = await convertUrlToFile(imageBase64, "image-processing-result.png");
  downloadFile(file);
};

export function useProjectManager() {
  return {
    isLoading: readonly(isLoading),
    isSaving: readonly(isSaving),
    loadProject,
    saveProject,
    exportPng,
  };
}
