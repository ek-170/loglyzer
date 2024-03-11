import type { Meta, StoryObj } from '@storybook/react';

import { Modal } from '@/app/_components/atoms/Modal';
import { ComponentProps, useState } from 'react';

const meta: Meta<typeof Modal> = {
  component: Modal,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  args: {
    children: (
      <div
        className="text-black shadow-slate-50 flex h-[200px] w-[300px] items-center
 justify-center rounded bg-white text-center shadow-xl"
      >
        This is Dialog
      </div>
    ),
  },
};

export default meta;
type Story = StoryObj<typeof Modal>;

const SampleModalWrapper = (args: ComponentProps<typeof Modal>) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  return (
    <>
      <button
        className="text-md mb-2 w-32 rounded-md bg-sky-400 p-1.5 text-center text-white"
        onClick={() => setIsOpen(!isOpen)}
      >
        {isOpen ? 'Close' : 'Open'} Modal
      </button>
      <Modal {...args} isOpen={isOpen} />
    </>
  );
};

export const SampleModal: Story = {
  args: {
    top: 'top-5',
    bottom: 'bottom-5',
  },
  render: (args) => <SampleModalWrapper {...args} />,
};
