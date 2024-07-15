/* eslint-disable tailwindcss/no-custom-classname */
import { Height, Width } from '@/app/_types/tailwind/Sizing';
import { ReactNode, useState } from 'react';
import { ThreeDots } from 'react-loader-spinner';
import { tv } from 'tailwind-variants';

export type TabItem = {
  id: string;
  header: ReactNode;
  children: ReactNode;
};

export type TabProps = {
  defaultId: string;
  items: TabItem[];
  width?: Width;
  height?: Height;
};

export const Tab = (props: TabProps) => {
  const { defaultId, items, width, height } = props;

  const [selectedId, setSelectedId] = useState<string>(defaultId);
  const widthClass = width || 'w-fit';
  const heightClass = height || 'h-fit';
  const tabHeader = tv({
    base: `border-b-2 bg-white px-2.5 py-1.5 text-gray-600`,
    variants: {
      selected: {
        selected: 'border-sky-500',
        notSelected: 'border-black-400 hover:border-black-500',
      },
    },
  });
  const selectedItem = items.find((i) => i.id === selectedId);
  return (
    <div className={`${widthClass} ${heightClass}`}>
      <div className="mb-4 flex">
        {
          // Generate Tab Headers
          items.map((i) => {
            return (
              <button
                key={i.id}
                className={tabHeader({
                  selected: selectedId === i.id ? 'selected' : 'notSelected',
                })}
                onClick={() => setSelectedId(i.id)}
                disabled={selectedId === i.id}
              >
                {i.header}
              </button>
            );
          })
        }
      </div>
      <div className="flex items-center justify-center">
        {selectedItem ? (
          selectedItem.children
        ) : (
          <ThreeDots height="50" width="50" color="gray" ariaLabel="loading" />
        )}
      </div>
    </div>
  );
};
