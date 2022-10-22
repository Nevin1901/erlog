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
import { ILoaderData } from "../models";

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
    <div className="flex">
      <div style={{ flexGrow: "0.5", minWidth: "800px" }}>
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
      <div
        style={{
          width: "6px",
          color: "black",
          backgroundColor: "black",
          cursor: "ew-resize",
        }}
        className="h-screen"
      ></div>
      <div style={{ flexGrow: "0.5" }} className="h-20 bg-white">
        <Outlet />
      </div>
    </div>
  );
}

export default Index;
