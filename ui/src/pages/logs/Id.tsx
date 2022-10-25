import axios from "axios";
import {
  Link,
  useLoaderData,
  useNavigate,
  useNavigation,
} from "react-router-dom";
import LoadingSpinner from "../../components/LoadingSpinner";
import Code from "../../components/Code";
import { ErLog } from "../../models";
import CrossIcon from "../../components/icons/CrossIcon";
import TrashIcon from "../../components/icons/TrashIcon";
import { Disclosure } from "@headlessui/react";
import ChevronRightIcon from "@heroicons/react/24/solid/ChevronRightIcon";
import { formatDistance } from "date-fns";

export default function Id() {
  const logs: ErLog[] = useLoaderData() as ErLog[];
  const router = useNavigate();
  const navigation = useNavigation();

  const deleteLog = async () => {
    if (logs.length < 1) {
      return;
    }

    await axios.post(`http://127.0.0.1:8080/ignore/${logs[0].id}`);

    router("/logs");
  };

  const toDate = (timeStamp: string) => {
    const current = new Date();

    const date = new Date(timeStamp);
    return formatDistance(date, current, { addSuffix: true });
  };

  if (logs.length === 0) {
    return <LoadingSpinner />;
  } else {
    if (navigation.state === "loading") {
      return <LoadingSpinner />;
    }
    return (
      <div className="max-h-screen overflow-y-scroll mt-1.5">
        <div className="flex items-center mb-1.5">
          <Link to="/logs">
            <CrossIcon />
          </Link>
          <TrashIcon onClick={() => deleteLog()} />
        </div>
        <div className="space-y-1">
          {logs.map((log) => (
            <Disclosure key={log.id} as="div">
              <Disclosure.Button className="flex items-center border-y-2 border-black w-full">
                <h1 className="font-bold">
                  {log.extraData.timestamp
                    ? toDate(log.extraData.timestamp)
                    : log.title}
                </h1>
                <ChevronRightIcon className="w-6 h-6 ui-open:rotate-90 ui-open:transform" />
              </Disclosure.Button>
              <Disclosure.Panel>
                <div key={log.id}>
                  <Code code={log.message} language={"javascript"}></Code>

                  <h1 className="font-bold">Extra Info</h1>
                  <Code
                    code={JSON.stringify(log.extraData, null, 2)}
                    language={"javascript"}
                  />
                </div>
              </Disclosure.Panel>
            </Disclosure>
          ))}
        </div>
      </div>
    );
  }
}
