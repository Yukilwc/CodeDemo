export class MessageError extends Error {
  code = 0;
  constructor(message: string, code: number) {
    super(message);
    this.name = "MessageError";
    this.code = code;
    // 在这里捕获堆栈信息，并将this作为targetObject，this.constructor作为constructorOpt
    Error.captureStackTrace(this, this.constructor);
  }
}
