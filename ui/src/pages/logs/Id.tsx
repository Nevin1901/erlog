import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import LoadingSpinner from "../../components/LoadingSpinner";
import { ErLog } from "../../types";
import Code from "../../components/Code";

export default function Id() {
  const [logs, setLogs] = useState<ErLog[]>([]);
  const params = useParams();

  useEffect(() => {
    const doWork = async () => {
      const { data } = await axios.post(
        `http://127.0.0.1:8080/logs/${params.id}`
      );

      setLogs(data);
    };
    doWork();
  }, []);
  if (logs.length === 0) {
    return <LoadingSpinner />;
  } else {
    return (
      <div>
        <a href="/logs">Go Back</a>
        <button className="ml-2">Ignore Logs Like This</button>
        {logs.map((log) => (
          <div key={log.ID}>
            <h1 className="font-semibold text-2xl">{log.title}</h1>

            <Code code={log.message} language={"javascript"}></Code>

            <h1 className="font-semibold text-2xl">Extra Info</h1>
            <Code
              code={JSON.stringify(log.extraData, null, 2)}
              language={"javascript"}
            />
          </div>
        ))}
      </div>
    );
  }
}
