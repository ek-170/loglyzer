// "use client"

import { Header } from '../../_components/organisms/Header';
import { SideBar } from './_components/SideBar';
import { Analysis } from './_types/type';

type AnalysesProps = {
  children?: React.ReactNode;
};

async function fetchAnalyses() {
  const res = await fetch('http://localhost:9765/api/v1/analyses', {
    method: 'POST',
  });
  return await res.json().then((r) => {
    return r as Analysis[];
  });
}

export default async function AnalysesLayout({ children }: AnalysesProps) {
  const analyses = await fetchAnalyses();

  return (
    <>
      <Header />
      <div className="flex flex-row">
        <main className="h-[calc(100lvh-80px)]">
          <SideBar analyses={analyses}></SideBar>
        </main>
        {children}
      </div>
    </>
  );
}
