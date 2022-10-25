import { ErLog } from "./models";

export const toLength = (text: string, len: number) => {
  return text.length > len ? text.substring(0, len - 3) + "..." : text;
};

export const sortLogs = (logs: ErLog[], sortDate: string) => {
  if (sortDate === "ascending") {
    return logs.sort(function (a, b) {
      return (
        new Date(b.extraData.timestamp).getDate() -
        new Date(a.extraData.timestamp).getDate()
      );
    });
  } else {
    return logs.sort(function (a, b) {
      return a.extraData.timestamp - b.extraData.timestamp;
    });
  }
};
