import React, { useEffect, useState } from "react";
import axios from "axios";
import { toLength } from "../utils";
import {
  Form,
  Outlet,
  redirect,
  useLoaderData,
  useNavigate,
} from "react-router-dom";
import LogContainer from "../components/LogContainer";

export interface LogCount {
  id: number;
  title: string;
  message: string;
  num: number;
  logType: string;
}

export interface ILoaderData {
  logData: LogCount[];
  search: string;
}

function Index() {
  const router = useNavigate();
  // const [logData, setLogData] = useState<LogCount[]>([]);
  const { logData, search }: ILoaderData = useLoaderData() as any;

  useEffect(() => {
    (document.getElementById("search") as any).value = search;
  }, [search]);

  if (!logData) {
    return <div></div>;
  }

  return (
    <div className="grid grid-cols-2">
      <div>
        <h1 className="font-semibold text-3xl">Logs</h1>
        <Form id="search-form" role="search">
          <input
            id="search"
            name="search"
            className="border-2 border-blue-500 rounded-sm mx-1 my-1"
            type="text"
            defaultValue={search}
          />
        </Form>
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
