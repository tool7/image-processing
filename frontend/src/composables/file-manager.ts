const loadProject = () => {
  console.log("load project");
};

const saveProject = () => {
  console.log("save project");
};

const exportPng = () => {
  console.log("export png");
};

export function useFileManager() {
  return {
    loadProject,
    saveProject,
    exportPng,
  };
}
