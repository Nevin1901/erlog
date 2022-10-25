import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function App() {
  const router = useNavigate();
  useEffect(() => {
    router("/logs");
  }, [router]);
  return (
    <div>
      <h1>Hello</h1>
    </div>
  );
}

export default App;
