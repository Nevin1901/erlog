import axios from "axios";
import { sortLogs } from "./utils";

export async function getLogs({ request }: any) {
  const url = new URL(request.url);
  let search = url.searchParams.get("search");

  if (!search) {
    search = "";
  }

  const { data } = await axios.post(
    "http://127.0.0.1:8080/count?search=" + search
  );
  const logData = data;

  return { logData, search };
}

export async function getLogById({ params }: any) {
  const { data } = await axios.post(`http://127.0.0.1:8080/logs/${params.id}`);
  const ordered = sortLogs(data, "ascending");
  return ordered;
}

export async function getIgnored() {
  const { data } = await axios.post("http://127.0.0.1:8080/ignored");
  return data;
}
