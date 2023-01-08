export namespace main {
	
	export class Base64Image {
	    width: number;
	    height: number;
	    base64: string;
	
	    static createFrom(source: any = {}) {
	        return new Base64Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.width = source["width"];
	        this.height = source["height"];
	        this.base64 = source["base64"];
	    }
	}
	export class TintRGB {
	    r: number;
	    g: number;
	    b: number;
	
	    static createFrom(source: any = {}) {
	        return new TintRGB(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.r = source["r"];
	        this.g = source["g"];
	        this.b = source["b"];
	    }
	}
	export class ImageOperation {
	    type: number;
	    level?: number;
	    tint?: TintRGB;
	    kernelSize?: number;
	    isEnabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ImageOperation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.level = source["level"];
	        this.tint = this.convertValues(source["tint"], TintRGB);
	        this.kernelSize = source["kernelSize"];
	        this.isEnabled = source["isEnabled"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

