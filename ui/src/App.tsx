import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function App() {
  const router = useNavigate();
  useEffect(() => {
    // router("/logs");
  }, [router]);
  return (
    <div className="bg-gray-100 flex flex-col items-center justify-center h-screen">
      <h1 className="font-medium text-2xl mb-3">Select a Project</h1>
      <div className="flex flex-row space-x-5">
        <div className="bg-gray-200 shadow-md w-48 h-48">
          <h1>
            Readable (put an icon here which you can upload an image or use an
            already existing generic one)
          </h1>
        </div>
        <div className="bg-gray-200 shadow-md w-48 h-48">
          also make it so you can select a default project or list of projects
        </div>
        <div className="bg-gray-200 shadow-md w-48 h-48">
          also add a light + with text below these to add a new project
        </div>
      </div>
    </div>
  );
}

export default App;
