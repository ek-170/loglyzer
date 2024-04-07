import { Button, PlusIcon, TrashBoxIcon } from '@/app/_components/atoms';
import { IconButton } from '@/app/_components/molecules';
import { Analysis } from '../_types/type';

export type AnalysisSideBarProps = {
  analyses: Analysis[];
};

export const SideBar = (props: AnalysisSideBarProps) => {
  return (
    <aside className="flex h-full w-[300px] flex-col divide-y-4 border border-primary-100">
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
                className="flex border-b px-3 py-1.5 hover:bg-primary-50"
              >
                <div className="grow overflow-x-auto">
                  <Button width="w-full">{a.id}</Button>
                </div>
                <IconButton endIcon={<TrashBoxIcon width={18} />} />
              </li>
            );
          })}
      </ul>
    </aside>
  );
};
