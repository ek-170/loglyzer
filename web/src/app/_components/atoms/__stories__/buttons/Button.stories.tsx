import type { Meta, StoryObj } from '@storybook/react';

import { Button } from '@/app/_components/atoms/buttons/Button';

const meta: Meta<typeof Button> = {
  component: Button,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  args: {
    onClick: () => {
      alert('button pushed!');
    },
    disabled: false,
  },
};

export default meta;
type Story = StoryObj<typeof Button>;

export const PositiveButton: Story = {
  args: {
    color: 'positive',
    padding: ['px-3', 'py-1.5'],
    width: 'w-16',
    children: <p className="text-white">Save</p>,
  },
};

export const NegativeButton: Story = {
  args: {
    color: 'negative',
    children: <p className="text-white">Cancel</p>,
  },
};
