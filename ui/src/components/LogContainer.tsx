import { LogCount } from "../models";
import { toLength } from "../utils";
import LinkWithQuery from "./LinkWithQuery";

interface IProps {
  log: LogCount;
}

function getBorder(logType: string) {
  switch (logType) {
    case "info":
      return "border-none";
    case "error":
      return "border-red-500";
    case "warning":
      return "border-yellow-500";
    default:
      return "";
  }
}

/**
 * Put delete and ignore buttons on top right
 * @param param0
 * @returns
 */
export default function LogContainer({ log }: IProps) {
  return (
    <LinkWithQuery to={`/logs/${log.id}`}>
      <div
        key={log.message}
        className={`border-2 bg-white ${getBorder(
          log.logType
        )} max-w-2xl rounded-md my-1.5 cursor-pointer px-1 py-1 hover:bg-gray-100`}
      >
        <p className="font-medium text-base">
          {log.title} <span className="text-xs text-slate-500">{log.num}</span>
        </p>
        <p className="text-sm text-slate-700">{toLength(log.message, 100)}</p>
      </div>
    </LinkWithQuery>
  );
}
