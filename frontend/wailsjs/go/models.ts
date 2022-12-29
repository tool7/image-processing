export namespace main {
	
	export class ImageOperation {
	    type: number;
	
	    static createFrom(source: any = {}) {
	        return new ImageOperation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	    }
	}
	export class ProcessedImage {
	    width: number;
	    height: number;
	    base64: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessedImage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.width = source["width"];
	        this.height = source["height"];
	        this.base64 = source["base64"];
	    }
	}

}

