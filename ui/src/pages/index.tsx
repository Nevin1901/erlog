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
      {logData.map((data) => (
        <div
          onClick={() => router(`/logs/${data.id}`)}
          key={data.message}
          className="bg-gray-200 rounded-sm my-1.5 cursor-pointer"
        >
          <a className="font-bold">{data.title}</a>
          <p>{toLength(data.message, 100)}</p>
        </div>
      ))}
      <h1>Hello</h1>
    </div>
  );
}

export default Index;
