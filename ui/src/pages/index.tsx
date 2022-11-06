import { useEffect } from "react";
import {
  Form,
  Outlet,
  useLoaderData,
  useNavigation,
  useParams,
} from "react-router-dom";
import IDSplit from "../components/IdSplit";
import LogContainer from "../components/LogContainer";
import { ILoaderData } from "../models";

function Index() {
  const { id } = useParams();
  const { logData, search }: ILoaderData = useLoaderData() as any;

  useEffect(() => {
    console.log(id);
    (document.getElementById("search") as any).value = search;
  }, [search]);

  if (!logData) {
    return <div></div>;
  }

  const isActive = (_id: number) => {
    if (!id) {
      return false;
    }

    if (parseInt(id) === _id) {
      return true;
    }

    return false;
  };

  return (
    <div className="h-screen">
      <div className="bg-gray-100 h-8 flex items-center px-1.5">
        <h1 className="font-bold text-xl">ErLog</h1>
      </div>
      <IDSplit className="h-full">
        <div className="bg-gray-100 p-1.5">
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
            <LogContainer
              log={data}
              key={data.id}
              selected={isActive(data.id)}
            />
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
