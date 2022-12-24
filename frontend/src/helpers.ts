export const getImageData = async (blob: Blob): Promise<ImageData> => {
  const bitmap = await createImageBitmap(blob);
  const [width, height] = [bitmap.width, bitmap.height];

  const canvas = document.createElement("canvas");
  canvas.width = width;
  canvas.height = height;

  const ctx = canvas.getContext("2d");
  if (!ctx) {
    throw new Error("Canvas context is null");
  }

  ctx.drawImage(bitmap, 0, 0);
  return ctx.getImageData(0, 0, width, height);
};

export const openImageFileSelector = (callback: (file: File) => void) => {
  const input = document.createElement("input");
  input.type = "file";
  input.accept = "image/png, image/jpeg";

  input.onchange = () => {
    if (!input.files) {
      return;
    }
    callback(input.files[0]);
  };

  input.click();
};
