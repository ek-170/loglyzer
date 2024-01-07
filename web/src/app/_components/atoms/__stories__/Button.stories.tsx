import type { Meta, StoryObj } from '@storybook/react';

import Button from '@/app/_components/atoms/Button';

const buttonElementProps = {
  onClick: () => {
    alert('button pushed!');
  },
  disabled: false,
};

const meta: Meta<typeof Button> = {
  component: Button,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  args: {
    buttonElementProps: buttonElementProps,
  },
};

export default meta;
type Story = StoryObj<typeof Button>;

export const PositiveButton: Story = {
  args: {
    color: 'positive',
    padding: ['px-3', 'py-1.5'],
    width: 'w-16',
    children: <p>Save</p>,
  },
};

export const NegativeButton: Story = {
  args: {
    color: 'negative',
    children: <p>Cancel</p>,
  },
};
