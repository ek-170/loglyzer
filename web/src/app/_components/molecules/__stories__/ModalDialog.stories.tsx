import type { Meta, StoryObj } from '@storybook/react';

import { ModalDialog } from '@/app/_components/molecules';
import { ComponentProps, useRef, useState } from 'react';

const meta: Meta<typeof ModalDialog> = {
  component: ModalDialog,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ModalDialog>;

const SampleModalDialogWrapper = (args: ComponentProps<typeof ModalDialog>) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const modalRef = useRef(null)
  const handleOnClick = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    if (e.target === modalRef.current) {
      setIsOpen(false);
    }
  }
  const children = (
    <div
      className="flex-column h-[200px] w-[300px] items-center justify-center rounded
 bg-white text-center text-black shadow-slate-50"
    >
      <p className='my-10'>
        This is Dialog
      </p>
      <button className='bg-primary-400 rounded p-1 text-white' onClick={() => setIsOpen(false)} >
        閉じる
      </button>
    </div>
  );
  return (
    <>
      <button
        className="text-md mb-2 w-fit rounded-md bg-sky-400 p-1.5 text-center text-white"
        onClick={() => setIsOpen(!isOpen)}
      >
        {isOpen ? 'Close' : 'Open'} ModalDialog
      </button>
      <div className='text-[80px]'>
        This is a Background Contents
      </div>
      <ModalDialog {...args} ref={modalRef} onClick={e=>handleOnClick(e)} isOpen={isOpen} >
        {children}
      </ModalDialog>
    </>
  );
};

export const SampleModalDialog: Story = {
  render: (args) => <SampleModalDialogWrapper {...args} />,
};

export const ModalDialogWithHeader: Story = {
  args: {
    header: (
      <div className='rounded-t-md bg-blue-400 p-2 text-white'>Dialog Header</div>
    ),
  },
  render: (args) => <SampleModalDialogWrapper {...args} />,
};
