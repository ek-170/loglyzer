import type { Meta, StoryObj } from '@storybook/react';

import { Header } from '@/app/_components/organisms/Header';

const meta: Meta<typeof Header> = {
  component: Header,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    (Story) => (
      <div className="w-screen">
        <Story />
        <div className="flex-col items-center justify-center">
          <div className="h-[1000px] w-screen border bg-primary-100 text-[80px]">
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
type Story = StoryObj<typeof Header>;

export const PositiveHeader: Story = {};
