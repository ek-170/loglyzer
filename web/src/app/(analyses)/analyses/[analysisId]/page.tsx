'use client';

type AnalysisPageProps = {
  params: { analysisId: string };
};

export default function AnalysisPage({ params }: AnalysisPageProps) {
  // fetch from client

  return (
    <main className="text-[100px]">
      {params.analysisId}
      aaaaaaa
    </main>
  );
}
