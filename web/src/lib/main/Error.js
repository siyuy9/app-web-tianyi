import EventBus from "./EventBus";

export default function AddErrorMessages(error, toast) {
  console.log(error);
  var send = toast
    ? toast.add
    : (error) => EventBus.emit("app-toast-add", error);
  if (!error.response) {
    send({
      severity: "error",
      summary: error.message,
      detail: error,
      life: 5000,
    });
    return;
  }
  send({
    severity: "error",
    summary: error.message,
    detail: error.response.data.error,
    life: 5000,
  });
}
