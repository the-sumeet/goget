export namespace main {
	
	export class HttpResponse {
	    status: string;
	    statusCode: number;
	    header: {[key: string]: string[]};
	    body: string;
	    time: number;
	
	    static createFrom(source: any = {}) {
	        return new HttpResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.statusCode = source["statusCode"];
	        this.header = source["header"];
	        this.body = source["body"];
	        this.time = source["time"];
	    }
	}

}

