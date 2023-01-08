import { main } from "../../wailsjs/go/models";

export interface ProjectState {
  originalImage: main.Base64Image;
  operations: Array<main.ImageOperation>;
}
