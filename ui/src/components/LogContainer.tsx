import { Link } from "react-router-dom";
import { LogCount } from "../pages";
import { ErLog } from "../types";
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
        )} max-w-3xl rounded-sm my-2 cursor-pointer mx-1 px-1.5 py-1`}
      >
        <a className="font-bold">{log.title}</a>
        <p>{toLength(log.message, 100)}</p>
        <p>{log.num}</p>
      </div>
    </LinkWithQuery>
  );
}
