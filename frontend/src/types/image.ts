import { main } from "../../wailsjs/go/models";

export enum ImageOperationType {
  Brightness = 1,
  Contrast,
  Saturation,
  Tint,
  Greyscale,
  Negative,
  Sepia,
  BoxBlur,
  MotionBlur,
  Sharpen,
  Emboss,
  EdgesHorizontal,
  EdgesVertical,
  Outline,
}

export const imageOperationSelectItems: Array<{ type: ImageOperationType; label: string }> = [
  { type: ImageOperationType.Brightness, label: "Brightness" },
  { type: ImageOperationType.Contrast, label: "Contrast" },
  { type: ImageOperationType.Saturation, label: "Saturation" },
  { type: ImageOperationType.Tint, label: "Tint" },
  { type: ImageOperationType.Greyscale, label: "Greyscale" },
  { type: ImageOperationType.Negative, label: "Negative" },
  { type: ImageOperationType.Sepia, label: "Sepia" },
  { type: ImageOperationType.BoxBlur, label: "Blur" },
  { type: ImageOperationType.MotionBlur, label: "Motion blur" },
  { type: ImageOperationType.Sharpen, label: "Sharpen" },
  { type: ImageOperationType.Emboss, label: "Emboss" },
  { type: ImageOperationType.EdgesHorizontal, label: "Horizontal edges" },
  { type: ImageOperationType.EdgesVertical, label: "Vertical edges" },
  { type: ImageOperationType.Outline, label: "Outline" },
];

export interface ImageOperationDraggableItem {
  id: string;
  operation: main.ImageOperation;
  isEnabled: boolean;
}

const colorComponentToHex = (c: number) => {
  var hex = c.toString(16);
  return hex.length == 1 ? "0" + hex : hex;
};

export const rgbToHex = (r: number, g: number, b: number, a: number) => {
  return "#" + colorComponentToHex(r) + colorComponentToHex(g) + colorComponentToHex(b) + colorComponentToHex(a);
};
