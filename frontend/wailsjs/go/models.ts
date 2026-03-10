export namespace aivis {
	
	export class Styles {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Styles(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class Speaker {
	    name: string;
	    speaker_uuid: string;
	    styles: Styles[];
	    version: string;
	
	    static createFrom(source: any = {}) {
	        return new Speaker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.speaker_uuid = source["speaker_uuid"];
	        this.styles = this.convertValues(source["styles"], Styles);
	        this.version = source["version"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace options {
	
	export class Config {
	    send_port: number;
	    recv_port: number;
	    enable_typing_msg: boolean;
	    realtime: boolean;
	    msg_keeping: boolean;
	    voice_control: boolean;
	    tts: boolean;
	    tts_option: tts.Option;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.send_port = source["send_port"];
	        this.recv_port = source["recv_port"];
	        this.enable_typing_msg = source["enable_typing_msg"];
	        this.realtime = source["realtime"];
	        this.msg_keeping = source["msg_keeping"];
	        this.voice_control = source["voice_control"];
	        this.tts = source["tts"];
	        this.tts_option = this.convertValues(source["tts_option"], tts.Option);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace pages {
	
	export class SelectedSpackerRsp {
	    SpackerId: number;
	
	    static createFrom(source: any = {}) {
	        return new SelectedSpackerRsp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SpackerId = source["SpackerId"];
	    }
	}

}

export namespace tts {
	
	export class Device {
	    Name: string;
	    Id: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Id = source["Id"];
	    }
	}
	export class Option {
	    Baseurl: string;
	    run: boolean;
	    log: string;
	    path: string;
	    args: string[];
	    now_spacker: number;
	    device: string;
	    cache: boolean;
	    cache_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Option(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Baseurl = source["Baseurl"];
	        this.run = source["run"];
	        this.log = source["log"];
	        this.path = source["path"];
	        this.args = source["args"];
	        this.now_spacker = source["now_spacker"];
	        this.device = source["device"];
	        this.cache = source["cache"];
	        this.cache_path = source["cache_path"];
	    }
	}

}

