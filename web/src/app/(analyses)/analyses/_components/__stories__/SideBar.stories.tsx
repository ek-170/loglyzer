import type { Meta, StoryObj } from '@storybook/react';

import { SideBar } from '@/app/(analyses)/analyses/_components/SideBar';
import { Analysis } from '../../_types/type';

const analyses: Analysis[] = [
  {
    id: 'loglyzer-test-a',
    parseSources: [],
  },
  {
    id: 'loglyzer-test-abc',
    parseSources: [],
  },
  {
    id: 'loglyzer-test-xyz-123',
    parseSources: [],
  },
];

const meta: Meta<typeof SideBar> = {
  component: SideBar,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    (Story) => (
      <div className="flex h-[1000px] flex-row">
        <Story />
        <div className="flex-col items-center justify-center">
          <div className="h-[1000px] w-[600px] border bg-primary-100 text-[80px]">
            <p className="my-20 text-center">Contents 1</p>
            <p className="my-20 text-center">Contents 2</p>
            <p className="my-20 text-center">Contents 3</p>
            <p className="my-20 text-center">Contents 4</p>
            <p className="my-20 text-center">Contents 5</p>
          </div>
        </div>
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof SideBar>;

export const SampleSideBar: Story = {
  args: {
    analyses: analyses,
  },
};
