import { useEffect, useState } from "react";

export default function LoadingSpinner() {
  const [show, setShow] = useState<boolean>(false);
  useEffect(() => {
    const interval = setInterval(() => {
      setShow(true);
      clearInterval(interval);
    }, 100);
  }, []);

  if (!show) {
    return <div className="h-screen"></div>;
  } else
    return (
      <div className="h-screen flex justify-center items-center overflow-hidden mx-auto my-auto">
        <svg
          className="animate-spin -ml-1 mr-3 h-8 w-8 text-gray-400"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            className="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            strokeWidth="4"
          ></circle>
          <path
            className="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
        {/* <h1 className="text-2xl text-gray-400">Loading...</h1> */}
      </div>
    );
}
