"use client"
import { Button, message } from "antd";
import { MessageError } from "./MessageError";
export default function CustomError() {
  const [messageApi, contextHolder] = message.useMessage();
  const callAPI = (type: number) => {
    try {
      if (type === 1) {
        throw new MessageError("消息错误", 200);
      } else {
        throw new Error("系统异常错误");
      }
    } catch (e) {
      console.error(e);
      if (e instanceof MessageError) {
        messageApi.warning(e.message);
      } else {
        messageApi.error(e.message);
      }
    }
  };
  return (
    <div>
      {contextHolder}
      <div>自定义错误类</div>
      <div className="mt16">
        <Button type="primary" onClick={() => callAPI(1)}>
          消息错误
        </Button>
      </div>
      <div className="mt16">
        <Button type="primary" onClick={() => callAPI(2)}>
          系统错误
        </Button>
      </div>
    </div>
  );
}
