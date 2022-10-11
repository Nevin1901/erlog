import React, { useEffect, useState } from "react";
import axios from "axios";
import { toLength } from "../utils";
import { redirect, useNavigate } from "react-router-dom";

interface LogCount {
  id: number;
  title: string;
  message: string;
  num: number;
}

function Index() {
  const router = useNavigate();
  const [logData, setLogData] = useState<LogCount[]>([]);

  useEffect(() => {
    const doWork = async () => {
      try {
        const { data } = await axios.post("http://127.0.0.1:8080/count");
        setLogData(data);
      } catch (err) {
        console.log(err);
      }
    };

    doWork();
  }, []);
  return (
    <div>
      <h1 className="font-semibold text-3xl">Logs</h1>
      {logData.map((data) => (
        <div
          onClick={() => router(`/logs/${data.id}`)}
          key={data.message}
          className="bg-gray-100 max-w-3xl rounded-md my-2 cursor-pointer mx-1 px-1.5 py-1"
        >
          <a className="font-bold">{data.title}</a>
          <p>{toLength(data.message, 100)}</p>
          <p>{data.num}</p>
        </div>
      ))}
      <h1>Hello</h1>
    </div>
  );
}

export default Index;
