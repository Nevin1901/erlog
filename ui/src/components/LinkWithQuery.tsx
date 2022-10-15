import { Link, useLocation } from "react-router-dom";

export default function LinkWithQuery({ children, to, ...props }: any) {
  const { search } = useLocation();
  console.log(search);

  return (
    <Link to={to + search} {...props}>
      {children}
    </Link>
  );
}
