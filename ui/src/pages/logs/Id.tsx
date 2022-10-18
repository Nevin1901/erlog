import axios from "axios";
import { useEffect, useState } from "react";
import {
  Link,
  useLoaderData,
  useNavigate,
  useNavigation,
  useParams,
} from "react-router-dom";
import LoadingSpinner from "../../components/LoadingSpinner";
import Code from "../../components/Code";
import { ErLog } from "../../models";

export default function Id() {
  // const [logs, setLogs] = useState<ErLog[]>([]);
  const logs: ErLog[] = useLoaderData() as ErLog[];
  const params = useParams();
  const router = useNavigate();
  const navigation = useNavigation();

  const deleteLog = async () => {
    if (logs.length < 1) {
      return;
    }

    const { data } = await axios.post(
      `http://127.0.0.1:8080/ignore/${logs[0].id}`
    );

    router("/logs");
  };

  useEffect(() => {
    // const doWork = async () => {
    //   const { data } = await axios.post(
    //     `http://127.0.0.1:8080/logs/${params.id}`
    //   );
    //   setLogs(data);
    // };
    // doWork();
  }, []);
  if (logs.length === 0) {
    return <LoadingSpinner />;
  } else {
    if (navigation.state === "loading") {
      return <LoadingSpinner />;
    }
    return (
      <div className="max-h-screen overflow-y-scroll">
        <Link to="/logs">
          <a href="/logs">Go Back</a>
        </Link>
        <button onClick={() => deleteLog()} className="ml-2">
          Ignore
        </button>
        {logs.map((log) => (
          <div key={log.id}>
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
