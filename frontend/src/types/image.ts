export enum ImageOperationType {
  Brightness = 1,
  Contrast,
  Saturation,
  Tint,
  Greyscale,
  Negative,
  Sepia,
  Emboss,
  EdgesVertical,
  EdgesHorizontal,
  MirrorVertical,
  MirrorHorizontal,
  RotateBy90,
  RotateBy180,
  RotateBy270,
}

export const imageOperationSelectItems: Array<{ type: ImageOperationType; label: string }> = [
  {
    type: ImageOperationType.Brightness,
    label: "Brightness",
  },
  {
    type: ImageOperationType.Contrast,
    label: "Contrast",
  },
  {
    type: ImageOperationType.Saturation,
    label: "Saturation",
  },
  {
    type: ImageOperationType.Tint,
    label: "Tint",
  },
  {
    type: ImageOperationType.Greyscale,
    label: "Greyscale",
  },
  {
    type: ImageOperationType.Negative,
    label: "Negative",
  },
  {
    type: ImageOperationType.Sepia,
    label: "Sepia",
  },
  {
    type: ImageOperationType.Emboss,
    label: "Emboss",
  },
  {
    type: ImageOperationType.EdgesVertical,
    label: "Vertical edges",
  },
  {
    type: ImageOperationType.EdgesHorizontal,
    label: "Horizontal edges",
  },
  {
    type: ImageOperationType.MirrorVertical,
    label: "Vertical mirror",
  },
  {
    type: ImageOperationType.MirrorHorizontal,
    label: "Horizontal mirror",
  },
  {
    type: ImageOperationType.RotateBy90,
    label: "Rotate by 90°",
  },
  {
    type: ImageOperationType.RotateBy180,
    label: "Rotate by 180°",
  },
  {
    type: ImageOperationType.RotateBy270,
    label: "Rotate by -90°",
  },
];

const colorComponentToHex = (c: number) => {
  var hex = c.toString(16);
  return hex.length == 1 ? "0" + hex : hex;
};

export const rgbToHex = (r: number, g: number, b: number, a: number) => {
  return "#" + colorComponentToHex(r) + colorComponentToHex(g) + colorComponentToHex(b) + colorComponentToHex(a);
};
