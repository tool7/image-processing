const loadProject = () => {
  console.log("load project");
};

const saveProject = () => {
  console.log("save project");
};

const exportPng = async (imageBase64: string) => {
  const link = document.createElement("a");

  link.style.display = "none";
  link.href = imageBase64;
  link.download = "image-processing-result.png";

  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

export function useFileManager() {
  return {
    loadProject,
    saveProject,
    exportPng,
  };
}
