import { useEffect } from "react";
import { Form, Outlet, useLoaderData } from "react-router-dom";
import IDSplit from "../components/IdSplit";
import LogContainer from "../components/LogContainer";
import { ILoaderData } from "../models";

function Index() {
  const { logData, search }: ILoaderData = useLoaderData() as any;

  useEffect(() => {
    (document.getElementById("search") as any).value = search;
  }, [search]);

  if (!logData) {
    return <div></div>;
  }

  return (
    <div>
      <div className="bg-gray-50 h-8 flex items-center px-1.5">
        <h1 className="font-bold text-xl">ErLog</h1>
      </div>
      <IDSplit>
        <div className="bg-gray-50 p-1.5 h-screen overflow-y-hidden">
          {/* <h1 className="font-semibold text-xl">Logs</h1> */}
          <Form id="search-form" role="search">
            <input
              id="search"
              placeholder="Search"
              name="search"
              className="bg-white p-0 border-2 border-gray-200 pl-1 focus:border-gray-300 focus:ring-gray-300 rounded-sm"
              type="text"
              defaultValue={search}
            />
          </Form>
          {logData.map((data) => (
            <LogContainer log={data} key={data.id} />
          ))}
        </div>
        <div className="">
          <Outlet />
        </div>
      </IDSplit>
    </div>
  );
}

export default Index;
