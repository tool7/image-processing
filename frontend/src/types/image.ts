export interface Color {
  r: number;
  g: number;
  b: number;
  a: number;
}

export enum ImageOperationType {
  Brightness = "brightness",
  Contrast = "contrast",
  Saturation = "saturation",
  Tint = "tint",
  Greyscale = "greyscale",
  Negative = "negative",
  Sepia = "sepia",
  Emboss = "emboss",
  EdgesVertical = "edges-vertical",
  EdgesHorizontal = "edges-horizontal",
  MirrorVertical = "mirror-vertical",
  MirrorHorizontal = "mirror-horizontal",
  RotateBy90 = "rotate-90",
  RotateBy180 = "rotate-180",
  RotateBy270 = "rotate-270",
}

export class ImageOperation {
  public level?: number;
  public tint?: any;

  constructor(public type: ImageOperationType, level?: number, tint?: any) {
    switch (type) {
      case ImageOperationType.Brightness:
      case ImageOperationType.Contrast:
      case ImageOperationType.Saturation:
        this.level = level;
        break;
      case ImageOperationType.Tint:
        this.level = level;
        this.tint = tint;
      default:
        break;
    }
  }
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
