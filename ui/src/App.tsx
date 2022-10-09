import React, { useEffect } from "react";
import { Router } from "react-router-dom";

function App() {
  useEffect(() => {
    window.location.href = "/logs";
  }, []);
  return (
    <div>
      <h1>Hello</h1>
    </div>
  );
}

export default App;
