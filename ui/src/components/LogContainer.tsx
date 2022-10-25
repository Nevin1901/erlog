import { LogCount } from "../models";
import { toLength } from "../utils";
import LinkWithQuery from "./LinkWithQuery";

interface IProps {
  log: LogCount;
}

function getBorder(logType: string) {
  switch (logType) {
    case "info":
      return "border-gray-500";
    case "error":
      return "border-red-500";
    case "warning":
      return "border-yellow-500";
    default:
      return "border-gray-500";
  }
}

export default function LogContainer({ log }: IProps) {
  return (
    <LinkWithQuery to={`/logs/${log.id}`}>
      <div
        key={log.message}
        className={`bg-gray-100 border-2 ${getBorder(
          log.logType
        )} max-w-3xl rounded-sm my-1.5 cursor-pointer px-1 py-1`}
      >
        <a href={`logs/${log.id}`} className="font-bold text-base">
          {log.title} <span className="text-xs text-gray-500">{log.num}</span>
        </a>
        <p className="text-sm text-gray-700">{toLength(log.message, 100)}</p>
      </div>
    </LinkWithQuery>
  );
}
