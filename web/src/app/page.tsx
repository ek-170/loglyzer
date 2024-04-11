import Link from 'next/link';

export default function Home() {
  return (
    <main className="m-5">
      <p>This page is under construction.</p>
      <Link href={'/analyses'}>
        <p>
          <strong className="text-blue-500 hover:underline hover:decoration-blue-500">
            Go Analyses
          </strong>
        </p>
      </Link>
    </main>
  );
}
