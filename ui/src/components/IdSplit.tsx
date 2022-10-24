import { useState } from "react";
import { useLocation } from "react-router-dom";
import Split from "react-split";

export default function IDSplit({ children }: any) {
  const [sizes, setSizes] = useState<number[]>([50, 50]);
  const location = useLocation();

  const handleDrag = (_sizes: number[]) => {
    setSizes(_sizes);
  };

  if (location.pathname === "/logs") return <>{children}</>;
  else
    return (
      <Split
        style={{ display: "flex" }}
        gutterSize={5}
        onDrag={handleDrag}
        cursor="ew-resize"
        sizes={sizes}
      >
        {children}
      </Split>
    );
}