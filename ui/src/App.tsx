import React, { useEffect } from "react";
import { Router, useNavigate } from "react-router-dom";

function App() {
  const router = useNavigate();
  useEffect(() => {
    router("/logs");
  }, []);
  return (
    <div>
      <h1>Hello</h1>
    </div>
  );
}

export default App;
