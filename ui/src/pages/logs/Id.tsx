import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ErLog } from "../../types";

export default function Id() {
  const [log, setLog] = useState<ErLog>();
  const params = useParams();

  useEffect(() => {
    const doWork = async () => {
      const { data } = await axios.post(
        `http://127.0.0.1:8080/logs/${params.id}`
      );
    };
    doWork();
  }, []);
  return (
    <div>
      <h1>{params.id}</h1>
    </div>
  );
}
