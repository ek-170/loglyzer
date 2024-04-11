"use client"

import { Button, PlusIcon, TrashBoxIcon } from '@/app/_components/atoms';
import { IconButton } from '@/app/_components/molecules';
import { Analysis } from '../_types/type';
import { useState } from 'react';
import Link from 'next/link';

export type AnalysisSideBarProps = {
  analyses: Analysis[];
};

export const SideBar = (props: AnalysisSideBarProps) => {
  const [selectedAnalysisId, setSelectedAnalysisId] = useState<string>("");

  return (
    <aside className="flex h-full w-[300px] flex-col divide-y-4 border-4 border-primary-100">
      <div className="grid place-items-center p-1.5">
        <IconButton endIcon={<PlusIcon width={20} color="green" />}>
          <span className="align-middle">
            <strong>Create New Analysis</strong>
          </span>
        </IconButton>
      </div>
      <ul className="overflow-y-auto">
        {props.analyses.length > 0 &&
          props.analyses.map((a) => {
            return (
              <li
                key={a.id}
                className={`flex border-b px-3 py-1.5 hover:bg-primary-50 ${selectedAnalysisId===a.id && "bg-primary-50"}`}
              >
                <div className="grow overflow-x-auto">
                  <Link className='w-full' href={`/analyses/${a.id}`}>
                    <Button width="w-full" onClick={() => setSelectedAnalysisId(a.id)}>
                        {a.id}
                    </Button>
                  </Link>
                </div>
                <IconButton endIcon={<TrashBoxIcon width={18} />} />
              </li>
            );
          })}
      </ul>
    </aside>
  );
};
