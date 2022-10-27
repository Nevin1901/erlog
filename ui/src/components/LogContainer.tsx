import { LogCount } from "../models";
import { toLength } from "../utils";
import LinkWithQuery from "./LinkWithQuery";

interface IProps {
  log: LogCount;
}

function getBorder(logType: string) {
  // switch (logType) {
  //   case "info":
  //     return "text-black";
  //   case "error":
  //     return "text-red-500";
  //   case "warning":
  //     return "text-yellow-500";
  //   default:
  //     return "text-black";
  // }
  return "";
}

export default function LogContainer({ log }: IProps) {
  return (
    <LinkWithQuery to={`/logs/${log.id}`}>
      <div
        key={log.message}
        className={`bg-white ${getBorder(
          log.logType
        )} max-w-3xl rounded-sm my-1.5 cursor-pointer px-1 py-1`}
      >
        <p className="font-medium text-base">
          {log.title} <span className="text-xs text-slate-500">{log.num}</span>
        </p>
        <p className="text-sm text-slate-700">{toLength(log.message, 100)}</p>
      </div>
    </LinkWithQuery>
  );
}
