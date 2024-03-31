import Link from "next/link";

export const Header = () =>
  {
    return (
      <div className="bg-primary-600 sticky top-0 flex h-20 w-full px-5 text-white">
        <Link className="grid place-items-center px-4" href={"/"}>
          <strong className="text-[24px]">LogLyzer</strong>
        </Link>
      </div>
    );
  };
