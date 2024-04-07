import { docsConfig } from '@/config/docs';
import { DocsSidebarNav } from '@/components/sidebar-nav';

interface DocsLayoutProps {
  children: React.ReactNode;
}

export default function DocsLayout({ children }: DocsLayoutProps) {
  return (
    <div className="md:grid md:grid-cols-[220px_1fr] md:gap-6 flex-1 lg:grid-cols-[240px_1fr] lg:gap-10">
      <aside className="md:sticky md:block fixed top-14 z-30 hidden h-[calc(100vh-3.5rem)] w-full shrink-0 overflow-y-auto border-r py-6 pr-2 lg:py-10">
        <DocsSidebarNav items={docsConfig.sidebarNav} />
      </aside>
      {children}
    </div>
  );
}
