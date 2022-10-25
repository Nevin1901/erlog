import React, { useEffect, useState } from "react";
import {
  Form,
  Outlet,
  useLoaderData,
  useLocation,
  useNavigate,
  useNavigation,
} from "react-router-dom";
import Split from "react-split";
import IDSplit from "../components/IdSplit";
import LogContainer from "../components/LogContainer";
import { ILoaderData } from "../models";

function Index() {
  const router = useNavigate();
  const location = useLocation();
  // const [logData, setLogData] = useState<LogCount[]>([]);
  const { logData, search }: ILoaderData = useLoaderData() as any;

  useEffect(() => {
    console.log(location);
    (document.getElementById("search") as any).value = search;
  }, [search]);

  if (!logData) {
    return <div></div>;
  }

  return (
    <IDSplit>
      <div className="mx-1.5 overflow-y-hidden">
        <h1 className="font-semibold text-xl">Logs</h1>
        <Form id="search-form" role="search">
          <input
            id="search"
            name="search"
            className="border-2 border-blue-500 rounded-sm my-1"
            type="text"
            defaultValue={search}
          />
        </Form>
        {logData.map((data) => (
          <LogContainer log={data} />
        ))}
        <h1>Hello</h1>
      </div>
      <div className="h-screen bg-white">
        <Outlet />
      </div>
    </IDSplit>
  );
}

export default Index;
