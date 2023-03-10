// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AppendImageOperation(arg1:main.ImageOperation):Promise<Error>;

export function GetOriginalImage():Promise<main.Base64Image>;

export function GetUserSelectedProjectFileContent():Promise<string>;

export function MirrorImageHorizontally():Promise<Error>;

export function MirrorImageVertically():Promise<Error>;

export function MoveImageOperation(arg1:number,arg2:number):Promise<Error>;

export function OpenImageFileSelector():Promise<boolean>;

export function ProcessImage(arg1:number):Promise<main.Base64Image>;

export function RemoveImageOperationAtIndex(arg1:number):Promise<Error>;

export function ReplaceImageOperationAtIndex(arg1:number,arg2:main.ImageOperation):Promise<Error>;

export function ResetAppState():Promise<void>;

export function RotateImageBy90Deg():Promise<Error>;

export function SetOriginalImage(arg1:string):Promise<void>;

export function ToggleImageOperation(arg1:number):Promise<Error>;

export function UpdateImageOperationAtIndex(arg1:number,arg2:main.ImageOperation):Promise<Error>;
