import React, { useEffect, useState } from "react";
import axios from "axios";
import { toLength } from "../utils";
import { Outlet, redirect, useLoaderData, useNavigate } from "react-router-dom";
import LogContainer from "../components/LogContainer";

export interface LogCount {
  id: number;
  title: string;
  message: string;
  num: number;
  logType: string;
}

function Index() {
  const router = useNavigate();
  // const [logData, setLogData] = useState<LogCount[]>([]);
  const logData: LogCount[] = useLoaderData() as LogCount[];

  useEffect(() => {
    // const doWork = async () => {
    //   try {
    //     const { data } = await axios.post("http://127.0.0.1:8080/count");
    //     setLogData(data);
    //   } catch (err) {
    //     console.log(err);
    //   }
    // };
    // doWork();
    console.log(logData);
  }, []);

  if (!logData) {
    return <div></div>;
  }

  return (
    <div className="grid grid-cols-2">
      <div>
        <h1 className="font-semibold text-3xl">Logs</h1>
        {logData.map((data) => (
          <LogContainer log={data} />
        ))}
        <h1>Hello</h1>
      </div>
      <div className="h-20 bg-white">
        <Outlet />
      </div>
    </div>
  );
}

export default Index;
