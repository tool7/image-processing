import { main } from "../../wailsjs/go/models";
import { ImageOperationDraggableItem } from "./image";

export interface ProjectState {
  originalImage: main.ProcessedImage; // TODO: Improve type
  operations: Array<ImageOperationDraggableItem>;
}
