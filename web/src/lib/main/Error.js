import EventBus from "./EventBus";

function send(error) {
  EventBus.emit("app-toast-add", error);
}

export default function AddErrorMessages(error, handle) {
  console.log(error);
  if (!handle) {
    handle = send;
  }

  if (!error.response) {
    handle({
      severity: "error",
      summary: error.message,
      detail: error,
      life: 5000,
    });
    return;
  }
  handle({
    severity: "error",
    summary: error.message,
    detail: error.response.data.error,
    life: 5000,
  });
}
