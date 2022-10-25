import { Link, useLocation, useNavigate } from "react-router-dom";

export default function LinkWithQuery({ children, to }: any) {
  const router = useNavigate();
  const { search } = useLocation();

  return <div onClick={() => router(to + search)}>{children}</div>;
}
