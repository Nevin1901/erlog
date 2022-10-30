import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Index from "./pages";
import Id from "./pages/logs/Id";
import { getIgnored, getLogById, getLogs } from "./loaders";
import Ignored from "./pages/ignored";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/logs",
    element: <Index />,
    loader: getLogs,
    children: [
      {
        path: "/logs/:id",
        element: <Id />,
        loader: getLogById,
        errorElement: <div>Log not found</div>,
      },
    ],
  },
  {
    path: "/ignored",
    element: <Ignored />,
    loader: getIgnored,
  },
]);

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
