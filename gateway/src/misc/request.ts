export class corsResponse extends Response {
    constructor(body?: any, init?: any) {
      super(body, init);
      this.headers.set("Access-Control-Allow-Origin", "*");
      this.headers.set("Access-Control-Allow-Methods", "OPTIONS, GET, PUT, POST, DELETE");
      this.headers.set("Access-Control-Allow-Headers", "Content-Type, Authorization, authorization, Cache-Control");
      // this.headers.set("Cache-Control", "private, max-age=60");
    }
  }