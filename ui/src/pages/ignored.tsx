import { useLoaderData } from "react-router-dom";
import { IgnoredLog } from "../models";

export default function Ignored() {
  const ignored: IgnoredLog[] = useLoaderData() as any;
  return (
    <div className="grid grid-cols-2">
      {ignored.map((ignore) => (
        <h1 key={ignore.preview}>{ignore.preview}</h1>
      ))}
    </div>
  );
}
