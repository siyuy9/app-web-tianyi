import EventBus from "./EventBus";

export default async function AddErrorMessages(error, toast) {
  await error;
  console.error(error);
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
