import axios from "axios";

export async function getLogs() {
  const { data } = await axios.post("http://127.0.0.1:8080/count");
  return data;
}

export async function getLogById({ params }: any) {
  const { data } = await axios.post(`http://127.0.0.1:8080/logs/${params.id}`);
  return data;
}
