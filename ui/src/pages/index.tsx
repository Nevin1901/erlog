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
    <div>
      <div className="h-8 flex items-center border-b-4 px-1.5">
        <h1 className="font-bold text-xl">ErLog</h1>
      </div>
      <IDSplit>
        <div className="mx-1.5 my-1.5 overflow-y-hidden">
          {/* <h1 className="font-semibold text-xl">Logs</h1> */}
          <Form id="search-form" role="search">
            <input
              id="search"
              name="search"
              className="p-0 border-2 pl-1 border-gray-300 focus:border-gray-500 focus:ring-gray-500 rounded-sm"
              type="text"
              defaultValue={search}
            />
          </Form>
          {logData.map((data) => (
            <LogContainer log={data} />
          ))}
        </div>
        <div className="h-screen bg-white">
          <Outlet />
        </div>
      </IDSplit>
    </div>
  );
}

export default Index;
