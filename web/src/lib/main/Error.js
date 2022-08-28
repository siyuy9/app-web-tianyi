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
  for (var [key, value] of Object.entries(error.response.data.errors)) {
    send({
      severity: "error",
      summary: error.message,
      detail: key === "body" ? value : `${key}: ${value}`,
      life: 5000,
    });
  }
}